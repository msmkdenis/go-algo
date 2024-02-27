package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(catalan(a))
}

func catalan(n int) int {
	res := 0

	if n <= 1 {
		return 1
	}

	for i := 0; i < n; i++ {
		res += catalan(i) * catalan(n-i-1)
	}
	return res
}
