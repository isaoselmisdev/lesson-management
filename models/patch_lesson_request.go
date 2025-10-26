package models

type PatchLessonRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Teacher     *string `json:"teacher"`
}
