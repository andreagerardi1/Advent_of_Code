package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 1
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	aux1 := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	times := make([]int, len(aux1))
	distances := make([]int, len(aux1))
	for i, k := range aux1 {
		times[i], _ = strconv.Atoi(k)
	}
	scanner.Scan()
	aux2 := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	for i, k := range aux2 {
		distances[i], _ = strconv.Atoi(k)
	}
	winnings := make([]int, len(times))
	for i := 0; i < len(times); i++ {
		for j := 0; j < times[i]; j++ {
			distance := (times[i] - j) * j
			if distance > distances[i] {
				winnings[i]++
			}
		}
	}
	for _, v := range winnings {
		total *= v
	}
	fmt.Println(total)
}
