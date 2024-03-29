package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/24414/run-report/106193533/

/**
  -- ПРИНЦИП РАБОТЫ --
  Создается структура hashtable, в которой содержится размер таблицы и слайс указателей на элементы (слайс "корзин").
  По условию задачи ключом может быть только целое число в диапазоне (1 ≤ n ≤ 10^9)
  Хэш функцией в таком случае выберем остаток целочисленного деления на размер таблицы, при этом размер таблицы возьмем 10.
  Таким образом мы гарантируем что значение хэш функции будет возвращать число от 0 до 9 включительно в зависимости от последнего знака в ключе.
  При этом с высокой вероятностью заполнятся все корзины таблицы, при этом мы можем гарантировать, что не будет 100% вероятности пустых корзин.
  Таким образом эффективно аллоцируем память под слайс корзин.

  Хранение данных в таблице организуем в виде двусвязного списка. С его же помощью будем разрешать коллизии (метод цепочек).

  Операции таблицы:
	1) put key value - в лучшем случае O(1) в худшем случаем O(n) по числу элементов в корзине
		- по хэш значению определяем номер корзины
        - если корзина пустая - добавляем новый элемент за O(1)
		- если корзина не пустая - внутренним методом итерируемся по списку элементов за O(n):
		  если находим элемент, совпадающий по ключу - заменяем его.
		  если не находим - добавляем новый элемент за O(1) в корзину, предварительно скорректировав ссылки у нового и текущего элемента в корзине

		Корректность:
		Хэш функция математически вернет значение от 0 до 9 включительно, мы не сможем выйти за границы слайса корзин.
		Если корзина пустая - мы добавим элемент в неё без изменений. Найти элемент сможем по хэш функции ключа, т.к. она детерминирована.
		Если корзина не пустая - мы точно по равенству ключа можем определить - есть ли элемент с таким же ключом в цепочке.
		Если элемент есть - мы заменяем в нем value, сохраняем принцип работы хэш-таблицы (замена значения).
		Если элемента нет - добавляем новый элемент в корзину по хэш функции и корректируем ссылки у нового и текущего элемента в корзине

	2) get key - в лучшем случае O(1) в худшем случаем O(n) по числу элементов в корзине
		- по ключу определяем хэш и номер корзины
		- если корзина пустая - возвращаем zero-value и флаг false
		- если корзина не пустая - внутренним методом итерируемся по списку элементов за O(n):
		  если находим элемент, совпадающий по ключу - возвращаем его и флаг true

		Корректность:
		Хэш функция математически вернет значение от 0 до 9 включительно, мы не сможем выйти за границы слайса корзин.
		Т.к. хэш функция детерминирована, то по ключу мы можем определить корзину.
		Если корзина не пустая - мы точно по равенству ключа можем определить - есть ли элемент с таким же ключом в цепочке.

	2) delete key - в лучшем случае O(1) в худшем случаем O(n) по числу элементов в корзине
		- по ключу определяем хэш и номер корзины
		- если корзина пустая - возвращаем zero-value и флаг false
		- если корзина не пустая - внутренним методом итерируемся по списку элементов за O(n):
		   если находим элемент, совпадающий по ключу - возвращаем его и флаг true
           корректируем ссылки у удаляемого, предыдущего и следующего элемента в корзине (GC уничтожит объект на который не будет ссылок)

		Корректность:
		Хэш функция математически вернет значение от 0 до 9 включительно, мы не сможем выйти за границы слайса корзин.
		Т.к. хэш функция детерминирована, то по ключу мы можем определить корзину.
		Если корзина не пустая - мы точно по равенству ключа можем определить - есть ли элемент с таким же ключом в цепочке.
		Удаление элемента по сути являются удалением ссылок на него и ссылок элемента на другие элементы.
		Если удаляемый элемент не содержал ссылку на предыдущий элемент - мы перемещаем следующий элементу корзину.
		В корзине всегда будет находиться элемент (если она не пустая полностью).

  -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  В общем случае O(1) при работе с методами таблицы.
  Общая сложность алгоритма составит O(n), где n - кол-во команд.

  При росте числа коллизий сложность методов таблицы в данной реализации деградирует до значений близких к O(n) где n кол-во элементов в корзине.
  В сравнении с хэш таблицей из стандартной библиотеки мы не можем игнорировать фактор коллизий, т.к.
  в данной реализации не предусмотрено расширение таблицы при переполнении.
  Таким образом у нас будет постоянное кол-во корзин, а значит при наполнении таблицы большим кол-вом данных
  неизбежны коллизии и деградация временной сложности операций к O(n).

  -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  O(n) элементов в корзинах т.к. потребность в памяти увеличивается линейно по мере заполнения таблицы (увеличения элементов в корзинах).
  При этом в силу способа разрешения коллизий дополнительная память потребуется для хранения указателей на следующие и предыдущие элементы
*/

/*
 * Т.к. поддерживать рехеширование и масштабирование хеш-таблицы не требуется - можно сразу задать размер таблицы
 */
type HashTable struct {
	size    int
	buckets []*node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*node, size),
	}
}

/*
 * Получаем value по ключу. Если value не найдено - возвращаем zero-value и флаг false.
 * Если ключ не найден (бакет пустой) - сложность O(1)
 * Если ключ найден O(n) в худшем случае (итерируемся по двусвязному списку), O(1) если в бакете только одно значение node
 */
func (h *HashTable) Get(key int) (int, bool) {
	node := h.getNode(key)
	if node == nil {
		return 0, false
	}

	return node.value, true
}

/*
 * Ищем value по ключу. Если value не найдено - возвращаем zero-value и флаг false.
 * Если значение найдено - возвращаем value при этом удаляем элемент.
 * Если ключ не найден (бакет пустой) - сложность O(1)
 * Если ключ найден O(n) в худшем случае (итерируемся по двусвязному списку), O(1) если в бакете только одно значение node
 */
func (h *HashTable) Delete(key int) (int, bool) {
	hash := h.hash(key)
	node := h.getNode(key)
	value := 0
	ok := false
	if node == nil {
		return 0, false
	}

	value, ok = node.value, true

	nodeAfterDelete := h.deleteNode(node)
	h.updateHeadIfNeeded(nodeAfterDelete, hash)

	return value, ok
}

/*
 * Если бакет по хэшу пустой - O(1) добавляем в него значение
 * Если в бакете существует n node - О(n) т.к. итерируемся по двусвязному списку для поиска по ключу.
 * Если не удалось обновить value (бакет не пустой и нет совпадений по ключу) - записываем в ячейку новую node O(1), обновляем ссылки
 */
func (h *HashTable) Put(key, value int) {
	hash := h.hash(key)
	head := h.buckets[hash]
	node := h.newNode(key, value)

	if head == nil {
		h.buckets[hash] = node
		return
	}

	existingNode := h.findByKey(key, head)
	if existingNode != nil {
		existingNode.value = value
		return
	}

	node.next = head
	head.prev = node
	h.buckets[hash] = node
}

/*
 * Обновляем ссылки у prev и next node относительно удаляемой node
 */
func (h *HashTable) deleteNode(node *node) *node {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	return node
}

/*
 * Если удаляемая node была первой в бакете - обновляем первое значение node в бакете
 */
func (h *HashTable) updateHeadIfNeeded(node *node, hash int) {
	if node.prev == nil {
		h.buckets[hash] = node.next
	}
	node.prev = nil
	node.next = nil
}

/*
 * Получаем хэш ключа.
 * Смотрим по хэшу в бакеты, если вернулся nil - прекращаем поиск, значение отсутствует.
 * Если в бакете есть значение - итерируемся по значениям для поиска совпадающего ключа.
 */
func (h *HashTable) getNode(key int) *node {
	hash := h.hash(key)
	node := h.buckets[hash]

	if node == nil {
		return nil
	}

	return h.findByKey(key, node)
}

/*
 * Итерируемся по двусвязному списку до совпадения ключа
 * Если совпадение получено - возвращаем node
 */
func (h *HashTable) findByKey(key int, head *node) *node {
	for head != nil {
		if head.key == key {
			return head
		}
		head = head.next
	}

	return nil
}

func (h *HashTable) hash(key int) int {
	if key < 0 {
		return -key % h.size
	}

	return key % h.size
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func (h *HashTable) newNode(key, value int) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func main() {
	data := getInputData()

	n, _ := strconv.Atoi(data[0])
	size := 10
	if n > 100 {
		size = n / 10
	}

	hashTable := NewHashTable(size)

	var answer strings.Builder
	for i := 1; i < len(data); i++ {
		command := strings.Split(data[i], " ")
		switch command[0] {
		case "get":
			key, _ := strconv.Atoi(command[1])
			value, ok := hashTable.Get(key)
			if !ok {
				answer.WriteString("None\n")
			} else {
				answer.WriteString(fmt.Sprintf("%v\n", value))
			}
		case "put":
			key, _ := strconv.Atoi(command[1])
			value, _ := strconv.Atoi(command[2])
			hashTable.Put(key, value)
		case "delete":
			key, _ := strconv.Atoi(command[1])
			value, ok := hashTable.Delete(key)
			if !ok {
				answer.WriteString("None\n")
			} else {
				answer.WriteString(fmt.Sprintf("%v\n", value))
			}
		}
	}
	fmt.Println(answer.String())
}

func getInputData() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var s []string
	for scanner.Scan() {
		bufStr := scanner.Text()
		s = append(s, bufStr)
	}

	return s
}
