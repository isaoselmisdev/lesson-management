package lessons

import (
	"encoding/json"
	"fmt"
	"lesson-management/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LessonHandler struct {
	service ILessonService
}

func NewLessonHandler(service ILessonService) *LessonHandler {
	return &LessonHandler{
		service: service,
	}
}

func (h *LessonHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Implement GET logic for lessons
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	// Use lessonID to fetch lesson detailss
	lesson, err := h.service.GetLesson(lessonID)
	if err != nil {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		fmt.Println("Error while fetching lesson: ", err)
		return
	}

	// Return lesson details as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(lesson)
	if err != nil {
		http.Error(w, "Failed to encode lesson data", http.StatusInternalServerError)
		return
	}
}

func (h *LessonHandler) List(w http.ResponseWriter, r *http.Request) {
	// Use lessonID to fetch lesson detailss
	lessons, err := h.service.GetAllLessons()
	if err != nil {
		http.Error(w, "Lessons not found", http.StatusNotFound)
		fmt.Println("Error while fetching lessons: ", err)
		return
	}

	// Return lesson details as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(lessons)
	if err != nil {
		http.Error(w, "Failed to encode lesson data", http.StatusInternalServerError)
		return
	}
}

func (h *LessonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody models.CreateLessonRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	lesson, err := h.service.CreateLesson(&requestBody)
	if err != nil {
		http.Error(w, "ınternal server error ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(lesson)
	if err != nil {
		http.Error(w, "Failed to encode lesson data", http.StatusInternalServerError)
		return
	}
}

func (h *LessonHandler) Update(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	_, err = h.service.GetLesson(lessonID)
	if err != nil {
		http.Error(w, "Lesson not found", http.StatusNotFound)
		fmt.Println("Error while fetching lesson: ", err)
		return
	}

	var requestBody models.PatchLessonRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validation
	if requestBody.Title != nil && *requestBody.Title == "" {
		http.Error(w, "Lesson name cannot be empty", http.StatusBadRequest)
		return
	}

	if requestBody.Description != nil && *requestBody.Description == "" {
		http.Error(w, "Lesson description cannot be empty", http.StatusBadRequest)
		return
	}

	if requestBody.Teacher != nil && *requestBody.Teacher == "" {
		http.Error(w, "Lesson teacher cannot be empty", http.StatusBadRequest)
		return
	}

	lesson, err := h.service.UpdateLesson(&requestBody, lessonID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error while updating lesson: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lesson)
}
