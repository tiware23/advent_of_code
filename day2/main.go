package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	name  string
	value int
}

var (
	forward = 0
	down    = 0
)

func main() {
	t, _ := readFile("input.txt")

	for _, v := range t {
		if v.name == "forward" {
			forward += v.value
		} else if v.name == "down" {
			down += v.value
		} else {
			down -= v.value
		}

	}
	fmt.Printf("Planned Course: %d\n", forward*down)

}

func readFile(file string) ([]Command, error) {
	var cValues []Command
	f, err := os.Open(file)
	if err != nil {
		return cValues, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		textIn := scanner.Text()
		textS := strings.Split(strings.TrimSpace(textIn), " ")
		textInt, _ := strconv.Atoi(textS[1])

		cValues = append(cValues, Command{name: textS[0], value: textInt})

	}

	err = f.Close()
	if err != nil {
		return cValues, err
	}

	if scanner.Err() != nil {
		return cValues, scanner.Err()
	}

	return cValues, nil
}
