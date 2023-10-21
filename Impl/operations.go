package Impl

var TodoItems = []Todo{}

func GetTodos() []Todo {
	return TodoItems
}

func GetTodoByID(id int) Todo {
	for _, todo := range TodoItems {
		if todo.Id == id {
			return todo
		}
	}
	return Todo{}
}

func AddTodoItem(title string) *Todo {
	temp := &Todo{
		Id:        len(TodoItems) + 1,
		Title:     title,
		Completed: false,
	}
	TodoItems = append(TodoItems, *temp)
	return temp
}

func UpdateTodoItem(id int, title string, completed bool) bool {
	for i, todo := range TodoItems {
		if todo.Id == id {
			TodoItems[i].Title = title
			TodoItems[i].Completed = completed
			return true
		}
	}
	return false
}

func DeleteTodoItem(id int) bool {
	for i, todo := range TodoItems {
		if todo.Id == id {
			TodoItems = append(TodoItems[:i], TodoItems[i+1:]...)
			return true
		}
	}
	return false
}
