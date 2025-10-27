package lessons

import (
	"errors"
	"lesson-management/entities"
	"lesson-management/models"
)

type ILessonService interface {
	GetLesson(id uint64) (*entities.Lesson, error)
	GetAllLessons() ([]*entities.Lesson, error)
	CreateLesson(lesson *models.CreateLessonRequest, teacherID uint) (*entities.Lesson, error)
	UpdateLesson(lesson *models.PatchLessonRequest, id uint64) (*entities.Lesson, error)
	DeleteLesson(id uint64) error
	GetTeacherLessons(teacherID uint) ([]*entities.Lesson, error)
	GetStudentLessons(studentID uint) ([]*entities.Lesson, error)
	AssignTeacherToLesson(lessonID uint64, teacherID uint) error
	EnrollStudentInLesson(lessonID uint64, studentID uint) error
	RemoveStudentFromLesson(lessonID uint64, studentID uint) error
	GetLessonStudents(lessonID uint64, teacherID uint) ([]entities.Student, error)
}

type LessonService struct {
	repo ILessonRepository
}

func NewLessonService(repo ILessonRepository) ILessonService {
	return &LessonService{
		repo: repo,
	}
}

func (s *LessonService) GetLesson(id uint64) (*entities.Lesson, error) {
	lesson, err := s.repo.GetLesson(uint(id))
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (s *LessonService) GetAllLessons() ([]*entities.Lesson, error) {
	lessons, err := s.repo.GetAllLessons()
	if err != nil {
		return nil, err
	}

	return lessons, err
}

func (s *LessonService) CreateLesson(lessonRequest *models.CreateLessonRequest, teacherID uint) (*entities.Lesson, error) {
	lesson := &entities.Lesson{
		Title:       lessonRequest.Title,
		Description: lessonRequest.Description,
		TeacherID:   teacherID,
	}

	err := s.repo.CreateLesson(lesson)
	if err != nil {
		return nil, err
	}

	return lesson, nil
}

func (s *LessonService) UpdateLesson(lessonRequest *models.PatchLessonRequest, id uint64) (*entities.Lesson, error) {
	var lesson entities.Lesson

	lesson.ID = uint(id)

	if lessonRequest.Title != nil {
		lesson.Title = *lessonRequest.Title
	}
	if lessonRequest.Description != nil {
		lesson.Description = *lessonRequest.Description
	}
	if lessonRequest.TeacherID != nil {
		lesson.TeacherID = *lessonRequest.TeacherID
	}

	err := s.repo.UpdateLesson(&lesson)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (s *LessonService) DeleteLesson(id uint64) error {
	return s.repo.DeleteLesson(uint(id))
}

func (s *LessonService) GetTeacherLessons(teacherID uint) ([]*entities.Lesson, error) {
	return s.repo.GetLessonsByTeacherID(teacherID)
}

func (s *LessonService) GetStudentLessons(studentID uint) ([]*entities.Lesson, error) {
	return s.repo.GetLessonsByStudentID(studentID)
}

func (s *LessonService) AssignTeacherToLesson(lessonID uint64, teacherID uint) error {
	return s.repo.AssignTeacherToLesson(uint(lessonID), teacherID)
}

func (s *LessonService) EnrollStudentInLesson(lessonID uint64, studentID uint) error {
	return s.repo.EnrollStudentInLesson(uint(lessonID), studentID)
}

func (s *LessonService) RemoveStudentFromLesson(lessonID uint64, studentID uint) error {
	return s.repo.RemoveStudentFromLesson(uint(lessonID), studentID)
}

func (s *LessonService) GetLessonStudents(lessonID uint64, teacherID uint) ([]entities.Student, error) {
	// Verify lesson belongs to teacher
	lesson, err := s.repo.GetLesson(uint(lessonID))
	if err != nil {
		return nil, err
	}

	if lesson.TeacherID != teacherID {
		return nil, errors.New("lesson does not belong to this teacher")
	}

	return s.repo.GetLessonStudents(uint(lessonID))
}
