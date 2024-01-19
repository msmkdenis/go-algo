package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type person struct {
	login   string
	solved  int
	penalty int
}

func main() {
	//const maxCapacity = 13 * 1024 * 1024
	//buf := make([]byte, maxCapacity)
	//
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Buffer(buf, maxCapacity)
	//
	//var inputData []string
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	inputData = append(inputData, line)
	//	if line == "" {
	//		break
	//	}
	//}
	//capacity, _ := strconv.Atoi(inputData[0])

	//participants := makeParticipants(inputData[1:], capacity)

	//sort.SliceStable(participants, func(i, j int) bool {
	//	if participants[i].solved == participants[j].solved {
	//		if participants[i].penalty == participants[j].penalty {
	//			return participants[i].login < participants[j].login
	//		}
	//		return participants[i].penalty < participants[j].penalty
	//	}
	//	return participants[i].solved > participants[j].solved
	//})

	participants, err := getInputData()
	if err != nil {
		showError(err)
	}

	if len(participants) > 1 {
		// сортируем
		//quickSort(participants, 0, len(participants)-1)

		sort.Slice(participants, func(i, j int) bool {
			if participants[i].solved == participants[j].solved {
				if participants[i].penalty == participants[j].penalty {
					return participants[i].login < participants[j].login
				}
				return participants[i].penalty < participants[j].penalty
			}
			return participants[i].solved > participants[j].solved
		})
	}

	//quickSort(participants, 0, len(participants)-1)

	for _, p := range participants {
		fmt.Println(p.login)
	}
}

func getInputData() (users []person, err error) {
	var input *os.File
	var bufStr string

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	bufStr = scanner.Text()
	n, err := strconv.Atoi(bufStr)
	if err != nil {
		return
	}

	users = make([]person, n)
	var user person
	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Text()

		strArr := strings.Split(bufStr, " ")
		user = person{
			login: strArr[0],
		}

		if strArr[1] != "0" {
			user.solved, err = strconv.Atoi(strArr[1])
		}
		if strArr[2] != "0" {
			user.penalty, err = strconv.Atoi(strArr[2])
		}

		users[i] = user
	}

	return
}

// getInputFromFile получение ввода из файла
func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// showError вывод ошибки
func showError(err interface{}) {
	panic(err)
}

func makeParticipants(inputData []string, capacity int) []person {
	participants := make([]person, 0, capacity)
	for _, v := range inputData {
		var p person
		_, _ = fmt.Sscanf(v, "%s %d %d", &p.login, &p.solved, &p.penalty)
		participants = append(participants, p)
	}
	return participants
}

// quickSort сортировка слайса
func quickSort(arr []person, left, right int) {

	// базовый случай
	if left >= right {
		return
	}

	pivot := sortPart(arr, left, right, comparator)

	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

// sortPart точка опоры и сортировка части слайса
func sortPart(arr []person, left, right int, compare func(a, b person) bool) int {
	p := arr[left]
	l := left + 1
	r := right

	for {
		for l <= r && compare(arr[r], p) {
			r--
		}

		for l <= r && !compare(arr[l], p) {
			l++
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
		} else {
			break
		}
	}

	arr[left], arr[r] = arr[r], arr[left]
	return r
}

// comparator сравнение значений
func comparator(a, b person) bool {

	// если равенство числа решённых задач
	if a.solved == b.solved {

		// если и в штрафах равенство - выводим по лексике
		if a.penalty == b.penalty {
			if strings.Compare(a.login, b.login) > 0 {
				return true
			}
			return false
		}

		// сравним по штрафам
		return a.penalty > b.penalty
	}

	// сравним по решенным
	return a.solved < b.solved
}
