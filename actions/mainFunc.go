package actions

import (
	"ToDoList/structures"
	"fmt"
	"os"
	"time"
)

func Help() string {
	fmt.Println()
	fmt.Println("Список доступных команд:")
	fmt.Println("add {заголовок задачи} {описание задачи} - позволяет добавлять задачу в список дел")
	fmt.Println("delete {заголовок задачи} - позволяет удалить задачу")
	fmt.Println("list - позволяет получать список задач")
	fmt.Println("done {заголовок задачи} - позволяет пометить задачу как выполненную")
	fmt.Println("log - позволяет получать список всех событий")
	fmt.Println("exit - позволяет закончить выполнение программы")
	fmt.Println()

	return ""
}

func AddTask(textSlice []string, list map[string]*structures.List) string {
	fmt.Println()
	if len(textSlice) < 3 {

		fmt.Println("ошибка формата ввода")
		return "ошибка формата ввода"
	}
	desc := ""
	for i := 2; i < len(textSlice); i++ {
		desc += textSlice[i]

		if i != len(textSlice)-1 {
			desc += " "
		}
	}

	list[textSlice[1]] = structures.MakeList(
		string(textSlice[1]),
		desc,
		time.Now(),
		false,
		time.Time{},
	)
	return ""
}

func DeleteTask(textSlice []string, list map[string]*structures.List) string {
	fmt.Println()
	if len(textSlice) < 2 {
		fmt.Println("ошибка формата ввода")
		return "ошибка формата ввода"
	}

	delete(list, textSlice[1])

	return ""
}

func GetTaskList(list map[string]*structures.List) {
	fmt.Println()
	fmt.Println("Список задач:")
	for _, l := range list {
		fmt.Println("Название:", l.GetTitle())
		fmt.Println("Описание:", l.GetDescription())
		fmt.Println("Время создания:", l.GetCreateTime())
		fmt.Println("Выполнено:", l.GetCompleted())
		if l.GetCompleteTime().IsZero() {
			fmt.Println("задача еще не сделана")
		} else {
			fmt.Println("Время выполнения:", l.GetCompleteTime())
		}

		fmt.Println()
	}
}

func CompleteTask(textSlice []string, list map[string]*structures.List) string {
	fmt.Println()
	if len(textSlice) < 2 {
		fmt.Println("ошибка формата ввода")
		return "ошибка формата ввода"
	}

	l, ok := list[textSlice[1]]
	if !ok {
		fmt.Println("задача не найдена")
		return "задача не найдена"
	}
	l.CompleteTask()

	return ""
}

func GetLogs(logs []*structures.Logs) string {
	fmt.Println()
	fmt.Println("журнал событий:")
	for _, l := range logs {
		fmt.Print(l.GetUserInput())
		fmt.Print(" ", l.GetUserInputError())
		fmt.Print(" ", l.GetUserInputTime())
		fmt.Println()
	}

	return ""
}

func Exit() {
	fmt.Println("До скорых встреч")
	os.Exit(0)
}
