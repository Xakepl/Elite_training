package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
	// return "TODO: реализуйте Summary"
}

func (t *Task) Complete() {
	// TODO: присвойте статус done исходной задаче.
	if t.Status != "done" {
		t.Status = "done"
	}
}

func (t *Task) SetPriority(priority int) bool {
	// TODO: проверьте диапазон 1–3, при ошибке объект не менять.
	if priority >= 1 && priority <= 3 {
		t.Priority = priority
		return true
	}
	return false
}

func findTaskByID(tasks []Task, id int) *Task {
	// TODO: верните &tasks[i] при совпадении, иначе nil.
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}
func tasksByStatus(tasks []Task, status string) []Task {
	var result []Task
	for i := range tasks {
		if tasks[i].Status == status {
			result = append(result, tasks[i])
		}
	}
	return result
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}
	// TODO: найдите ID=2, завершите и установите приоритет 3.
	// TODO: посчитайте завершённые задачи.
	for _, task := range tasks {
		fmt.Println(task.Summary())
	}

	fmt.Printf("Задача с ID = 99: ")
	if findTaskByID(tasks, 99) == nil {
		fmt.Printf("Задача не найдена")
	}
	fmt.Printf("\n")
	fmt.Println(tasks[1].Status, tasks[1].SetPriority(3), tasks[1].Summary())
	tasks[1].Complete()
	fmt.Println("Task 2 (in_progress -> " + tasks[1].Status + ")")
	mapDone := make(map[string]int)
	for _, task := range tasks {
		mapDone[task.Status]++
	}
	fmt.Println("Завершённые задачи:", mapDone["done"])

	fmt.Println("============================|DOP TASK|==============================")
	fmt.Printf("Задачи со статусом new: ")
	for _, task := range tasksByStatus(tasks, "new") {
		fmt.Println(task.Summary())
	}
	fmt.Printf("Задачи со статусом done: ")
	for _, task := range tasksByStatus(tasks, "done") {
		fmt.Println(task.Summary())
	}
	fmt.Printf("Задачи со статусом in_progress:")
	for _, task := range tasksByStatus(tasks, "in_progress") {
		fmt.Println(task.Summary())
	}
	if tasksByStatus(tasks, "in_progress") == nil {
		fmt.Printf("нет задач")
	}
}
