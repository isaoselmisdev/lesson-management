package models

type CreateLessonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TeacherID   uint   `json:"teacher_id"`
}
