// 問題URL: https://atcoder.jp/contests/xxx/taks/xxx_x
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)

	n := ni()
	m := ni()
	balls := make(map[int]int)
	for i := range m {
		balls[i+1] = -1
	}

	for range n {
		c := ni()
		s := ni()

		balls[c] = max(balls[c], s)
	}
	var ans []string

	for i := 1; i <= len(balls); i++ {
		ans = append(ans, strconv.Itoa(balls[i]))
	}
	fmt.Println(strings.Join(ans, " "))
}

func ns() string  { sc.Scan(); return sc.Text() }
func ni() int     { i, _ := strconv.Atoi(ns()); return i }
func nf() float64 { f, _ := strconv.ParseFloat(ns(), 64); return f }
func nis(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = ni()
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func digitSum(x int) int {
	s := 0
	for ; x > 0; x /= 10 {
		s += x % 10
	}
	return s
}
func isEven(x int) bool {
	return x%2 == 0
}
