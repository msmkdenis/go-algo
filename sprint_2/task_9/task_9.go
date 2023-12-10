package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://contest.yandex.ru/contest/22779/problems/I

/*
Астрологи объявили день очередей ограниченного размера.
Тимофею нужно написать класс MyQueueSized, который принимает параметр max_size,
означающий максимально допустимое количество элементов в очереди.

Помогите ему —– реализуйте программу, которая будет эмулировать работу такой очереди.
Функции, которые надо поддержать, описаны в формате ввода.

Формат ввода
В первой строке записано одно число — количество команд, оно не превосходит 5000.
Во второй строке задан максимально допустимый размер очереди, он не превосходит 5000.
Далее идут команды по одной на строке. Команды могут быть следующих видов:

push(x) — добавить число x в очередь;
pop() — удалить число из очереди и вывести на печать;
peek() — напечатать первое число в очереди;
size() — вернуть размер очереди;
При превышении допустимого размера очереди нужно вывести «error».
При вызове операций pop() или peek() для пустой очереди нужно вывести «None».
*/

type LinkedList struct {
	head        *Node
	tail        *Node
	maxSize     int
	currentSize int
}

func NewLinkedList(maxSize int) *LinkedList {
	return &LinkedList{
		head:        nil,
		tail:        nil,
		maxSize:     maxSize,
		currentSize: 0,
	}
}

type Node struct {
	value int
	next  *Node
}

func NewNode(value int, next *Node) *Node {
	return &Node{
		value: value,
		next:  next,
	}
}

func (s *LinkedList) PushBack(value int) {
	if s.currentSize == s.maxSize {
		fmt.Println("error")
		return
	}
	if s.currentSize == 0 {
		s.head = NewNode(value, nil)
		s.tail = s.head
		s.currentSize++
		return
	} else if s.currentSize == 1 {
		newNode := NewNode(value, nil)
		s.head.next = newNode
		s.tail = newNode
		s.currentSize++
		return
	} else {
		newNode := NewNode(value, nil)
		s.tail.next = newNode
		s.tail = newNode
		s.currentSize++
	}
}

func (s *LinkedList) peek() {
	if s.currentSize == 0 {
		fmt.Println("None")
	} else {
		fmt.Println(s.head.value)
	}
}

func (s *LinkedList) size() {
	fmt.Println(s.currentSize)
}

func (s *LinkedList) PopFront() {
	if s.currentSize == 0 {
		fmt.Println("None")
	} else if s.currentSize == 1 {
		fmt.Println(s.head.value)
		s.head = nil
		s.tail = nil
		s.currentSize--
	} else {
		fmt.Println(s.head.value)
		s.head = s.head.next
		s.currentSize--
	}
}

type Queue interface {
	PushBack(value int)
	peek()
	size()
	PopFront()
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

	queueSize, _ := strconv.Atoi(inputData[1])
	queue := NewLinkedList(queueSize)
	processCommands(inputData[2:], queue)
}

func processCommands(input []string, queue Queue) {
	for _, v := range input {
		command := strings.Split(v, " ")
		if command[0] == "peek" {
			queue.peek()
		} else if command[0] == "size" {
			queue.size()
		} else if command[0] == "pop" {
			queue.PopFront()
		} else if command[0] == "push" {
			x, _ := strconv.Atoi(command[1])
			queue.PushBack(x)
		}
	}
}
