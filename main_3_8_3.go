package main

func removeDuplicates(input, output chan string) {
	prev := ""

	for curr := range input {
		if curr != prev {
			prev = curr
			output <- curr
		}
	}

	close(output)
}
