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

	s := ns()

	for len(s) > 4 {
		if s == "dream" || s == "dreamer" || s == "erase" || s == "eraser" {
			fmt.Fprintln(wr, "YES")
			return
		}
		w := s[:5]

		if len(s) >= 8 {
			switch w {
			case "dream":
				if string(s[5:7]) == "er" {
					if string(s[7]) != "a" {

						s = s[7:]
					} else {
						s = s[5:]
					}
				} else {
					s = s[5:]
				}
			case "erase":
				if string(s[4:6]) == "er" {
					if string(s[7]) != "a" {
						s = s[6:]
					} else {
						s = s[5:]
					}
				} else {
					s = s[5:]
				}
			default:
				fmt.Fprintln(wr, "NO")
				return
			}
		} else {
			if s == "dream" || s == "erase" {
				fmt.Fprintln(wr, "YES")
				return
			} else {
				fmt.Fprintln(wr, "NO")
				return
			}
		}
	}

	fmt.Fprintln(wr, "NO")
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
