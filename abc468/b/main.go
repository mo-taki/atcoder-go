// 問題URL: https://atcoder.jp/contests/abc468/tasks/abc468_b
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

	ni()
	d := ni()

	s := ns()

	ans := 0
	for range d {
		s = "." + s + "."
	}

	// 2
	// .....G...G...G.....
	for i := 0 + d; i < len(s)-d; i++ {
		if s[i] == '.' {
			h := s[i-d : i+d+1]
			if !strings.Contains(h, "G") {
				ans++
			}
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
