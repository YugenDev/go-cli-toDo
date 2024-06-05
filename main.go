package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
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
		fmt.Print("> ")

		name, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTask(file, tasks)
		task.ListLastTask(tasks)

	case "delete":

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número")
			return
		}

		if len(os.Args) < 3 {
			fmt.Println("Proporciona un ID para eliminar")
			return
		}

		if !task.TaskExist(tasks, id) {
			fmt.Println("ID no existe")
			return
		}

		
		tasks = task.DeleteTask(tasks, id)
		task.SaveTask(file, tasks)
	}

}

func printUsage() {
	fmt.Println("Uso: go-cli-ToDo [list|add|complete|delete]")
}
