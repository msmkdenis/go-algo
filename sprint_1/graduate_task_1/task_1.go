package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	data := makeIntSlice(s[1], length)

	zeroCounter := 0
	distance := make([]int, length)
	for i, v := range data {
		if v == 0 {
			if zeroCounter == 0 {
				zeroCounter++
				k := 0
				for i >= 0 {
					distance[k] = i
					i--
					k++
				}
				rightDistance := 1
				for k < length {
					distance[k] = rightDistance
					k++
					rightDistance++
				}
			} else {
				zeroCounter++
				k := 0
				for i >= 0 {
					if distance[k] > i {
						distance[k] = i
					}
					i--
					k++
				}
				rightDistance := 1
				for k < length {
					distance[k] = rightDistance
					k++
					rightDistance++
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

func makeIntSlice(s string, len int) []int {
	result := make([]int, 0, len)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
