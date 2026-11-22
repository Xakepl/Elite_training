package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая задача"
	case "in_progress":
		return "В работе"
	case "done":
		return "Выполнено"
	default:
		return "Неизвестный статус"
	}
	// TODO: обработайте new, in_progress, done через switch.
	// TODO: предусмотрите ветку default для незнакомого статуса.
}

func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	count := 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	// TODO: пройдите по срезу через for range.
	// TODO: сравните каждый элемент с target и верните счётчик.
	return count
}

func completionPercent(statuses []string) float64 {
	doneCount := countStatus(statuses, "done")
	if len(statuses) == 0 {
		fmt.Println("Нет данных!")
		return 0
	}
	return float64(doneCount) / float64(len(statuses)) * 100
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	// TODO: добавьте ещё один статус new через append.
	statuses = append(statuses, "new")
	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))
	mapStatus := make(map[string]int)
	for _, status := range statuses {
		mapStatus[status]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d \n", mapStatus["new"], mapStatus["in_progress"], mapStatus["done"])
	// TODO: создайте map[string]int для подсчёта статусов.
	// TODO: выведите new, in_progress, done в фиксированном порядке.
	fmt.Printf("blocked: %s \n", statusLabel("blocked"))

	statuses0 := []string{}
	mapStatus0 := make(map[string]int)
	for _, status := range statuses0 {
		mapStatus0[status]++
	}
	fmt.Printf("Статистика statuses0: new=%d, in_progress=%d, done=%d \n", mapStatus0["new"], mapStatus0["in_progress"], mapStatus0["done"])
	statuses2 := []string{"done"}
	mapStatus2 := make(map[string]int)
	for _, status := range statuses2 {
		mapStatus2[status]++
	}
	fmt.Printf("Статистика statuses2: new=%d, in_progress=%d, done=%d \n", mapStatus2["new"], mapStatus2["in_progress"], mapStatus2["done"])
	statuses3 := []string{"new", "new"}
	mapStatus3 := make(map[string]int)
	for _, status := range statuses3 {
		mapStatus3[status]++
	}
	fmt.Printf("Статистика statuses3: new=%d, in_progress=%d, done=%d \n", mapStatus3["new"], mapStatus3["in_progress"], mapStatus3["done"])

	statuses4 := []string{"done", "done"}
	mapStatus4 := make(map[string]int)
	for _, status := range statuses4 {
		mapStatus4[status]++
	}
	fmt.Printf("Статистика statuses4: new=%d, in_progress=%d, done=%d \n", mapStatus4["new"], mapStatus4["in_progress"], mapStatus4["done"])

	statuses5 := []string{""}
	mapStatus5 := make(map[string]int)
	for _, status := range statuses5 {
		mapStatus5[status]++
	}
	fmt.Printf("Статистика statuses5: new=%d, in_progress=%d, done=%d \n", mapStatus5["new"], mapStatus5["in_progress"], mapStatus5["done"])
	fmt.Printf("=======================|DOP TASK|========================== \n")
	fmt.Printf("A\n")
	fmt.Println(findStatus(statuses, "done"))
	fmt.Println(findStatus(statuses, "new"))
	fmt.Println(findStatus(statuses, "blocked"))
	s := []string{}
	fmt.Println(findStatus(s, "done"))
	fmt.Printf("B\n")
	fmt.Printf("Процент завершения: %.2f%%\n", completionPercent(statuses))
	statusesEmpty := []string{}
	fmt.Printf("Процент завершения пустого среза: %.2f%%\n", completionPercent(statusesEmpty))
	statusesAllDone := []string{"done", "done", "done"}
	fmt.Printf("Процент завершения всех done: %.2f%%\n", completionPercent(statusesAllDone))
	statusesAllNew := []string{"new", "new", "new"}
	fmt.Printf("Процент завершения всех new: %.2f%%\n", completionPercent(statusesAllNew))
	statuses50 := []string{"new", "done", "done", "new"}
	fmt.Printf("Процент завершения 50/50: %.2f%%\n", completionPercent(statuses50))
}
