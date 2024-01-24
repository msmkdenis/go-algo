package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
В компании, где работает Тимофей, заботятся о досуге сотрудников и устраивают различные кружки по интересам.
Когда кто-то записывается на занятие, в лог вносится название кружка.

По записям в логе составьте список всех кружков, в которые ходит хотя бы один человек.
*/
type OrderedSet struct {
	data   map[string]struct{}
	order  []string
	Length int
}

func (o *OrderedSet) Add(key string) {
	if _, ok := o.data[key]; ok {
		o.data[key] = struct{}{}
		return
	}

	o.data[key] = struct{}{}
	o.order = append(o.order, key)
	o.Length++
}

func (o *OrderedSet) Get(i int) string {
	return o.order[i]
}

func main() {

	activitySet := OrderedSet{
		data:  make(map[string]struct{}),
		order: make([]string, 0),
	}

	const maxCapacity = 15 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	counter := 0
	for scanner.Scan() {
		if counter == 0 {
			counter++
			continue
		}
		line := scanner.Text()
		activitySet.Add(line)
		if line == "" {
			break
		}
	}

	for i := 0; i < activitySet.Length; i++ {
		fmt.Println(activitySet.Get(i))
	}
}
