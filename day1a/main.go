package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (int, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return 0, errors.New("Empty line")
	} else {
		return strconv.Atoi(line)
	}
}

func main() {
	// Loop over stdin, store previous line in var
	// If previous var exists, subtract from current var and output ± diff

	var previous_value int
	increases := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	previous_value, err := parseLine(scanner.Text())
	if err != nil {
		panic(err)
	}
	for scanner.Scan() {
		value, err := parseLine(scanner.Text())
		if err != nil {
			continue
		}

		if (value - previous_value) >= 0 {
			increases += 1
		}

		previous_value = value
	}

	fmt.Printf("%d increases\n", increases)
}
