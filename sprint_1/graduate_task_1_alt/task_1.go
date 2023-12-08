package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/run-report/102337928/

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
	const maxCapacity = 10 * 1024 * 1024
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
	length, _ := strconv.Atoi(inputData[0])
	homeAddresses := makeIntSlice(inputData[1], length)

	distance := make([]int, length)

	counter := len(homeAddresses)
	for i, v := range homeAddresses {
		if v == 0 {
			counter = 0
		} else {
			counter++
			distance[i] = counter
		}
	}

	counter = len(homeAddresses)
	for i := len(homeAddresses) - 1; i >= 0; i-- {
		if homeAddresses[i] == 0 {
			counter = 0
		} else {
			counter++
			if distance[i] > counter {
				distance[i] = counter
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

func makeIntSlice(s string, len int) []int {
	result := make([]int, 0, len)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
