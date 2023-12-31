package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22781/run-report/103774302/

/*
-- ПРИНЦИП РАБОТЫ --
Для реализации интерфейса Deque с константной сложностью вставки и удаления элементов в начало и конец воспользуемся
слайсом с заранее заданным параметром capacity и имплементируем кольцевой буфер.
Также Deque неограниченного размера с константной сложностью вставки и удаления в начало и конец можно реализовать например с помощью связных списков,
но указанная реализация потребует больше памяти (необходимо также хранить указатели на следующие элементы).

-- РЕАЛИЗАЦИЯ --
Для реализации интерфейса Deque на основе слайса с заранее заданным параметром capacity требуется поддерживать информацию о:
	- размере слайса
	- количестве элементов в слайсе
	- индексе головы очереди
	- индексе хвоста очереди
В процессе обработки команд динамически высчитываются индексы хвоста и головы очереди, операции вставки выполняются по индексу.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Отслеживая положения головы и хвоста очереди, а также осуществляя операции вставки и удаления в начало и конец по соответствующим индексам,
мы добились того, что слайсу не требуются перестановки внутреннего массива при вставке элементов в начало.
Таким образом все операции вставки и удаления в начало и конец будут выполняться за O(1).

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Каждая операция интерфейса Deque в отдельности выполняется с константной сложностью О(1).
Общая временная сложность операций с интерфейсом Deque линейно зависит от кол-ва операций с интерфейсом и равна O(n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность интерфейса Deque линейно зависит от объёма входных данных (элементов добавленных в слайс),
соответственно будет занято n количество памяти. Отсюда можно сделать вывод, что пространственная сложность O(n).
*/

type SliceDeque struct {
	capacity int
	data     []int
	head     int
	tail     int
	size     int
}

func NewSliceDeque(capacity int) *SliceDeque {
	return &SliceDeque{
		capacity: capacity,
		data:     make([]int, capacity),
	}
}

func (d *SliceDeque) PushFront(x int) error {
	if d.isFull() {
		return errors.New("error")
	}
	d.head = (d.head - 1 + d.capacity) % d.capacity
	d.data[d.head] = x
	d.size++
	return nil
}

func (d *SliceDeque) PushBack(x int) error {
	if d.isFull() {
		return errors.New("error")
	}
	d.data[d.tail] = x
	d.tail = (d.tail + 1) % d.capacity
	d.size++
	return nil
}

func (d *SliceDeque) PopFront() (int, error) {
	if d.isEmpty() {
		return 0, errors.New("error")
	}
	x := d.data[d.head]
	d.head = (d.head + 1) % d.capacity
	d.size--
	return x, nil
}

func (d *SliceDeque) PopBack() (int, error) {
	if d.isEmpty() {
		return 0, errors.New("error")
	}
	d.tail = (d.tail - 1 + d.capacity) % d.capacity
	x := d.data[d.tail]
	d.size--
	return x, nil
}

func (d *SliceDeque) isFull() bool {
	return d.size == len(d.data)
}

func (d *SliceDeque) isEmpty() bool {
	return d.size == 0
}

type Deque interface {
	PushFront(x int) error
	PushBack(x int) error
	PopFront() (int, error)
	PopBack() (int, error)
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
	sliceDeque := NewSliceDeque(dequeueCapacity)
	processCommands(inputData[2:], sliceDeque)
}

func processCommands(inputData []string, deque Deque) {
	for _, v := range inputData {
		command := strings.Split(v, " ")
		if command[0] == "push_back" {
			x, _ := strconv.Atoi(command[1])
			err := deque.PushBack(x)
			if err != nil {
				fmt.Println(err)
			}
		}
		if command[0] == "push_front" {
			x, _ := strconv.Atoi(command[1])
			err := deque.PushFront(x)
			if err != nil {
				fmt.Println(err)
			}
		}
		if command[0] == "pop_front" {
			x, err := deque.PopFront()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(x)
			}
		}
		if command[0] == "pop_back" {
			x, err := deque.PopBack()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(x)
			}
		}
	}
}
