package service

import "user-management-api/internal/repository"

type demo2Service struct {
	repo repository.Demo2Repository
}

func NewDemo2Service(repo repository.Demo2Repository) Demo2Service {
	return &demo2Service{
		repo: repo,
	}
}

// CreateDemo implements [Demo2Service].
func (d *demo2Service) CreateDemo() {
	panic("unimplemented")
}

// DeleteDemo implements [Demo2Service].
func (d *demo2Service) DeleteDemo() {
	panic("unimplemented")
}

// GetAllDemos implements [Demo2Service].
func (d *demo2Service) GetAllDemos() {
	panic("unimplemented")
}

// UpdateDemo implements [Demo2Service].
func (d *demo2Service) UpdateDemo() {
	panic("unimplemented")
}
