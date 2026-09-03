package author

type Author struct {
	name   string
	author string
}

func New(name string) Author {
	return Author{name: name}
}

func (a Author) GetName() string {
	return a.name
}

func (a Author) GetAuthor() string {
	return a.author
}
func (a Author) Speak() string {
	return "Hello, my name is " + a.name + " and I am the author of this book."
}
