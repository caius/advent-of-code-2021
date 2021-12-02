package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	Horizontal int
	Vertical   int
	Aim        int
}

func (s Submarine) Stats() string {
	return fmt.Sprintf("Horizontal: %d\nVertical: %d\nPosition: %d", s.Horizontal, s.Vertical, (s.Horizontal * s.Vertical))
}

func main() {
	sub := Submarine{}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := strings.TrimSpace(input.Text())
		bits := strings.Split(line, " ")
		command := bits[0]
		distance, err := strconv.Atoi(bits[1])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Moving %s %d\n", command, distance)

		switch command {
		case "forward":
			sub.Horizontal += distance
			sub.Vertical += (sub.Aim * distance)
		case "up":
			sub.Aim -= distance
		case "down":
			sub.Aim += distance
		default:
			panic("Unknown command")
		}
	}

	fmt.Println(sub.Stats())
}
