package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type City struct {
	min   float64
	max   float64
	count uint32
	total float64
}

func strToFloat(input string) float64 {
	s, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return s
}

func main() {
	mappedvals := make(map[string]City, 600)

	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), ";")
		cityName := splitted[0]
		measurement := strToFloat(splitted[1])

		if val, ok := mappedvals[cityName]; ok {
			if measurement < val.min {
				val.min = measurement
			}
			if measurement > val.max {
				val.max = measurement
			}
			val.count++
			val.total += measurement
			mappedvals[cityName] = val
		} else {
			mappedvals[cityName] = City{
				min:   measurement,
				max:   measurement,
				count: 1,
				total: measurement,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k, v := range mappedvals {
		fmt.Printf("%s: min: %.2f, max: %.2f, avg: %.2f\n", k, v.min, v.max, v.total/float64(v.count))
	}
}
