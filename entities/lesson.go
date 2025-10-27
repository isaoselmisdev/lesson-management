package entities

import "time"

type Lesson struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	TeacherID   uint      `json:"teacher_id"`
	Teacher     Teacher   `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	Students    []Student `gorm:"many2many:lesson_students;" json:"students,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
