package main

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)

		select {
		case n := <-firstChan:
			outputChan <- n * n

		case n := <-secondChan:
			outputChan <- n * 3

		case <-stopChan:
			return
		}
	}()

	return outputChan
}
