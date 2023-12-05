package main

import (
	"fmt"
	"slices"
)

/*
Основная теорема арифметики говорит: любое число раскладывается на произведение простых множителей единственным образом, с точностью до их перестановки. Например:

Число 8 можно представить как 2 × 2 × 2.
Число 50 –— как 2 × 5 × 5 (или 5 × 5 × 2, или 5 × 2 × 5). Три варианта отличаются лишь порядком следования множителей.
Разложение числа на простые множители называется факторизацией числа.

Напишите программу, которая производит факторизацию переданного числа.
*/

func main() {
	var number int
	fmt.Scan(&number)
	factors := factorization(number)
	slices.Sort(factors)
	for n, v := range factors {
		if n == len(factors)-1 {
			fmt.Print(v)
		} else {
			fmt.Print(v, " ")
		}
	}
}

func factorization(num int) []int {
	var factors []int

	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			factors = append(factors, i)
			num /= i
		}
	}

	if num != 1 {
		factors = append(factors, num)
	}
	return factors
}
