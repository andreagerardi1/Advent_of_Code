package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var mp = map[string]int{"red": 12, "green": 13, "blue": 14}
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		var redMax, greenMax, blueMax int
		aux := strings.Split(text, ":")
		game := strings.Split(aux[0], " ")
		gameNumber, _ := strconv.Atoi(game[1])
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
		if greenMax <= mp["green"] && redMax <= mp["red"] && blueMax <= mp["blue"] {
			sum += gameNumber
		}
	}
	fmt.Println(sum)
}
