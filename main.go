package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)



func printMenu() {
	fmt.Printf("Menu:\n1) Add a new item\n2) Mark as complete\n3) Delete an item\n4) Exit\n\n")
}

func promptMenuOption(reader *bufio.Reader, text string) (int, error) {
	fmt.Print(text)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	value, err := strconv.Atoi(input)
	fmt.Printf("Debug: %v %v", value, err)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func main() {

	todoList := TodoList{}
	running := true
	reader := bufio.NewReader(os.Stdin)

	for running {
		fmt.Printf("Your todo list:\n\n")
		for i := 0; i < len(todoList.Items); i++ {
			fmt.Printf("%v - %v", i+1, todoList.Items[i].Name)
		}

		fmt.Printf("\n");
		printMenu()
		option, err := promptMenuOption(reader, "Enter option number (1-4): ")
		if err != nil {
			fmt.Errorf("%v", err)
			continue
		}

		fmt.Printf("You selected %v", option)

		

	}


	// todoList.addTodoItem(TodoItem{Name: "Buy some milk", Description: "Go to the grocery store to buy some milk", Done: false})

}