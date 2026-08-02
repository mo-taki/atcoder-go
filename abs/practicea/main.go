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
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	second := strings.Split(scanner.Text(), " ")
	b, _ := strconv.Atoi(second[0])
	c, _ := strconv.Atoi(second[1])

	scanner.Scan()
	s := scanner.Text()

	fmt.Println(a+b+c, s)
}
