package models

type PatchLessonRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	TeacherID   *uint   `json:"teacher_id"`
}
