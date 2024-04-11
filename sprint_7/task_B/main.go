package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Lesson struct {
	start, end int
}

func main() {
	inputData := getInputData()
	n, _ := strconv.Atoi(inputData[0])

	lessons := make([]Lesson, n)
	for i := 0; i < n; i++ {
		parts := strings.Fields(inputData[i+1])
		lessons[i] = Lesson{parseTime(parts[0]), parseTime(parts[1])}
	}

	sort.Slice(lessons, func(i, j int) bool {
		if lessons[i].end == lessons[j].end {
			return lessons[i].start < lessons[j].start
		}
		return lessons[i].end < lessons[j].end
	})

	schedule := make([]Lesson, 0)
	lastEnd := 0
	for _, lesson := range lessons {
		if lesson.start >= lastEnd {
			schedule = append(schedule, lesson)
			lastEnd = lesson.end
		}
	}

	fmt.Println(len(schedule))
	for _, lesson := range schedule {
		startHour := lesson.start / 60
		startMinute := lesson.start % 60
		endHour := lesson.end / 60
		endMinute := lesson.end % 60
		startStr := fmt.Sprintf("%d", startHour)
		if startMinute != 0 {
			startStr += fmt.Sprintf(".%02d", startMinute)
		}
		endStr := fmt.Sprintf("%d", endHour)
		if endMinute != 0 {
			endStr += fmt.Sprintf(".%02d", endMinute)
		}
		fmt.Printf("%s %s\n", startStr, endStr)
	}
}

func parseTime(timeStr string) int {
	parts := strings.Split(timeStr, ".")
	hours, _ := strconv.Atoi(parts[0])
	minutes := 0
	if len(parts) > 1 {
		minutes, _ = strconv.Atoi(parts[1])
	}
	return hours*60 + minutes
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
