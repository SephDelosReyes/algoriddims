package main

import (
	"fmt"
)

func main() {
	fmt.Println(3 + (10-3)/2)
}

func linearSearch(a []int, val int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == val {
			return true
		}
	}
	return false
}

func binarySearch(a []int, val int) bool {
	hi := len(a) - 1
	lo := 0
	for {
		if lo <= hi {
			m := lo + (hi-lo)/2
			v := a[m]
			fmt.Println("hi: ", hi, "lo: ", lo, "m: ", m, "v: ", v)
			if v == val {
				return true
			} else if v > val {
				hi = m - 1
			} else {
				lo = m + 1
			}
		} else {
			break
		}
	}
	return false
}
