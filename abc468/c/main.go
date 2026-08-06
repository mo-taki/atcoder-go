// 問題URL: https://atcoder.jp/contests/abc468/tasks/abc468_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

var n int
var cur []int
var used [11]bool
var ans int
var p []int
var q []int

func main() {
	defer wr.Flush()
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)

	n = ni()

	for range n {
		p = append(p, ni())
	}

	for range n {
		q = append(q, ni())
	}

	dfs()

	fmt.Fprintln(wr, ans)
}

func dfs() {
	if len(cur) == n {
		if slices.Compare(cur, p) == 1 && slices.Compare(q, cur) == 1 {
			ans++
		}
		return
	}
	for v := 1; v <= n; v++ {
		if used[v] {
			continue
		}
		used[v] = true
		cur = append(cur, v)
		dfs()
		cur = cur[:len(cur)-1]
		used[v] = false
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
