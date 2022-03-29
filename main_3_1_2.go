package main

var groupCity = map[int][]string{
	10:   {}, // города с населением 10-99 тыс. человек
	100:  {}, // города с населением 100-999 тыс. человек
	1000: {}, // города с населением 1000 тыс. человек и более
}

var cityPopulation = map[string]int{}

func main() {
	for group, cities := range groupCity {
		for _, city := range cities {
			if _, ok := cityPopulation[city]; !ok && group == 100 {
				cityPopulation[city] = group
			} else if ok && group != 100 {
				delete(cityPopulation, city)
			}
		}
	}
}
