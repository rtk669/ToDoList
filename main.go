package main

import (
	"ToDoList/actions"
	"ToDoList/structures"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	list := map[string]*structures.List{}

	logs := []*structures.Logs{}

	for {
		fmt.Print("Введите комманду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("ошибка")
			return
		}

		text := scanner.Text()

		textSlice := strings.Fields(text)

		if len(textSlice) == 0 {
			fmt.Println()
			continue
		}
		cmd := textSlice[0]

		switch {
		case cmd == "add":
			isError := actions.AddTask(textSlice, list)

			structures.AddLog(&logs, textSlice, isError)

			continue

		case cmd == "delete":
			isError := actions.DeleteTask(textSlice, list)

			structures.AddLog(&logs, textSlice, isError)

			continue

		case cmd == "list":
			actions.GetTaskList(list)

			structures.AddLog(&logs, textSlice, "")

			continue

		case cmd == "done":
			isError := actions.CompleteTask(textSlice, list)

			structures.AddLog(&logs, textSlice, isError)

			continue

		case cmd == "log":
			isError := actions.GetLogs(logs)

			structures.AddLog(&logs, textSlice, isError)

			continue

		case cmd == "help":
			isError := actions.Help()

			structures.AddLog(&logs, textSlice, isError)

			continue

		case cmd == "exit":
			actions.Exit()

		default:
			fmt.Println("Вы ввели неизвестную команду")

			structures.AddLog(&logs, textSlice, "Вы ввели неизвестную команду")

			fmt.Println("")
		}

	}

}
