package repository

type UserRepository interface {
	FindAll()
	Create()
	FindByUUID()
	Update()
	Delete()
}

type DemoRepository interface {
	FindAll()
	Create()
	FindByUUID()
	Update()
	Delete()
}
type Demo2Repository interface {
	FindAll()
	Create()
	FindByUUID()
	Update()
	Delete()
}
