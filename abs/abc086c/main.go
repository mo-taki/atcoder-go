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

	t := ni()
	x := ni()
	y := ni()

	d := x + y
	if t < d || isEven(t) != isEven(d) {
		fmt.Fprintln(wr, "No")
		return
	}

	for range n - 1 {
		t = ni() - t
		x = abs(ni() - x)
		y = abs(ni() - y)

		d = x + y

		if t < d || isEven(t) != isEven(d) {
			fmt.Fprintln(wr, "No")
			return
		}
	}
	fmt.Fprintln(wr, "Yes")
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

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
