// 問題URL: https://atcoder.jp/contests/abc463/tasks/abc463_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)

	n := ni()

	var hs, ls = []int{}, []int{}

	for range n {
		h := ni()
		hs = append(hs, h)
		l := ni()
		ls = append(ls, l)
	}

	ans := make([]int, n)
	ans[n-1] = hs[n-1]
	for i := n - 2; i >= 0; i-- {
		ans[i] = max(hs[i], ans[i+1])
	}

	q := ni()

	for range q {
		t := ni()
		k := sort.Search(n, func(i int) bool { return ls[i] > t })
		fmt.Fprintln(wr, ans[k])
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
