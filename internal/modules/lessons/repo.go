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
	GetLessonsByTeacherID(teacherID uint) ([]*entities.Lesson, error)
	GetLessonsByStudentID(studentID uint) ([]*entities.Lesson, error)
	AssignTeacherToLesson(lessonID uint, teacherID uint) error
	EnrollStudentInLesson(lessonID uint, studentID uint) error
	RemoveStudentFromLesson(lessonID uint, studentID uint) error
	GetLessonStudents(lessonID uint) ([]entities.Student, error)
}

type LessonRepository struct{}

func NewLessonRepository() ILessonRepository {
	return &LessonRepository{}
}

func (r *LessonRepository) GetLesson(id uint) (entities.Lesson, error) {
	var lesson entities.Lesson
	result := common.DB.Preload("Teacher").Preload("Students").First(&lesson, id)
	return lesson, result.Error
}

func (r *LessonRepository) GetAllLessons() ([]*entities.Lesson, error) {
	var lessons []*entities.Lesson
	result := common.DB.Preload("Teacher").Preload("Students").Find(&lessons)

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

func (r *LessonRepository) GetLessonsByTeacherID(teacherID uint) ([]*entities.Lesson, error) {
	var lessons []*entities.Lesson
	result := common.DB.Where("teacher_id = ?", teacherID).Preload("Students").Find(&lessons)
	return lessons, result.Error
}

func (r *LessonRepository) GetLessonsByStudentID(studentID uint) ([]*entities.Lesson, error) {
	var lessons []*entities.Lesson
	result := common.DB.Model(&entities.Lesson{}).
		Joins("JOIN lesson_students ON lessons.id = lesson_students.lesson_id").
		Where("lesson_students.student_id = ?", studentID).
		Preload("Teacher").
		Find(&lessons)
	return lessons, result.Error
}

func (r *LessonRepository) AssignTeacherToLesson(lessonID uint, teacherID uint) error {
	result := common.DB.Model(&entities.Lesson{}).Where("id = ?", lessonID).Update("teacher_id", teacherID)
	return result.Error
}

func (r *LessonRepository) EnrollStudentInLesson(lessonID uint, studentID uint) error {
	lesson := &entities.Lesson{}
	student := &entities.Student{}

	if err := common.DB.First(lesson, lessonID).Error; err != nil {
		return err
	}

	if err := common.DB.First(student, studentID).Error; err != nil {
		return err
	}

	// Check if already enrolled
	var count int64
	common.DB.Model(&entities.Lesson{ID: lessonID}).Association("Students").Count()
	common.DB.Table("lesson_students").
		Where("lesson_id = ? AND student_id = ?", lessonID, studentID).
		Count(&count)

	if count > 0 {
		return fmt.Errorf("student already enrolled in this lesson")
	}

	return common.DB.Model(lesson).Association("Students").Append(student)
}

func (r *LessonRepository) RemoveStudentFromLesson(lessonID uint, studentID uint) error {
	lesson := &entities.Lesson{}
	student := &entities.Student{}

	if err := common.DB.First(lesson, lessonID).Error; err != nil {
		return err
	}

	if err := common.DB.First(student, studentID).Error; err != nil {
		return err
	}

	return common.DB.Model(lesson).Association("Students").Delete(student)
}

func (r *LessonRepository) GetLessonStudents(lessonID uint) ([]entities.Student, error) {
	var lesson entities.Lesson
	if err := common.DB.Preload("Students").First(&lesson, lessonID).Error; err != nil {
		return nil, err
	}
	return lesson.Students, nil
}
