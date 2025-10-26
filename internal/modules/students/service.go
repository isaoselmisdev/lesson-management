package students

import (
	"lesson-management/entities"
	"lesson-management/models"
)

type IStudentService interface {
	GetStudentByID(id uint64) (*entities.Student, error)
	GetAllStudents() ([]entities.Student, error)
	CreateStudent(student *models.CreateStudentRequest) (*entities.Student, error)
	UpdateStudent(student *models.PatchStudentRequest, id uint64) (*entities.Student, error)
}

type StudentService struct {
	repo IStudentRepository
}

func NewStudentService(repo IStudentRepository) IStudentService {
	return &StudentService{
		repo: repo,
	}
}

func (s *StudentService) GetStudentByID(id uint64) (*entities.Student, error) {
	student, err := s.repo.GetStudentByID(uint(id))
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *StudentService) GetAllStudents() ([]entities.Student, error) {
	students, err := s.repo.GetAllStudents()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *StudentService) CreateStudent(request *models.CreateStudentRequest) (*entities.Student, error) {
	student := &entities.Student{
		Name:  request.Name,
		Email: request.Email,
	}

	err := s.repo.CreateStudent(student)
	if err != nil {
		return nil, err
	}

	return student, nil

}

func (s *StudentService) UpdateStudent(request *models.PatchStudentRequest, id uint64) (*entities.Student, error) {
	var student entities.Student

	student.ID = uint(id)

	if request.Name != nil {
		student.Name = *request.Name
	}
	if request.Email != nil {
		student.Email = *request.Email
	}

	err := s.repo.UpdateStudent(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}
