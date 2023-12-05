package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	var total int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var counter int
		text := scanner.Text()
		winning := strings.Fields(strings.Split(strings.Split(text, ":")[1], "|")[0])
		mine := strings.Fields(strings.Split(strings.Split(text, ":")[1], "|")[1])
		fmt.Println(winning, mine)
		for _, k := range winning {
			if slices.Contains(mine, k) {
				counter++
			}
		}
		total += int(math.Pow(2.0, float64(counter-1)))
	}
	fmt.Println(total)
}
