package models

type CreateLessonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Teacher     string `json:"teacher"`
}
