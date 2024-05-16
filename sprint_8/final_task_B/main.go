package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/26133/run-report/114169301/

/*
-- ПРИНЦИП РАБОТЫ --
  Алгоритм проверки возможности разбить строку T на слова из списка words основан на использовании префиксного дерева (Trie) и динамического программирования.
  Мы создаем Trie для быстрого поиска слов и их префиксов, что позволяет эффективно проверять подстроки строки T.
  Затем используем динамическое программирование для определения, можно ли разбить строку T на допустимые слова.

-- РЕАЛИЗАЦИЯ --
  1. Создаем Trie и вставляем в него все слова из списка words.
  2. Создаем булевый массив dp размером len(T) + 1 и инициализируем dp[0] как true.
     dp[i] будет означать, что подстрока T[0:i] может быть разбита на слова из словаря.
  3. Проходим по каждому символу строки T. Если dp[i] равно true, проверяем все возможные подстроки, начиная с позиции i,
     используя Trie для поиска соответствующих слов. Если подстрока найдена в Trie, устанавливаем dp[j+1] в true.
  4. После завершения проверки, если dp[len(T)] равно true, значит строку T можно разбить на слова из списка words, иначе - нет.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Алгоритм гарантирует корректность, так как проверяет все возможные разбиения строки T, используя эффективное хранение и поиск слов в Trie.
  Использование динамического программирования позволяет систематически проверять все возможные подстроки, обеспечивая нахождение верного решения.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Построение Trie:
  Вставка каждого слова в Trie занимает O(l) времени, где l - длина слова.
  Если у нас m слов, и l - их средняя длина, то общее время на построение Trie составляет O(m * l).

  Заполнение массива dp:
  Внешний цикл проходит по каждому символу строки T (n символов).
  Внутренний цикл проходит по подстрокам начиная с текущей позиции i, и каждый раз мы проверяем символ в Trie.
  В худшем случае, каждый символ строки T проверяется на наличие в Trie, что занимает O(l) времени, где l - максимальная длина слова в Trie.
  Таким образом, для каждой позиции i внутренний цикл выполняется максимум за O(l).
  Следовательно, общая временная сложность алгоритма будет O(n * l) для заполнения dp.

  В итоге, общая временная сложность алгоритма составляет O(m * l + n * l).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность алгоритма определяется объемом памяти,
  необходимым для хранения Trie, динамического массива dp и строки T.
  Trie: В худшем случае, Trie будет содержать все символы всех слов из списка words.
  Если m - количество слов, а l - средняя длина слова, то
  пространственная сложность для Trie будет O(m * l).
  Массив dp: Булевый массив dp имеет длину len(T) + 1, что дает пространственную сложность O(n), где n - длина строки T.
  В совокупности, пространственная сложность алгоритма составляет O(n + m * l).
*/

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search проверяет, есть ли слово в Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith проверяет, есть ли в Trie слово, начинающееся с данного префикса
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// canSegmentString проверяет, можно ли разбить строку T на слова из words
func canSegmentString(T string, words []string) bool {
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	// Создаем dp массив и инициализируем dp[0] в true
	dp := make([]bool, len(T)+1)
	dp[0] = true

	// Проверяем каждую позицию в строке T
	for i := 0; i < len(T); i++ {
		if dp[i] {
			node := trie.root
			for j := i; j < len(T); j++ {
				ch := rune(T[j])
				if _, exists := node.children[ch]; !exists {
					break
				}
				node = node.children[ch]
				if node.isEnd {
					dp[j+1] = true
				}
			}
		}
	}

	// Возвращаем результат для всей строки
	return dp[len(T)]
}

func main() {
	data := getInputData()

	// Читаем строку T
	T := data[0]

	// Читаем количество слов n
	n, _ := strconv.Atoi(data[1])

	words := data[2 : 2+n]

	if canSegmentString(T, words) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func getInputData() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	const maxCapacity = 10240 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		bufStr := scanner.Text()
		s = append(s, bufStr)
	}

	return s
}
