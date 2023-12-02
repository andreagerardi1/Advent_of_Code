package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		var redMax, greenMax, blueMax int
		aux := strings.Split(text, ":")
		draws := strings.Split(aux[1], ";")
		for _, v := range draws {
			colors := strings.Split(v, ",")
			for _, k := range colors {
				colNum := strings.Split(k, " ")
				num, _ := strconv.Atoi(colNum[1])
				if colNum[2] == "green" && num > greenMax {
					greenMax = num
				} else if colNum[2] == "red" && num > redMax {
					redMax = num
				} else if colNum[2] == "blue" && num > blueMax {
					blueMax = num
				}
			}
		}
		mult := greenMax * redMax * blueMax
		sum += mult
	}
	fmt.Println(sum)
}
