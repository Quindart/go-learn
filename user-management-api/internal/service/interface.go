package service

type UserService interface {
	GetAllUsers()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type DemoService interface {
	GetAllDemos()
	CreateDemo()
	UpdateDemo()
	DeleteDemo()
}

type Demo2Service interface {
	GetAllDemos()
	CreateDemo()
	UpdateDemo()
	DeleteDemo()
}
