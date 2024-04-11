package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type GoldHeap struct {
	costPerKg, weight int
}

func main() {
	inputData := getInputData()

	M, _ := strconv.Atoi(inputData[0])
	n, _ := strconv.Atoi(inputData[1])

	goldHeaps := make([]GoldHeap, n)
	for i := 0; i < n; i++ {
		parts := strings.Fields(inputData[i+2])
		cost, _ := strconv.Atoi(parts[0])
		weight, _ := strconv.Atoi(parts[1])
		goldHeaps[i] = GoldHeap{cost, weight}
	}

	// Сортируем кучи по стоимости за килограмм в порядке убывания
	sort.Slice(goldHeaps, func(i, j int) bool {
		return goldHeaps[i].costPerKg > goldHeaps[j].costPerKg
	})

	totalCost := 0
	remainingCapacity := M

	// Проходим по отсортированным кучам и добавляем их в рюкзак,
	// пока у нас есть место
	for _, heap := range goldHeaps {
		if remainingCapacity >= heap.weight {
			totalCost += heap.costPerKg * heap.weight
			remainingCapacity -= heap.weight
		} else {
			totalCost += heap.costPerKg * remainingCapacity
			break
		}
	}

	fmt.Println(totalCost)
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
