package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type person struct {
	login   string
	solved  int
	penalty int
}

func main() {
	const maxCapacity = 150 * 1024 * 1024
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

	capacity, _ := strconv.Atoi(inputData[0])
	participants := makeParticipants(inputData[1:], capacity)

	sort.SliceStable(participants, func(i, j int) bool {
		if participants[i].solved == participants[j].solved {
			if participants[i].penalty == participants[j].penalty {
				return participants[i].login < participants[j].login
			}
			return participants[i].penalty < participants[j].penalty
		}
		return participants[i].solved > participants[j].solved
	})

	for _, p := range participants {
		fmt.Println(p.login)
	}
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
