package tasks

import "fmt"

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTask(tasks []Task) {
	if (len(tasks)) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {

		status := ""
		if task.Complete {
			status = "✓"
		} else {
			status = "X"
		}

		fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}
}

func AddTask(tasks []Task, name string) []Task {
	newTask := task{
		ID: 10,
		Name: name,
		Complete: false,
	}

	return append(tasks, newTask)
}
