package lessons

import (
	"lesson-management/entities"
	"lesson-management/models"
)

type ILessonService interface {
	GetLesson(id uint64) (*entities.Lesson, error)
	GetAllLessons() ([]*entities.Lesson, error)
	CreateLesson(lesson *models.CreateLessonRequest) (*entities.Lesson, error)
	UpdateLesson(lesson *models.PatchLessonRequest, id uint64) (*entities.Lesson, error)
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

func (s *LessonService) CreateLesson(lessonRequest *models.CreateLessonRequest) (*entities.Lesson, error) {
	lesson := &entities.Lesson{
		Title:       lessonRequest.Title,
		Description: lessonRequest.Description,
		Teacher:     lessonRequest.Teacher,
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
	if lessonRequest.Teacher != nil {
		lesson.Teacher = *lessonRequest.Teacher
	}

	err := s.repo.UpdateLesson(&lesson)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

// func (s *LessonService) DeleteLesson(id uint64) error {}
