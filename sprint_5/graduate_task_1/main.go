package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/24810/run-report/108681156/

/*
-- ПРИНЦИП РАБОТЫ --
  Для реализации алгоритма - воспользуемся бинарной кучей.
  Бинарная куча дает гарантию, что на вершине будет наименьший или наибольший элемент, доступный за О(1).
  При этом операции добавления и удаления элемента в бинарной куче (с гарантией сохранения структуры кучи) - O(log n)

-- РЕАЛИЗАЦИЯ --
  Создадим структуру бинарной кучи, кучу построим на основе слайса.
  Элементы для целей сохранения структуры кучи (дерево) будем хранить по следующему принципу:
     Вершина кучи имеет индекс [1]
     Потомок вершины кучи (и каждого другого узла дерева) имеет индекс [2*i] и [2*i+1]

  Реализуем для кучи следующие экспортируемые методы:
   - Push - добавление элемента в кучу
     Добавляем элемент в конец слайса.
     Запускаем просеивание вверх для добавленного элемента (по индексу элементу):
		- определяем индекс родителя (индекс элемента делится на 2)
        - запускаем цикл с условием: если индекс = 0 - мы добавили вершину (или дошли до вершины), просеивание не требуется
			- по функции компаратору проверяем приоритетность
		 		если родитель > потомка - меняем местами, определяем заново индексы, продолжаем цикл
				если родитель <= потомка - просеивание не требуется, выходим из цикла

	- Pop - извлечение элемента с вершины кучи
	  Извлекаем элемент слайса по индексу 1 и возвращаем его.
      Заменяем вершину кучи на последний элемент слайса.
      Удаляем последний элемент слайса.
	  Запускаем просеивание вниз с вершины кучи:
		- определяем индексы потомков (индекс * 2 и индекс * 2 + 1)
		- запускаем цикл с условием: индекс левого потомка меньше длины слайса элементов:
			- определяем индекс большего из двух потомков
			- по функции компаратору проверяем приоритетность между потомками и выбираем наибольший
		    	если потомок > родителя - меняем местами, продолжаем цикл
				если потомок <= родителя - просеивание не требуется, выходим из цикла

  В конструктор кучи передадим набор элементов.
  Вызовем метод upSifting для каждого элемента в наборе, у которого есть потомки.
  Потомки гарантированно есть у первых heapSize/2 вершин.
  После добавления всех элементов будем последовательно вызывать метод Pop сохраняя результат.
  В итоге получим отсортированный набор данных.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Гарантии, даваемые бинарной кучей (методами Push и Pop), сохраняются во время работы алгоритма.
  После каждого добавления или удаления элемента на вершине кучи будет элемент с наибольшим приоритетом.
  Таким образом последовательно извлекая элементы с вершины кучи мы получим отсортированный набор данных.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Временная сложность алгоритма быстрой сортировки:
  - построение кучи: O(N)
  - построение сортированной последовательности: O(N * log N) - удаляем N элементов из кучи

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность алгоритма: O(1).
  Используется внутренний слайс кучи, заранее аллоцируем память под весь набор элементов.
  Дополнительная память не используется: не используем рекурсии или вспомогательные массивы.
*/

type Person struct {
	login   string
	solved  int
	penalty int
}

type PersonHeap struct {
	arr []Person
}

func New(persons []Person) *PersonHeap {
	arr := make([]Person, 1, len(persons)+1)
	arr = append(arr, persons...)

	p := &PersonHeap{
		arr: arr,
	}
	p.init()

	return p
}

// Append добавляем элемент в бинарную кучу
func (h *PersonHeap) Append(p Person) {
	// добавляем элемент в конец
	h.arr = append(h.arr, p)
	// запускаем просеивание вверх добавленного элемента
	h.upSifting(len(h.arr) - 1)
}

// upSifting просеивания вверх
func (h *PersonHeap) upSifting(index int) {

	// определяем индекс родителя
	// сравнение делаем только с родителем
	// в бинарной куче именно вершина "не менее" чем любой из потомков
	parentIndex := index / 2

	for parentIndex != 0 {
		if h.less(index, parentIndex) {
			h.swap(index, parentIndex)
			index = parentIndex
			parentIndex = index / 2
		} else {
			break
		}
	}
}

func (h *PersonHeap) init() {
	for i := len(h.arr) / 2; i > 0; i-- {
		h.downSifting(i)
	}
}

// Pop получение приоритетного элемента из кучи
func (h *PersonHeap) Pop() Person {
	// записываем вершину кучи в ответ
	p := h.arr[1]
	// заменяем вершину кучи на последний элемент
	h.arr[1] = h.arr[len(h.arr)-1]
	// удаляем последний элемент
	h.arr = slices.Delete(h.arr, len(h.arr)-1, len(h.arr))
	// запускаем просеивание вниз с вершины кучи
	h.downSifting(1)
	return p
}

// downSifting просеивания вниз
func (h *PersonHeap) downSifting(index int) {
	// индекс левого потомка
	left := index * 2
	// индекс правого потомка
	right := index*2 + 1

	for left < len(h.arr) {
		// определяем индекс большего из двух потомков
		largest := left
		if right < len(h.arr) && h.less(right, left) {
			largest = right
		}
		// проверяем необходимость сделать swap с большим потомком
		if h.less(largest, index) {
			h.swap(largest, index)
			index = largest
			left = index * 2
			right = index*2 + 1
		} else {
			break
		}
	}
}

// Компаратор для сравнения элементов в куче
func (h *PersonHeap) less(i, j int) bool {
	return !comparator(h.arr[i], h.arr[j])
}

// Меняем местами два элемента в куче
func (h *PersonHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func main() {
	participants := getInputData()

	personHeap := New(participants)

	sb := strings.Builder{}
	for i := 0; i < len(participants); i++ {
		sb.WriteString(personHeap.Pop().login + "\n")
	}

	fmt.Println(sb.String())
}

func comparator(a, b Person) bool {
	// сначала сравниваем по количеству задач
	if a.solved != b.solved {
		return a.solved < b.solved
	}

	// если равенство числа решённых задач
	// сравниваем по штрафам
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}

	// если равенство решенных задач и штрафов
	// сравниваем по лексике
	return strings.Compare(a.login, b.login) > 0
}

func getInputData() []Person {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	var bufStr string
	bufStr = scanner.Text()
	n, _ := strconv.Atoi(bufStr)

	users := make([]Person, n)
	var user Person
	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Text()
		strArr := strings.Split(bufStr, " ")
		user = Person{}
		user.login = strArr[0]
		user.solved, _ = strconv.Atoi(strArr[1])
		user.penalty, _ = strconv.Atoi(strArr[2])

		users[i] = user
	}

	return users
}
