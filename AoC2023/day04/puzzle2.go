package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	m := make(map[int]int)
	var cext int
	var total int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var cint int
		text := scanner.Text()
		winning := strings.Fields(strings.Split(strings.Split(text, ":")[1], "|")[0])
		mine := strings.Fields(strings.Split(strings.Split(text, ":")[1], "|")[1])
		for i := 0; i <= m[cext]; i++ {
			cint = 1
			for _, k := range winning {
				if slices.Contains(mine, k) {
					m[cext+cint] += 1
					cint++
				}
			}
		}
		cext++
	}
	total += cext
	for _, v := range m {
		total += v
	}
	fmt.Println(total)
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Println("took this much: ", executionTime)
}
