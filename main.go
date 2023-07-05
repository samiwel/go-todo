package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samiwel/go-todo/todo"
)

var todoList todo.TodoList = todo.TodoList{}
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func prompt(text string) string {
	fmt.Printf("%v: ", text)
	value, _ := reader.ReadString('\n')
	return strings.TrimSuffix(value, "\n")
}

func addTodoItem() {
	name := prompt("What do you need to do?")
	description := prompt("Description (optional)")

	todoList.Add(todo.TodoItem{Name: name, Description: description, Done: false})
}

func markAsComplete() {
	number := prompt("Select a todo item")

	index, err := strconv.Atoi(strings.TrimSuffix(number, "\n"))
	if err != nil {
		fmt.Printf("Could not find a todo item for %v", number)
		return
	}

	todoList.Items[index - 1].Complete()
}

func printMenu() {
	fmt.Printf("Menu:\n1) Add a new item\n2) Mark as complete\n3) Delete an item\n4) Exit\n\n")
}

func promptMenuOption(reader *bufio.Reader, text string, lower int, upper int) (int, error) {
	fmt.Printf("%v [%v-%v]: ", text, lower, upper)
	input, _ := reader.ReadString('\n')

	value, err := strconv.Atoi(strings.TrimSuffix(input, "\n"))
	if err != nil || value < lower || value > upper {
		return 0, fmt.Errorf("select a menu option between %v and %v", lower, upper)
	}

	return value, nil
}

func todoStatus(complete bool) string {
	if complete {
		return "✅"
	} else {
		return " "
	}
}

func deleteTodoItem() {
	number := prompt("Select a todo item")

	index, err := strconv.Atoi(number)
	if err != nil {
		fmt.Printf("Could not find a todo item for %v", number)
		return
	}

	todoList.Remove(index)
}

func main() {
	running := true

	for running {
		fmt.Printf("Your todo list:\n\n")
		for i := 0; i < len(todoList.Items); i++ {
			fmt.Printf("%v - %v %v\n", i+1, todoList.Items[i].Name, todoStatus(todoList.Items[i].Done))
		}

		fmt.Printf("\n");
		printMenu()
		option, _ := promptMenuOption(reader, "Select an option", 1, 4)

		switch option {
		case 1:
			addTodoItem()
			continue
		case 2:
			markAsComplete()
			continue
		case 3:
			deleteTodoItem()
			continue
		case 4:
			os.Exit(0)
		}

	
	}

}