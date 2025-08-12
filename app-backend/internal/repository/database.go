package repository

import (
	"time"
	"willdo/internal/config"
	"willdo/internal/errors"
	"willdo/internal/models"

	"gorm.io/gorm"
)

type DatabaseEventRepository struct {
	db *gorm.DB
}

func NewDatabaseEventRepository(db *gorm.DB) *DatabaseEventRepository {
	return &DatabaseEventRepository{db: db}
}

func (r *DatabaseEventRepository) GetAll() (models.Events, error) {
	var events models.Events
	result := r.db.Find(&events)
	return events, result.Error
}

func (r *DatabaseEventRepository) GetByID(id int) (*models.Event, error) {
	var event models.Event
	result := r.db.First(&event, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrEventNotFound
		}
		return nil, result.Error
	}
	return &event, nil
}

func (r *DatabaseEventRepository) Create(event *models.Event) error {
	now := time.Now().Format(config.DateTime)
	event.CreatedAt = now
	event.UpdatedAt = now
	return r.db.Create(event).Error
}

func (r *DatabaseEventRepository) Update(event *models.Event) error {
	event.UpdatedAt = time.Now().Format(config.DateTime)
	result := r.db.Save(event)
	if result.RowsAffected == 0 {
		return errors.ErrEventNotFound
	}
	return result.Error
}

func (r *DatabaseEventRepository) Delete(id int) error {
	result := r.db.Delete(&models.Event{}, id)
	if result.RowsAffected == 0 {
		return errors.ErrEventNotFound
	}
	return result.Error
}
