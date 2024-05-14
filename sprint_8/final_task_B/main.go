package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/26133/run-report/114014665/

/*
-- ПРИНЦИП РАБОТЫ --
  Алгоритм проверки возможности разбить строку T на слова из списка words основан на использовании префиксного дерева (Trie) и динамического программирования.
  Мы создаем Trie для быстрого поиска слов и их префиксов, что позволяет эффективно проверять подстроки строки T.
  Затем используем динамическое программирование для определения, можно ли разбить строку T на допустимые слова.

-- РЕАЛИЗАЦИЯ --
  1. Создаем Trie и вставляем в него все слова из списка words.
  2. Создаем булевый массив dp размером len(T)+1 и инициализируем dp[0] как true.
     dp[i] будет означать, что подстрока T[0:i] может быть разбита на слова из словаря.
  3. Проходим по каждому символу строки T. Если dp[i] равно true, проверяем все возможные подстроки, начиная с позиции i,
     используя Trie для поиска соответствующих слов. Если подстрока найдена в Trie, устанавливаем dp[j+1] в true.
  4. После завершения проверки, если dp[len(T)] равно true, значит строку T можно разбить на слова из списка words, иначе - нет.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Алгоритм гарантирует корректность, так как проверяет все возможные разбиения строки T, используя эффективное хранение и поиск слов в Trie.
  Использование динамического программирования позволяет систематически проверять все возможные подстроки, обеспечивая нахождение верного решения.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Временная сложность алгоритма определяется длиной строки T и количеством слов в словаре.
  В худшем случае, временная сложность составляет O(n*m*l), где
  n - длина строки T,
  m - количество слов в словаре,
  l - средняя длина слова.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность определяется размером Trie и массивом dp, что составляет O(n + k), где
  n - длина строки T,
  k - общее количество символов во всех словах словаря.
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
