package main


type User struct {
	Name  string `json:"name" db:"name"`
	Age   int    `json:"age" db:"age"`
	Clazz string `json:"clazz" db:"clazz"` 
}

func main() {

}
