package repository

import (
	"fmt"
	"user-management-api/internal/models"
)

type demoRepository struct {
	Demos []models.Demo
}

func NewDemoRepository() DemoRepository {
	return &demoRepository{
		Demos: make([]models.Demo, 0),
	}
}

func (repo *demoRepository) FindAll() {
	fmt.Println("Find all demo code")
}
func (repo *demoRepository) Create() {

}
func (repo *demoRepository) FindByUUID() {

}

func (repo *demoRepository) Update() {

}
func (repo *demoRepository) Delete() {

}
