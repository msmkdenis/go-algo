package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/run-report/102242196/

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

	distance := make([]int, length)
	if isZeroSlice { // проверка на граничный случай - если в слайсе только нули
		for i, _ := range data {
			distance[i] = 0
		}
	} else {
		// первичное составление слайса расстояний
		// двигаемся слева направо по слайсу адресов
		a := -1                  // используем -1 пока не встретим 0 (-1 расстояние невозможно - заменим)
		fromLeftZeroCounter := 0 // счетчик, чтобы понять, что встретили 0 и -1 не надо использовать
		fromLeftCounter := 0     // счетчик расстояний слева
		for i, v := range data {
			if v != 0 && fromLeftZeroCounter == 0 { // если не 0 и 0 не встречался ставим -1
				distance[i] = a
			} else if v == 0 { // если 0 - ставим 0 в расстояние и сбрасываем счетчик расстояний
				fromLeftZeroCounter++ // флаг встретившегося 0
				distance[i] = 0
				fromLeftCounter = 0
			} else {
				fromLeftCounter++ // увеличиваем счетчик расстояний справа от 0 и ставим его в слайс
				distance[i] = fromLeftCounter
			}
		}

		//идем справа налево по слайсу расстояний
		fromRightZeroCounter := 0
		fromRightCounter := 0
		for i := len(distance) - 1; i >= 0; i-- {
			if distance[i] != 0 && fromRightZeroCounter == 0 { // если не 0 и 0 не встречался пропускаем шаг (мы все еще не дошли до 0)
				continue
			} else if distance[i] == 0 { // если расстояние 0 - обновляем счетчик встреченных 0 и очищаем счетчик расстояний
				fromRightZeroCounter++
				fromRightCounter = 0
			} else {
				fromRightCounter++                                       // увеличиваем счетчик расстояний
				if distance[i] > fromRightCounter || distance[i] == -1 { // если расстояние больше текущего счетчика или равно -1 заменяем
					distance[i] = fromRightCounter
				}
			}
		}
	}

	// чтобы не печатать в цикле - съедает очень много ресурсов :D
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
