// 問題URL: https://atcoder.jp/contests/abc472/tasks/abc472_c
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
	m := ni()
	k := ni()

	s := 0

	ate := []int{}

	for i := 0; i < n; i++ {
		a := ni()

		if s+a > k {
			fmt.Fprintln(wr, "No")
			ate = append(ate, 0)
		} else {
			fmt.Fprintln(wr, "Yes")
			ate = append(ate, a)
			s += a
		}

		if len(ate) >= m {
			s = s - ate[0]
			ate = ate[1:]
		}

	}
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
