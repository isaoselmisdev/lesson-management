package auth

import (
	"fmt"
	"lesson-management/entities"
	"lesson-management/pkg/common"
)

type IAuthRepository interface {
	FindAdminByEmail(email string) (*entities.Admin, error)
	FindTeacherByEmail(email string) (*entities.Teacher, error)
	FindStudentByEmail(email string) (*entities.Student, error)
	FindAdminByID(id uint) (*entities.Admin, error)
	FindTeacherByID(id uint) (*entities.Teacher, error)
	FindStudentByID(id uint) (*entities.Student, error)
	CreateAdmin(admin *entities.Admin) error
	CreateTeacher(teacher *entities.Teacher) error
	CreateStudent(student *entities.Student) error
}

type AuthRepository struct{}

func NewAuthRepository() IAuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) FindAdminByEmail(email string) (*entities.Admin, error) {
	var admin entities.Admin
	result := common.DB.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (r *AuthRepository) FindTeacherByEmail(email string) (*entities.Teacher, error) {
	var teacher entities.Teacher
	result := common.DB.Where("email = ?", email).First(&teacher)
	if result.Error != nil {
		return nil, result.Error
	}
	return &teacher, nil
}

func (r *AuthRepository) FindStudentByEmail(email string) (*entities.Student, error) {
	var student entities.Student
	result := common.DB.Where("email = ?", email).First(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func (r *AuthRepository) FindAdminByID(id uint) (*entities.Admin, error) {
	var admin entities.Admin
	result := common.DB.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (r *AuthRepository) FindTeacherByID(id uint) (*entities.Teacher, error) {
	var teacher entities.Teacher
	result := common.DB.First(&teacher, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &teacher, nil
}

func (r *AuthRepository) FindStudentByID(id uint) (*entities.Student, error) {
	var student entities.Student
	result := common.DB.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

// GetUserByRole fetches user from appropriate table based on role
func (r *AuthRepository) GetUserByRole(role string, userID uint) (interface{}, string, error) {
	switch role {
	case "admin":
		user, err := r.FindAdminByID(userID)
		if err != nil {
			return nil, "", fmt.Errorf("admin not found")
		}
		return user, user.Role, nil
	case "teacher":
		user, err := r.FindTeacherByID(userID)
		if err != nil {
			return nil, "", fmt.Errorf("teacher not found")
		}
		return user, user.Role, nil
	case "student":
		user, err := r.FindStudentByID(userID)
		if err != nil {
			return nil, "", fmt.Errorf("student not found")
		}
		return user, user.Role, nil
	default:
		return nil, "", fmt.Errorf("invalid role")
	}
}

func (r *AuthRepository) CreateAdmin(admin *entities.Admin) error {
	return common.DB.Create(admin).Error
}

func (r *AuthRepository) CreateTeacher(teacher *entities.Teacher) error {
	return common.DB.Create(teacher).Error
}

func (r *AuthRepository) CreateStudent(student *entities.Student) error {
	return common.DB.Create(student).Error
}
