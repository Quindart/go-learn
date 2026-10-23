package service

import "user-management-api/internal/repository"

type userService struct {
	UserRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		UserRepo: repo,
	}
}

func (service *userService) GetAllUsers() {
	service.UserRepo.FindAll()
}

func (service *userService) CreateUser() {

}

func (service *userService) UpdateUser() {

}
func (service *userService) DeleteUser() {

}
