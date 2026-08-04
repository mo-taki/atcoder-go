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

func main() {
	defer wr.Flush()
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Split(bufio.ScanWords)

	n := ni()

	var cards []int

	for range n {
		cards = append(cards, ni())
	}

	c := [2]int{0, 0}
	i := 0

	for {
		maxC := slices.Max(cards)
		c[i] += maxC
		idx := slices.Index(cards, maxC)
		cards = slices.Delete(cards, idx, idx+1)

		if len(cards) == 0 {
			fmt.Fprintln(wr, c[0]-c[1])
			return
		}

		if i == 0 {
			i = 1
		} else {
			i = 0
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
