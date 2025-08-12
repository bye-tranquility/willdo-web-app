package repository

import "willdo/internal/models"

type EventRepository interface {
	GetAll() (models.Events, error)
	GetByID(id int) (*models.Event, error)
	Create(event *models.Event) error
	Update(event *models.Event) error
	Delete(id int) error
}
