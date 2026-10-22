package repository

import "user-management-api/internal/models"

type InMemoryUserRepository struct {
	Users []models.User
}

func NewUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		Users: make([]models.User, 0),
	}
}
