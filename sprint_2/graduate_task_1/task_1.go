package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data []int
	head int
	tail int
	size int
}

func NewDeque(capacity int) *Deque {
	return &Deque{
		data: make([]int, capacity),
	}
}

func (d *Deque) PushFront(x int) {
	if d.size == len(d.data) {
		fmt.Println("error")
		return
	}
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = x
	d.size++
}

func (d *Deque) PushBack(x int) {
	if d.size == len(d.data) {
		fmt.Println("error")
		return
	}
	d.data[d.tail] = x
	d.tail = (d.tail + 1) % len(d.data)
	d.size++
}

func (d *Deque) PopFront() (int, bool) {
	if d.size == 0 {
		fmt.Println("error")
		return 0, false
	}
	x := d.data[d.head]
	fmt.Println(x)
	d.head = (d.head + 1) % len(d.data)
	d.size--
	return x, true
}

func (d *Deque) PopBack() (int, bool) {
	if d.size == 0 {
		fmt.Println("error")
		return 0, false
	}
	d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	x := d.data[d.tail]
	fmt.Println(x)
	d.size--
	return x, true
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

	dequeueCapacity, _ := strconv.Atoi(inputData[1])
	deque := NewDeque(dequeueCapacity)
	processCommands(inputData[2:], deque)
}

func processCommands(inputData []string, deque *Deque) {
	for _, v := range inputData {
		command := strings.Split(v, " ")
		if command[0] == "push_back" {
			x, _ := strconv.Atoi(command[1])
			deque.PushBack(x)
		}
		if command[0] == "push_front" {
			x, _ := strconv.Atoi(command[1])
			deque.PushFront(x)
		}
		if command[0] == "pop_front" {
			deque.PopFront()
		}
		if command[0] == "pop_back" {
			deque.PopBack()
		}
	}
}
