package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Loop over stdin, store previous line in var
	// If previous var exists, subtract from current var and output ± diff

	var previous_value int
	increases := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if (value - previous_value) >= 0 {
			increases += 1
		}

		previous_value = value
	}

	fmt.Printf("%d increases\n", increases)
}
