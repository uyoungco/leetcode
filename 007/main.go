package main

import (
	"fmt"
	"math"
)

// 7. 整数反转

var reverse = func(x int) (rev int) {
	for x != 0 {
		if rev <  math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev * 10 + digit
	}
	return rev
}

func main() {
	fmt.Println(reverse(123))
}
