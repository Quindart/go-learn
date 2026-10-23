package repository

import (
	"fmt"
	"user-management-api/internal/models"
)

type InMemoryUserRepository struct {
	Users []models.User
}

func NewUserRepository() UserRepository {
	return &InMemoryUserRepository{
		Users: make([]models.User, 0),
	}
}

func (repo *InMemoryUserRepository) FindAll() {
	fmt.Println("Find all users!!!")
}

func (repo *InMemoryUserRepository) Create() {
	fmt.Println("Create users!!!")

}

func (repo *InMemoryUserRepository) FindByUUID() {

}
func (repo *InMemoryUserRepository) Delete() {

}
func (repo *InMemoryUserRepository) Update() {

}
