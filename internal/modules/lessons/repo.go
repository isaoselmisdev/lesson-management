package lessons

import (
	"fmt"
	"lesson-management/entities"
	"lesson-management/pkg/common"
)

type ILessonRepository interface {
	GetLesson(id uint) (entities.Lesson, error)
	GetAllLessons() ([]*entities.Lesson, error)
	CreateLesson(lesson *entities.Lesson) error
	UpdateLesson(lesson *entities.Lesson) error
	DeleteLesson(id uint) error
}

type LessonRepository struct{}

func NewLessonRepository() ILessonRepository {
	return &LessonRepository{}
}

func (r *LessonRepository) GetLesson(id uint) (entities.Lesson, error) {
	var lesson entities.Lesson
	result := common.DB.Preload("Students").First(&lesson, id)
	return lesson, result.Error
}

func (r *LessonRepository) GetAllLessons() ([]*entities.Lesson, error) {
	var lessons []*entities.Lesson
	result := common.DB.Preload("Students").Find(&lessons)

	return lessons, result.Error
}

func (r *LessonRepository) CreateLesson(lesson *entities.Lesson) error {
	return common.DB.Create(lesson).Error
}

func (r *LessonRepository) UpdateLesson(lesson *entities.Lesson) error {
	result := common.DB.Save(lesson)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (r *LessonRepository) DeleteLesson(id uint) error {
	return common.DB.Delete(&entities.Lesson{}, id).Error
}
