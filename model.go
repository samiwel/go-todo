package main

type TodoItem struct {
	Name string;
	Description string;
	Done bool;
}

func (todoItem *TodoItem) Complete() {
	todoItem.Done = true
}

type TodoList struct {
	Items []TodoItem;
}

func (tl *TodoList) add(item TodoItem) {
	tl.Items = append(tl.Items, item)
}