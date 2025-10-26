package entities

import "time"

// entity
type Lesson struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Teacher     string    `json:"teacher"`
	Students    []Student `gorm:"many2many:lesson_students;" json:"students"`
	CreatedAt   time.Time `json:"created_at"`
}
