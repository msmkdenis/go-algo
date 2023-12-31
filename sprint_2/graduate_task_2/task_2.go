package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22781/run-report/103774231/

/*
-- ПРИНЦИП РАБОТЫ --
Входными данными является строка операндов и знаков операций, записанных в обратной польской нотации.

Калькуляция выражения, записанного в обратной польской нотации возможна с использованием интерфейса stack (стек).
Стек по своей сути является интерфейсом над структурой данных, т.к. предоставляет контракт в виде определенной реализации методов,
как минимум методов Push - добавление элемента на вершину стека и Pop - удаление элемента с вершины stack.
Указанные два метода реализуют логику LIFO для работы с данными через интерфейс стек.

-- РЕАЛИЗАЦИЯ --
Преобразуем строку данных в слайс строк c разделителем " ".
В цикле переберём каждый элемент полученного слайса.
Если элемент не является знаком операций, то добавляем его на вершину стека (push(x)).
Если элемент является знаком (*, +, /, -), то извлекаем два последних элемента из стека (pop()): первый элемент - b, второй - a,
выполняем арифметическое действие и записываем его в переменную result, которую добавляем на вершину стека.
В зависимости от знака выполняем логику:
	- если "+" - выполняем а + b,
	- если "-" - выполняем a - b,
	- если "*" - выполняем a * b,
	- если "/" - выполняем math.floor(a/b) - по условию требуется осуществить целочисленное деление.
В результате на вершине stack окажется результат выражения, введенного в обратной польской нотации.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
В постфиксной нотации операнды расположены перед знаками операций.
Соответственно мы можем "прочитать" элементы строки справа налево, при этом если встречается знак операций мы должны
извлечь два следующих после знака элемента из коллекции и совершить над нами арифметическое действие, затем результат добавить к коллекции.
Удобным способом реализации такой алгоритма является использование интерфейса стек и его двух ключевых методов Push и Pop,
т.к. стек реализует логику LIFO.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
В основе реализованного стека лежит слайс.
Метод Push добавляет элемент в конец слайса (на вершину стека), в общем амортизационная сложность добавления элемента в конец слайса составляет О(1).
Также нам известен размер входящей строки мы можем создать стек, использующий слайс с заранее заданным параметром capacity
и избежать необходимости копирования внутреннего массива, заранее аллоцировав память.
Метод Pop удаляет последний элемент из слайса (с вершины стека), удаление последнего элемента слайса выполняется за О(1).
Общая временная сложность алгоритма линейна (зависит от кол-ва операндов и знаков операций) и составляет 0(n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность алгоритма линейно зависит от объёма входных данных (операндов и знаков операций),
соответственно будет занято n количество памяти. Отсюда можно сделать вывод, что пространственная сложность O(n).
*/

type SimpleStack struct {
	stack    []int
	capacity int
}

func NewSimpleStack(capacity int) *SimpleStack {
	return &SimpleStack{
		capacity: capacity,
		stack:    make([]int, 0, capacity),
	}
}

func (s *SimpleStack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *SimpleStack) Pop() int {
	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:(len(s.stack) - 1)]
	return x
}

type Stack interface {
	Push(x int)
	Pop() int
}

func calculateRPN(inputTokens []string, stack Stack) int {
	for _, token := range inputTokens {
		switch token {
		case "+", "-", "*", "/":
			b := stack.Pop()
			a := stack.Pop()

			var calculation int
			switch token {
			case "+":
				calculation = a + b
			case "-":
				calculation = a - b
			case "*":
				calculation = a * b
			case "/":
				calculation = floorDiv(a, b)
			}
			stack.Push(calculation)
		default:
			digit, _ := strconv.Atoi(token)
			stack.Push(digit)
		}
	}

	return stack.Pop()
}

func floorDiv(a, b int) int {
	r := a / b
	if (a^b) < 0 && (r*b != a) {
		r--
	}
	return r
}

func main() {
	const maxCapacity = 150 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var inputData []string
	for scanner.Scan() {
		line := scanner.Text()
		inputData = append(inputData, line)
		if line == "" {
			break
		}
	}

	inputTokens := strings.Split(inputData[0], " ")

	stack := NewSimpleStack(len(inputTokens))
	result := calculateRPN(inputTokens, stack)

	fmt.Println(result)
}
