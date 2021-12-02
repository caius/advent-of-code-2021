package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BatchSize = 3

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var values []int
	var sums []int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		values = append(values, value)
	}

	// Loop number of times to get n batches of BatchSize
	loops := len(values) - (BatchSize - 1)
	for i := 0; i < loops; i++ {
		window := values[i : i+BatchSize]
		sum := 0
		for _, n := range window {
			sum += n
		}

		sums = append(sums, sum)
	}

	increases := 0
	previous := sums[0]
	for _, n := range sums[1:] {
		diff := n - previous
		if diff > 0 {
			increases += 1
		}
		previous = n
	}

	fmt.Printf("%d increases\n", increases)
}
