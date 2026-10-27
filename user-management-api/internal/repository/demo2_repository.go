package repository

import "user-management-api/internal/models"

type demo2Repository struct {
	demo2 []models.Demo2
}

func NewDemo2Repository() Demo2Repository {
	return &demo2Repository{
		demo2: make([]models.Demo2, 0),
	}
}

// Create implements [Demo2Repository].
func (d *demo2Repository) Create() {
	panic("unimplemented")
}

// Delete implements [Demo2Repository].
func (d *demo2Repository) Delete() {
	panic("unimplemented")
}

// FindAll implements [Demo2Repository].
func (d *demo2Repository) FindAll() {
	panic("unimplemented")
}

// FindByUUID implements [Demo2Repository].
func (d *demo2Repository) FindByUUID() {
	panic("unimplemented")
}

// Update implements [Demo2Repository].
func (d *demo2Repository) Update() {
	panic("unimplemented")
}
