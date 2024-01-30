package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/24414/run-report/106188277/

/**
  -- ПРИНЦИП РАБОТЫ --
  Создаем индекс поиска в виде ассоциативного массива: ключ - уникальные слова из документов, значение - слайс структур вхождений слова в каждый документ (индекс документа - кол-во вхождений).
  Далее в цикле для каждого запроса: собираем множество уникальных слов из запроса.
  Итерируемся по словам из множества и собираем слайс структур вхождений слов из запроса в разные документы.
  Итерируемся по полученному слайсу и строим ассоциативный массив релевантности: ключ - индекс документа, значение кол-во вхождений слова в документ.
  На основе ассоциативного массива релевантности документов запросу строим слайс структур очков релевантности.
  В цикле до 5 раз выбираем наиболее релевантный документ из слайса, "обнуляя" найденный.

  -- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  В результате построения индекса получаем ассоциативный массив, в котором ключ - уникальные слова из документов, значение - слайс структур вхождений слова в каждый документ.
  В силу специфики устройства map (все ключи являются уникальными) мы уверены, что индекс составлен по всем словам документа без повторов.

  Далее для каждого запроса строим множества  слов на основе map, где значение пустая структура (идиоматичный способ построения set для go).
  Опять же в силу специфики устройства map (все ключи являются уникальными) мы уверены, что множество не содержит повторяющихся значений.

  На основе указанного выше индекса (по всем документам) и множества слов конкретного запроса рассчитываем релевантность запроса.
  Итерируемся по множеству слов и строим слайс структур вхождений слов запроса в документы.
  Т.к. индекс для каждого слова документа содержит корректный массив вхождений (указано выше благодаря устройству map), мы получим корректный массив вхождений для каждого слова запроса.

  Далее итерируясь по массиву вхождений составляем map релевантности, где ключ - номер документа, значение - кол-во вхождений.
  Благодаря устройству map мы можем гарантировать уникальность ключей (номеров документов), при это быстро посчитать релевантность.

  Далее нам необходимо выбрать 5 или менее документов с наибольшей релевантность.
  Составим массив структур из map релевантности и итерируясь по нему n раз (до 5 раз) каждый раз будем находить наиболее релевантный документ и затем "обнулять его",
  проставляя значения {-1, -1}, что гарантирует нам, что данный элемент однозначно не будет выбран наиболее релевантным в следующей итерации (индекс документа и кол-во вхождений не могут быть отрицательными).
  Совокупность гарантий устройства ассоциативного массива и способа отбора релевантных документов гарантируют корректность алгоритма.

  -- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Создание индекса - линейная сложность O(n * m) где n - кол-во документов, m - кол-во слов в документе
  Вывод ответа - i операций ниже, где i - кол-во запросов:
  		 O(n) - создание set уникальных слов из запроса (n - кол-во слов в запросе)
  		 O(s) - создание слайса структур вхождений слов из запроса в документы (s - кол-во слов в множестве)
  		 O(m) - создание map релевантности по уникальным словам (m - кол-во уникальных слов в запросе)
  		 O(k) - создание слайса структур с релевантностью для последующего поиска 5 "максимальных" элементов
  	     O(h * k) - поиск в слайсе структур релевантности до 5 самых больших вхождений (h <= 5)
  В общем случае временная сложность составит линейное время O((n * m) + (i * (n + m + k + h * k))) - сумма линейных операций выше.

  -- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  O(n) доп. памяти на построение индекса
  O(n + m + k) доп. памяти где n - кол-во уникальных слов, m - map релевантности и k - слайс структур с релевантностью документов для запроса
*/

const ResultLimit = 5

type occurrence struct {
	docIdx int
	count  int
}

type score struct {
	docIdx int
	weight int
}

/*
 * Строим поисковый индекс по документам.
 * Для каждого уникального слова строим map в которой ключ - слово, значение - слайс структур вхождений слова в каждый документ.
 * Получим индекс, в котором находится информация о всех уникальных словах и частоте вхождений
 * в тот или иной документ.
 */
func buildIndex(input []string, docCount int) map[string][]occurrence {
	var index = make(map[string][]occurrence)
	for i := 1; i < docCount+1; i++ {
		words := strings.Split(input[i], " ")
		for a := 0; a < len(words); a++ {
			occurrences, found := index[words[a]]
			if !found || occurrences[len(occurrences)-1].docIdx != i {
				index[words[a]] = append(occurrences, occurrence{i, 1})
			} else {
				occurrences[len(occurrences)-1].count++
			}
		}
	}

	return index
}

/*
 * Получаем set уникальных слов из соответствующего запроса
 */
func getUniqueWords(input string) map[string]struct{} {
	words := strings.Split(input, " ")
	uniqueWords := make(map[string]struct{})

	for a := 0; a < len(words); a++ {
		uniqueWords[words[a]] = struct{}{}
	}

	return uniqueWords
}

/*
 * Рассчитываем релевантность запроса относительно документа по index map, построенной ранее.
 */
func calcRelevance(uniqueWords map[string]struct{}, index map[string][]occurrence) []score {
	// Итерируемся по set уникальных слов из запроса
	// если слово встречается в индексе - заполняем слайс вхождений слов из запроса в документы
	var queryWords []occurrence
	for k, _ := range uniqueWords {
		if occurrences, found := index[k]; found {
			queryWords = append(queryWords, occurrences...)
		}
	}

	// Строим map релевантности, в которой ключ - индекс документа, значение - количество вхождений слов из запроса в документ
	scoresDocs := make(map[int]int)
	for _, q := range queryWords {
		scoresDocs[q.docIdx] += q.count
	}

	// Из полученной map релевантности собираем слайс структур score, которую в последствие сможем отсортировать
	var scores []score
	for k, v := range scoresDocs {
		scores = append(scores, score{k, v})
	}

	n := ResultLimit
	if len(scores) < ResultLimit {
		n = len(scores)
	}

	// Ищем 5 наиболее релевантных документов
	// Каждый найденный документ "обнуляем" для следующей итерации
	maxScores := make([]score, 0, n)
	k := 0
	for k < n {
		maxScore := score{
			docIdx: -1,
			weight: -1,
		}
		maxIdx := -1
		for i := 0; i < len(scores); i++ {
			if scores[i].weight > maxScore.weight || (scores[i].weight == maxScore.weight && scores[i].docIdx < maxScore.docIdx) {
				maxScore = scores[i]
				maxIdx = i
			}
		}
		maxScores = append(maxScores, maxScore)
		scores[maxIdx] = score{-1, -1}
		k++
	}

	return maxScores
}

/*
 * Готовим вывод с учетом ограничения на кол-во релевантных документов к выводу.
 */
func printResult(scores []score) {
	var p strings.Builder
	for k := 0; k < len(scores); k++ {
		if k == len(scores)-1 {
			p.WriteString(strconv.Itoa(scores[k].docIdx))
		} else {
			p.WriteString(strconv.Itoa(scores[k].docIdx) + " ")
		}
	}

	fmt.Println(p.String())
}

func main() {
	input := getInputData()
	docCount, _ := strconv.Atoi(input[0])

	//Строим индекс поиска в виде map по всем документам
	index := buildIndex(input, docCount)

	//Для каждого запроса вычисляем релевантность и выводим ответ
	for i := docCount + 2; i < len(input); i++ {
		uniqueWords := getUniqueWords(input[i])
		scores := calcRelevance(uniqueWords, index)
		printResult(scores)
	}
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
