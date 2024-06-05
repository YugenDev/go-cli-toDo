package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	task "github.com/YugenDev/go-cli-toDo/tasks"
)

func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}

	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "list":
		task.ListTask(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("¿Que tarea quieres agregar?")

		name, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTask(file, tasks)
	}

}

func printUsage() {
	fmt.Println("Uso: go-cli-ToDo [list|add|complete|delete]")
}
