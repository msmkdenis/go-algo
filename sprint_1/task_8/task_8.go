package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Тимофей записал два числа в двоичной системе счисления и попросил Гошу вывести их сумму, также в двоичной системе.
Встроенную в язык программирования возможность сложения двоичных чисел применять нельзя. Помогите Гоше решить задачу.

Решение должно работать за O(N), где N –— количество разрядов максимального числа на входе.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}
	fmt.Println(addBinary(s[0], s[1]))
}

func addBinary(a string, b string) string {
	carry := 0
	result := ""

	// Iterate over the strings from right to left
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		sum := carry

		// Add the corresponding digits from both strings
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// Append the sum mod 2 to the result string
		result = strconv.Itoa(sum%2) + result

		// Update the carry for the next iteration
		carry = sum / 2
	}

	// If there is still a carry, prepend it to the result string
	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}

	return result
}
