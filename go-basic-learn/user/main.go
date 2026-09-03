package user

type User struct {
	name string
	age  int
}

func New(name string, age int) *User {
	return &User{
		name: name,
		age:  age,
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int {
	return u.age
}
