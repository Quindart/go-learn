package service

import "user-management-api/internal/repository"

type UserService struct {
	UserRepo *repository.InMemoryUserRepository
}

func NewUserService(repo *repository.InMemoryUserRepository) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}
