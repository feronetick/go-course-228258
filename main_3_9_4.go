package main

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		sum := 0

		for {
			select {
			case n := <-arguments:
				sum += n

			case <-done:
				output <- sum
				return
			}
		}
	}()

	return output
}
