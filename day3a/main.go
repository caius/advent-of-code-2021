package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ReportEntry struct {
	Value string
}

func (re *ReportEntry) Bits() []int {
	var values []rune
	for _, runeValue := range re.Value {
		values = append(values, runeValue)
	}

	var bits []int
	bits = make([]int, len(values))
	for idx, runeValue := range values {
		var value int
		if runeValue == 48 {
			value = 0
		} else {
			value = 1
		}
		// Reverse the order so bit 1 is at index 0
		bits[len(values)-1-idx] = value
	}

	return bits
}

func main() {
	var entries []ReportEntry

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := strings.TrimSpace(input.Text())
		entry := ReportEntry{Value: line}
		entries = append(entries, entry)
		fmt.Printf("%v\n", entry.Bits())
	}

	// Map from rows into columns in array entries
	var columns [][]int
	columns = make([][]int, len(entries[0].Bits()))
	for _, entry := range entries {
		for idx, value := range entry.Bits() {
			columns[idx] = append(columns[idx], value)
		}
	}

}
