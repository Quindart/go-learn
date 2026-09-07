package todo

type todo struct {
	title string
	name string
}

func NewTodo(title string , name string) todo{

	return todo{
		title: title,
		name: name,
	}
}

