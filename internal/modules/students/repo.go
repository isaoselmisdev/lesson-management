package students

import (
	"fmt"
	"lesson-management/entities"
	"lesson-management/pkg/common"
)

type IStudentRepository interface {
	GetStudentByID(id uint) (entities.Student, error)
	GetAllStudents() ([]entities.Student, error)
	CreateStudent(student *entities.Student) error
	DeleteStudent(id uint) error
	UpdateStudent(student *entities.Student) error
}
type StudentRepository struct{}

func NewStudentRepository() IStudentRepository {
	return &StudentRepository{}
}

func (r *StudentRepository) GetStudentByID(id uint) (entities.Student, error) {
	var student entities.Student
	result := common.DB.First(&student, id)
	return student, result.Error
}

func (r *StudentRepository) GetAllStudents() ([]entities.Student, error) {
	var students []entities.Student
	result := common.DB.Find(&students)
	return students, result.Error
}

func (r *StudentRepository) CreateStudent(student *entities.Student) error {
	return common.DB.Create(student).Error
}

func (r *StudentRepository) DeleteStudent(id uint) error {
	return common.DB.Delete(&entities.Student{}, id).Error
}
func (r *StudentRepository) UpdateStudent(student *entities.Student) error {
	result := common.DB.Save(student)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
