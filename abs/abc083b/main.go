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
	a := ni()
	b := ni()

	cnt := 0

	sum := 0

	for i := 1; i <= n; i++ {
		sn := strconv.Itoa(i)
		for _, v := range sn {
			sum += int(v - '0')
		}
		if sum >= a && b >= sum {
			cnt += i
		}
		sum = 0
	}
	fmt.Fprintln(wr, cnt)
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
