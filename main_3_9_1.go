package main

func work()

func main() {
	quit := make(chan int)

	go func() {
		work()
		close(quit)
	}()

	<-quit
}
