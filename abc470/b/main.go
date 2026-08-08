// 問題URL: https://atcoder.jp/contests/abc470/tasks/abc470_b
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

	m := make(map[int]int)

	for range n {
		c := ni()
		m[c]++
	}

	values := []int{}
	for _, v := range m {
		values = append(values, v)
	}

	mostV := 0
	mostK := 0
	ans := 0

	for k, v := range m {
		if mostV <= v {
			mostV = v
			mostK = k
		}
	}

	for k, v := range m {
		if k != mostK {
			ans += v
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
