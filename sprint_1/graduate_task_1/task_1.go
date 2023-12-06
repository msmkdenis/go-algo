package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/run-report/102130766/

/*
Тимофей ищет место, чтобы построить себе дом.
Улица, на которой он хочет жить, имеет длину n, то есть состоит из n одинаковых идущих подряд участков.
Каждый участок либо пустой, либо на нём уже построен дом.

Общительный Тимофей не хочет жить далеко от других людей на этой улице.
Поэтому ему важно для каждого участка знать расстояние до ближайшего пустого участка.
Если участок пустой, эта величина будет равна нулю — расстояние до самого себя.

Помогите Тимофею посчитать искомые расстояния.
Для этого у вас есть карта улицы.
Дома в городе Тимофея нумеровались в том порядке, в котором строились, поэтому их номера на карте никак не упорядочены.
Пустые участки обозначены нулями.
*/

func main() {
	const maxCapacity = 10240 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}
	length, _ := strconv.Atoi(s[0])
	data, isZeroSlice := makeIntSlice(s[1], length)

	/*
		Основная идея алгоритма:
		1. Составить карту нулей (записывать индексы) в слайс zeroIndexes
		2. Составлять расстояния с учетом карты нулей, избегая итерирования по всему слайсу с номерами домов.
	*/

	zeroIndexes := make([]int, 0) // составляем карту нулей (записываем индексы)
	for i, v := range data {
		if v == 0 {
			zeroIndexes = append(zeroIndexes, i)
		}
	}

	distance := make([]int, length)
	if isZeroSlice { // проверка на граничный случай - если в слайсе только нули
		for i, _ := range data {
			distance[i] = 0
		}
	} else {
		zeroCounter := 0 // счетчик на первое вхождение, чтобы не проверять расстояния слева
		previousZero := 0
		for i, v := range data {
			if v == 0 { // двигаемся по слайсу с номерами домов, ищем вхождение нуля
				if zeroCounter == 0 { // если это первое вхождение нуля среди номеров домов
					zeroCounter++
					k := 0
					previousZero = i // запоминаем индекс нуля для последующих вычислений
					for i >= 0 {     // двигаемся слева направо, заполняя расстояния
						distance[k] = i // если индекс 0 например 3, то расстояние от края будет 3
						i--             // в цикле уменьшаем расстояние от края до 0
						k++             // увеличиваем индекс в слайсе расстояний
					}
					rightDistance := 1
					for k < length { // заполняем расстояния справа от нуля до конца слайса расстояний
						distance[k] = rightDistance
						k++
						rightDistance++
					}
				} else { // если это не первое вхождение нуля
					zeroCounter++
					k := previousZero     // ведём отсчет от предыдущего вхождения нуля, т.к. мы запоминали индекс ранее
					l := i - previousZero // расстояние от предыдущего вхождения нуля до текущего
					for l >= 0 {
						if distance[k] > l { // в определенный момент расстояние до текущего нуля будет меньше расстояния до предыдущего
							distance[k] = l
						}
						l--
						k++
					}
					index := zeroCounter               // индекс текущего вхождения нуля
					if index+1 <= len(zeroIndexes)-1 { // проверяем, что у нас не граничное значение и справа еще есть нули
						rightDistance := 1
						for k < zeroIndexes[index+1] { // заполняем расстояния справа от нуля до следующего вхождения нуля
							distance[k] = rightDistance
							k++
							rightDistance++
						}
					} else { // если нет следующего вхождения нуля, то заполняем расстояния до конца слайса с номерами домов
						rightDistance := 1
						for k < length {
							distance[k] = rightDistance
							k++
							rightDistance++
						}
					}
					previousZero = i // запоминаем индекс нуля для последующих вычислений
				}
			}
		}
	}

	var res strings.Builder
	for n, v := range distance {
		if n == len(distance)-1 {
			res.WriteString(strconv.Itoa(v))
		} else {
			res.WriteString(strconv.Itoa(v) + " ")
		}
	}

	fmt.Println(res.String())
}

func makeIntSlice(s string, len int) ([]int, bool) {
	result := make([]int, 0, len)
	isZeroSlice := true
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		if i != 0 {
			isZeroSlice = false
		}
		result = append(result, i)
	}
	return result, isZeroSlice
}
