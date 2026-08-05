// 問題URL: https://atcoder.jp/contests/xxx/taks/xxx_x
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

	k := nf()
	km := k / 1000.0

	if km > 70 {
		fmt.Fprintf(wr, "%02d\n", 89)
	} else if km >= 35 {
		fmt.Fprintf(wr, "%02d\n", int(((km-30)/5 + 80)))
	} else if km >= 6 {
		fmt.Fprintf(wr, "%02d\n", int(km+50))
	} else if km >= 0.1 {
		fmt.Fprintf(wr, "%02d\n", int(km*10))
	} else {
		fmt.Fprintf(wr, "%02d\n", 0)
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
