package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Event представляет событие начала или конца сессии
type Event struct {
	time  int  // Время события
	start bool // Флаг начала или конца сессии
}

func main() {
	// Чтение данных
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	events := make([]Event, 0, 2*n)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		times := strings.Fields(line)
		start, _ := strconv.Atoi(times[0])
		end, _ := strconv.Atoi(times[1])
		events = append(events, Event{time: start, start: true})
		events = append(events, Event{time: end, start: false})
	}

	// Сортировка событий по времени, если времена равны, то начало идет перед окончанием
	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].start
		}
		return events[i].time < events[j].time
	})

	maxSessions := 0     // Максимальное число одновременных сессий
	currentSessions := 0 // Текущее число одновременных сессий
	bestTime := 0        // Время, на котором было максимальное число сессий

	// Обработка событий
	for _, event := range events {
		if event.start {
			currentSessions++
			// Обновляем максимальное число сессий и время, если текущее число сессий больше
			if currentSessions > maxSessions {
				maxSessions = currentSessions
				bestTime = event.time
			}
		} else {
			currentSessions--
		}
	}

	// Вывод результата
	fmt.Println(bestTime)
}
