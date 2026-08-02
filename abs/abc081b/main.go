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
	// N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	second := strings.Split(scanner.Text(), " ")

	cnt := 0

	for {

		for i, v := range second {
			val, _ := strconv.Atoi(v)
			if val%2 == 1 {
				fmt.Println(cnt)
				return
			} else {
				x, _ := strconv.Atoi(second[i])
				second[i] = strconv.Itoa(x / 2)
			}
		}
		cnt++

	}

}
