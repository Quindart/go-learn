package service

import "user-management-api/internal/repository"

type demoService struct {
	repo repository.DemoRepository
}

func NewDemoService(repo repository.DemoRepository) DemoService {
	return &demoService{
		repo: repo,
	}
}

func (service *demoService) GetAllDemos() {
	service.repo.FindAll()

}
func (service *demoService) CreateDemo() {
	service.repo.Create()

}
func (service *demoService) UpdateDemo() {

}
func (service *demoService) DeleteDemo() {}
