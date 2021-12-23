// Advent of Code 2021 - Day1
// Author: Thiago N Cavalcante
// 2021-12-23

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	measurement, err := readFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	n := 3
	listSum := []int{}

	for i := 0; i < len(measurement); i++ {
		t := measurement[i:n]
		listSum = append(listSum, sum(t))
		n++

		if cap(measurement) < n {
			break
		}
	}

	sliceOfMe := checkMeasurement(listSum)
	fmt.Printf("There are %d sums\n", len(sliceOfMe))
}

func sum(list []int) int {
	count := 0
	for _, v := range list {
		count += v
	}
	return count
}

func checkMeasurement(list []int) []int {

	countM := []int{}

	for i := 0; i < len(list); i++ {
		n := i + 1
		if n == len(list) {
			break
		} else {
			if list[i] >= list[n] {
				continue
			} else {
				countM = append(countM, list[n])
			}
		}

	}

	return countM
}

func readFile(file string) ([]int, error) {
	var sValues []int
	f, err := os.Open(file)
	if err != nil {
		return sValues, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		textToInt, _ := strconv.Atoi(scanner.Text())
		sValues = append(sValues, textToInt)
	}

	if scanner.Err() != nil {
		return sValues, scanner.Err()
	}

	return sValues, nil

}
