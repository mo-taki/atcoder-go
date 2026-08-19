// 問題URL: https://atcoder.jp/contests/abc462/tasks/abc462_b
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

	s := make([]int, n)
	t := make([]string, n)

	for i := range n {
		k := ni()
		for range k {
			a := ni()
			s[a-1]++
			t[a-1] += strconv.Itoa(i+1) + " "
		}
	}

	for i := range s {
		if s[i] == 0 {
			fmt.Fprintln(wr, s[i])
		} else {
			fmt.Fprintln(wr, s[i], strings.TrimRight(t[i], " "))
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
