package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func explodeBits(value string) []int {
	var values []int
	for _, runeValue := range value {
		var value int
		if runeValue == 48 {
			value = 0
		} else {
			value = 1
		}

		values = append(values, value)
	}

	return values
}

func binary_array_2_int(arr []int) int {
	var str string
	for _, n := range arr {
		str = fmt.Sprintf("%s%d", str, n)
	}

	dec, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(dec)
}

func main() {
	var data [][]int

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := strings.TrimSpace(input.Text())
		bits := explodeBits(line)

		// Initialize data structure now we know number of columns
		if len(data) == 0 {
			data = make([][]int, len(bits))
		}

		for i, n := range bits {
			data[i] = append(data[i], n)
		}
	}

	// Figure out most/least values in columns
	var column_most []int
	var column_least []int

	for _, column := range data {
		// Count number of 1s/0s in column
		ones, zeros := 0, 0
		for _, n := range column {
			if n == 0 {
				zeros++
			} else {
				ones++
			}
		}

		if zeros < ones {
			// More 1s in column
			column_most = append(column_most, 1)
			column_least = append(column_least, 0)
		} else {
			// More 0s in column
			column_most = append(column_most, 0)
			column_least = append(column_least, 1)
		}
	}

	gamma_rate := binary_array_2_int(column_most)
	epsilon_rate := binary_array_2_int(column_least)

	fmt.Printf("Gamma Rate: %d\n", gamma_rate)
	fmt.Printf("Epsilon Rate: %d\n", epsilon_rate)
	fmt.Printf("Power Consumption: %d\n", gamma_rate*epsilon_rate)
}
