package models

// Event represents a to-do list event/task
type Event struct {
	ID          int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Description string `json:"description" validate:"required" gorm:"not null"`
	Completed   bool   `json:"completed" gorm:"default:false"`
	DueDate     string `json:"due" validate:"omitempty,date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Events []*Event
