package auth

import (
	"errors"
	"fmt"
	"lesson-management/entities"
	"lesson-management/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(email, password, role string) (*models.LoginResponse, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
	GenerateToken(userID uint, role string, name string) (string, error)
	RegisterAdmin(name, email, password string) (*models.CreateUserResponse, error)
	RegisterTeacher(name, email, password string) (*models.CreateUserResponse, error)
	RegisterStudent(name, email, password string) (*models.CreateUserResponse, error)
}

type AuthService struct {
	repo   IAuthRepository
	secret string
}

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func NewAuthService(repo IAuthRepository) IAuthService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-secret-key-change-in-production" // Fallback for development
	}

	return &AuthService{
		repo:   repo,
		secret: secret,
	}
}

func (s *AuthService) Login(email, password, role string) (*models.LoginResponse, error) {
	var userID uint
	var userName string

	switch role {
	case "admin":
		admin, err := s.repo.FindAdminByEmail(email)
		if err != nil {
			return nil, errors.New("invalid credentials")
		}
		if !s.checkPasswordHash(password, admin.Password) {
			return nil, errors.New("invalid credentials")
		}
		userID = admin.ID
		userName = admin.Name

	case "teacher":
		teacher, err := s.repo.FindTeacherByEmail(email)
		if err != nil {
			return nil, errors.New("invalid credentials")
		}
		if !s.checkPasswordHash(password, teacher.Password) {
			return nil, errors.New("invalid credentials")
		}
		userID = teacher.ID
		userName = teacher.Name

	case "student":
		student, err := s.repo.FindStudentByEmail(email)
		if err != nil {
			return nil, errors.New("invalid credentials")
		}
		if !s.checkPasswordHash(password, student.Password) {
			return nil, errors.New("invalid credentials")
		}
		userID = student.ID
		userName = student.Name

	default:
		return nil, errors.New("invalid role")
	}

	token, err := s.GenerateToken(userID, role, userName)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &models.LoginResponse{
		Token: token,
		Role:  role,
		Name:  userName,
	}, nil
}

func (s *AuthService) GenerateToken(userID uint, role string, name string) (string, error) {
	claims := &JWTClaims{
		UserID: userID,
		Role:   role,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *AuthService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		// Verify token is not expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			return nil, errors.New("token expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *AuthService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) RegisterAdmin(name, email, password string) (*models.CreateUserResponse, error) {
	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return nil, err
	}

	admin := &entities.Admin{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "admin",
	}

	if err := s.repo.CreateAdmin(admin); err != nil {
		return nil, err
	}

	return &models.CreateUserResponse{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Role:  admin.Role,
	}, nil
}

func (s *AuthService) RegisterTeacher(name, email, password string) (*models.CreateUserResponse, error) {
	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return nil, err
	}

	teacher := &entities.Teacher{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "teacher",
	}

	if err := s.repo.CreateTeacher(teacher); err != nil {
		return nil, err
	}

	return &models.CreateUserResponse{
		ID:    teacher.ID,
		Name:  teacher.Name,
		Email: teacher.Email,
		Role:  teacher.Role,
	}, nil
}

func (s *AuthService) RegisterStudent(name, email, password string) (*models.CreateUserResponse, error) {
	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return nil, err
	}

	student := &entities.Student{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "student",
	}

	if err := s.repo.CreateStudent(student); err != nil {
		return nil, err
	}

	return &models.CreateUserResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
		Role:  student.Role,
	}, nil
}
