package main

import "sync"

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			wg := new(sync.WaitGroup)
			wg.Add(2)

			go func() {
				wg.Done()
				x1 = fn(x1)
			}()

			go func() {
				wg.Done()
				x2 = fn(x2)
			}()

			wg.Wait()

			out <- x1 + x2
		}
	}()
}
