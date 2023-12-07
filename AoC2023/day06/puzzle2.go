package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var strAux string
	var total int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	aux1 := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	for _, k := range aux1 {
		strAux += k
	}
	time, _ := strconv.Atoi(strAux)
	fmt.Println(time)
	scanner.Scan()
	strAux = ""
	aux2 := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	for _, k := range aux2 {
		strAux += k
	}
	distance, _ := strconv.Atoi(strAux)
	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			total++
		}
	}
	fmt.Println(total)
}
