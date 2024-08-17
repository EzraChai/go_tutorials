package leetcode7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Reverse(x int) int {
	if x > 0 {
		xs := strconv.Itoa(x)
		xsA := strings.Split(xs, "")
		slices.Reverse(xsA)
		xsAJ := strings.Join(xsA, "")
		rev, err := strconv.Atoi(xsAJ)
		if err != nil {
			fmt.Println("FAIL")
		}
		if rev > 2147483648 {
			return 0
		}
		return rev
	} else if x < 0 {
		xs := strconv.Itoa(x)
		xsA := strings.Split(xs, "")
		xSA := xsA[1:]
		slices.Reverse(xSA)
		xsAJ := strings.Join(xSA, "")
		rev, err := strconv.Atoi(`-` + xsAJ)
		if err != nil {
			fmt.Println("FAIL")
		}
		if rev < -2147483648 {
			return 0
		}
		return rev
	} else {
		return 0
	}
}
