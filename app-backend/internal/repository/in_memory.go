/* For test purposes mostly */

package repository

import (
	"time"
	"willdo/internal/config"
	"willdo/internal/errors"
	"willdo/internal/models"
)

type InMemoryEventRepository struct {
	events models.Events
}

func NewInMemoryEventRepository() *InMemoryEventRepository {
	events := models.Events{
		&models.Event{
			ID:          1,
			Description: "party",
			Completed:   true,
			DueDate:     (time.Now().Add(30 * time.Minute)).Format(config.DateTime),
			CreatedAt:   time.Now().Format(config.DateTime),
			UpdatedAt:   time.Now().Format(config.DateTime),
		},
		&models.Event{
			ID:          2,
			Description: "cinema",
			Completed:   false,
			DueDate:     time.Now().Format(config.DateTime),
			CreatedAt:   time.Now().Format(config.DateTime),
			UpdatedAt:   time.Now().Format(config.DateTime),
		},
	}

	return &InMemoryEventRepository{
		events: events,
	}
}

func (r *InMemoryEventRepository) GetAll() (models.Events, error) {
	return r.events, nil
}

func (r *InMemoryEventRepository) GetByID(id int) (*models.Event, error) {
	i := r.findIndexByEventID(id)
	if i == -1 {
		return nil, errors.ErrEventNotFound
	}
	return r.events[i], nil
}

func (r *InMemoryEventRepository) Create(event *models.Event) error {
	event.ID = r.getNextID()
	now := time.Now().Format(config.DateTime)
	event.CreatedAt = now
	event.UpdatedAt = now
	r.events = append(r.events, event)
	return nil
}

func (r *InMemoryEventRepository) Update(event *models.Event) error {
	i := r.findIndexByEventID(event.ID)
	if i == -1 {
		return errors.ErrEventNotFound
	}
	event.UpdatedAt = time.Now().Format(config.DateTime)
	r.events[i] = event
	return nil
}

func (r *InMemoryEventRepository) Delete(id int) error {
	i := r.findIndexByEventID(id)
	if i == -1 {
		return errors.ErrEventNotFound
	}
	r.events = append(r.events[:i], r.events[i+1:]...)
	return nil
}

func (r *InMemoryEventRepository) getNextID() int {
	if len(r.events) == 0 {
		return 1
	}
	lastEvent := r.events[len(r.events)-1]
	return lastEvent.ID + 1
}

func (r *InMemoryEventRepository) findIndexByEventID(id int) int {
	for i, event := range r.events {
		if event.ID == id {
			return i
		}
	}
	return -1
}
