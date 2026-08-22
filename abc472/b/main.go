// 問題URL: https://atcoder.jp/contests/abc472/tasks/abc472_b
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
	sum := 0
	i := []int{}

	for range n {
		l := ni()
		sum += l
		i = append(i, l)
	}

	half := sum / 2
	halfIdx := 0
	s := 0
	for idx, v := range i {
		s += v
		if s > half {
			halfIdx = idx
			break
		}
	}

	front := 0
	back := 0

	for _, v := range i {
		if front < half {
			front += v
		} else {
			back += v
		}
	}

	af := 0
	ab := 0
	bf := 0
	bb := 0

	for idx, v := range i {
		if idx < halfIdx {
			af += v
		} else {
			ab += v
		}
	}

	for idx, v := range i {
		if idx < halfIdx+1 {
			bf += v
		} else {
			bb += v
		}
	}

	a := abs(af - ab)
	b := abs(bf - bb)

	fmt.Fprintln(wr, min(a, b))
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
