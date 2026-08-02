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
	second := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(second[0])
	b, _ := strconv.Atoi(second[1])

	x := a * b
	if x%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
