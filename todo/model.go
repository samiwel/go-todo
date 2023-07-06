package todo

type TodoItem struct {
	Name        string
	Description string
	Done        bool
}

func (todoItem *TodoItem) Complete() {
	todoItem.Done = true
}

type TodoList struct {
	Items []TodoItem
}

func (tl *TodoList) Add(item TodoItem) {
	tl.Items = append(tl.Items, item)
}

func (tl *TodoList) Remove(item int) {
	index := item - 1
	ret := make([]TodoItem, 0)
	ret = append(ret, tl.Items[:index]...)
	if index > 0 {
		ret = append(ret, tl.Items[index-1:]...)
	}
	tl.Items = ret
}
