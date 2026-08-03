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
	y := ni()

	for j := 0; j <= n; j++ {
		for i := 0; i <= n; i++ {
			if (n-i-j) >= 0 && n == i+j+(n-i-j) && i*10000+j*5000+(n-i-j)*1000 == y {
				fmt.Fprintln(wr, i, j, n-i-j)
				return
			}
		}
	}

	fmt.Fprintln(wr, "-1", "-1", "-1")
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
