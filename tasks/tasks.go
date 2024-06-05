package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

var Reset = "\033[0m"
var Cyan = "\033[36m"

func ListTask(tasks []Task) {

	fmt.Println(" ")

	if (len(tasks)) == 0 {
		fmt.Println(Cyan + "No hay tareas" + Reset)
		return
	}

	for _, task := range tasks {

		status := ""
		if task.Complete {
			status = "✓"
		} else {
			status = " "
		}

		fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}
	fmt.Println(" ")
}

func AddTask(tasks []Task, name string) []Task {

	newTask := Task{
		ID:       getNextID(tasks),
		Name:     name,
		Complete: false,
	}

	return append(tasks, newTask)

}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writter := bufio.NewWriter(file)

	_, err = writter.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writter.Flush()
	if err != nil {
		panic(err)
	}
}

func getNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	} else {
		return tasks[len(tasks)-1].ID + 1
	}
}

func ListLastTask(tasks []Task) Task {
	fmt.Println("")
	fmt.Println(Cyan + "¡Se añadió la nueva tarea!" + Reset)
	lastTask := tasks[len(tasks)-1]
	fmt.Printf("-> [%s] %d %s\n", " ", lastTask.ID, lastTask.Name)
	return lastTask
}

func DeleteTask(tasks []Task, id int) []Task{
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func TaskExist(tasks []Task, id int) bool {
	for _, task := range tasks {
		if task.ID == id {
			return true
		}
	}
	return false
}