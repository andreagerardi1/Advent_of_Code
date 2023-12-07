package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sourcesString := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	var source []int
	for _, k := range sourcesString {
		n, _ := strconv.Atoi(k)
		source = append(source, n)
	}
	sourcebool := make([]bool, len(source))
	scanner.Scan()
	var c int
	for scanner.Scan() {
		//fmt.Println("riga", c, ":", scanner.Text())
		if len(scanner.Text()) == 0 || strings.Contains(scanner.Text(), ":") {
			sourcebool = make([]bool, len(source))
		} else {
			nums := strings.Fields(scanner.Text())
			d, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			r, _ := strconv.Atoi(nums[2])
			for i, k := range source {
				if k >= s && k <= s+r && !sourcebool[i] {
					source[i] -= s - d
					sourcebool[i] = true
				}
			}
			//fmt.Println("all'iterazione ", c, " source è: ", source)
		}
		c++
	}
	min := source[0]
	for _, k := range source {
		if k < min {
			min = k
		}
	}
	fmt.Println(min)
}
