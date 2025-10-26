package models

type CreateStudentRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
