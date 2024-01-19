package main

//https://contest.yandex.ru/contest/23815/run-report/105262261/

/*
-- ПРИНЦИП РАБОТЫ --
Сломанный массив фактически все равно остается отсортированным массивом, поэтому к нему применим принцип алгоритма бинарного поиска.
Разница с классическим алгоритмом бинарного поиска состоит в том, что нам необходимо делать проверки на то является ли
одна из части массива после разделения отсортированной.

-- РЕАЛИЗАЦИЯ --
Выбираем начальные, конечный и средний элементы массива.
Проверяем является ли левая часть массива отсортированной.

Если левая часть отсортирована и искомый элемент находится в ней - конечный элемент бин. поиска становится центральным, сдвинутым на 1 влево.
Далее поиск осуществляется в отсортированной части массива по классическому алгоритму бин. поиска.

Если левая часть отсортирована, но искомый элемент в ней не находится - начальный элемент поиска становится центральным, сдвинутым на 1 вправо.
Т.к. мы убедились, что в левой части искомого элемента быть не может (она отсортирована), мы можем отбросить эту часть массива и продолжить поиск в правой части.

Если левая часть не отсортирована - исходя из принципа сломанного массива - правая часть является отсортированной.

Проверяем находится ли искомый элемент в правой отсортированной части.
Если искомый элемент находится в правой отсортированной части - начальный элемент поиска становится центральным, сдвинутым на 1 вправо.
Далее поиск осуществляется в отсортированной части массива по классическому алгоритму бин. поиска.

Если правая часть оказалась отсортированной, но искомый элемент в ней отсутствует - конечный элемент поиска становится центральным, сдвинутым на 1 влево.
Т.к. мы убедились, что в правой части искомого элемента быть не может (она отсортирована), мы можем отбросить эту часть массива и продолжить поиск в левой части.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Сломанный массив из условия - это "сдвинутый" отсортированный по возрастанию массив.
Таким образом внутри массива у нас есть только один элемент, который больше предыдущего и меньше следующего.
Такой элемент ранее был конечным элементом отсортированного массива.
Фактически "сломанный массив" состоит из 2-х отсортированный по возрастанию массивов.

Мы можем применить стандартный подход бинарного поиска учитывая указанную особенность, т.к. при взятии
среднего значения могут быть лишь 3 следующих варианта:
- Левая часть массива отсортирована, правая нет.
  Понимаем есть ли искомый элемент в левой части и либо работаем с ней в стандартном бин поиске, либо полностью отбрасываем.
- Левая часть не отсортирована - значит правая часть отсортирована.
  Понимаем есть ли искомый элемент в правой части и либо работаем с ней в стандартном бин поиске, либо полностью отбрасываем.
- Левая и правая части отсортированы (мы попали в бывший конечный элемент).
  Так как искомый элемент находится в отсортированной части, то мы можем его найти в стандартном бин поиске.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Каждый раз мы делим массив на 2 равные части.
Временная сложность равна O(log(n)) - стандартному алгоритму бин. поиска.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
В ходе работы алгоритмы не создаются дополнительные промежуточные данные.
Пространственная сложность зависит от объёма входных данных (элементов добавленных в слайс),
соответственно будет занято n количество памяти. Отсюда можно сделать вывод, что пространственная сложность O(n).
*/

func brokenSearch(arr []int, k int) int {
	start := 0
	end := len(arr) - 1

	// Пока начальный индекс меньше или равен конечному
	for start <= end {
		mid := start + (end-start)/2 // Избегаем переполнения

		// Если нашли элемент, возвращаем его индекс
		if arr[mid] == k {
			return mid
		}

		// Определяем, отсортирована ли левая часть
		if arr[start] <= arr[mid] {
			// Если элемент находится в отсортированной левой части
			if k >= arr[start] && k < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			// Если отсортирована правая часть
			if k > arr[mid] && k <= arr[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	// Если элемент не найден, возвращаем -1
	return -1
}
