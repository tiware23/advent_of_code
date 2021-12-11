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

	sliceOfMe := checkMeasurement(measurement)

	fmt.Println("Your puzzle answer was", len(sliceOfMe))
}

func checkMeasurement(list []int) []int {

	countM := []int{}

	for i := 0; i < len(list); i++ {
		n := i + 1
		if n == len(list) {
			break
		} else {
			if list[i] > list[n] {
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

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		textToInt, _ := strconv.Atoi(scanner.Text())
		sValues = append(sValues, textToInt)
	}

	err = f.Close()
	if err != nil {
		return sValues, err
	}

	if scanner.Err() != nil {
		return sValues, scanner.Err()
	}

	return sValues, nil
}
