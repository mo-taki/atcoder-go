// https://atcoder.jp/contests/abc469/tasks/abc469_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)

	n := ni()
	s := ns()

	s = "x" + s + "x"

	ans := 0

	for i := 1; i < n+1; i++ {
		if string(s[i]) == "o" {
			continue
		}
		if string(s[i-1]) == "x" && string(s[i+1]) == "x" {
			ans++
		}
	}
	fmt.Fprintln(wr, ans)

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
