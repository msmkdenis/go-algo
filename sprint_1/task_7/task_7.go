package main

import "fmt"

/*
Вася реализовал функцию, которая переводит целое число из десятичной системы в двоичную.
Но, кажется, она получилась не очень оптимальной.

Попробуйте написать более эффективную программу.

Не используйте встроенные средства языка по переводу чисел в бинарное представление.
*/

func main() {
	var number uint
	fmt.Scan(&number)
	fmt.Println(uintToBinary(number))
}

func uintToBinary(num uint) string {
	if num == 0 {
		return "0"
	}

	binary := ""
	for num > 0 {
		remainder := num % 2
		binary = fmt.Sprint(remainder) + binary
		num /= 2
	}

	return binary
}
