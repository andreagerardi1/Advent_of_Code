package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mp := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	var tot int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var c1, c2 int
		text := []rune(scanner.Text())
		var min int = len(string(text))
		var max int = 0
		for k, v := range mp {
			idx := strings.Index(string(text), k)
			if idx < min && idx != -1 {
				min = idx
				c1 = v
			}
		}
		for i := 0; i < min; i++ {
			n, e := strconv.Atoi(string(text[i]))
			if e == nil {
				c1 = n
				break
			}
		}

		for k, v := range mp {
			idx := strings.LastIndex(string(text), k)
			if idx > max {
				max = idx
				c2 = v
			}
		}
		for i := len(text) - 1; i >= max; i-- {
			n, e := strconv.Atoi(string(text[i]))
			if e == nil {
				c2 = n
				break
			}
		}
		totP := (c1*10 + c2)
		fmt.Println("number of the line: ", totP)
		tot += totP
	}
	fmt.Println(tot)
}
