package lessons

import (
	"encoding/json"
	"fmt"
	"lesson-management/models"
	"lesson-management/pkg/middleware"
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

	lesson, err := h.service.CreateLesson(&requestBody, requestBody.TeacherID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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

func (h *LessonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteLesson(lessonID)
	if err != nil {
		http.Error(w, "Failed to delete lesson", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Teacher handlers
func (h *LessonHandler) GetTeacherLessons(w http.ResponseWriter, r *http.Request) {
	// Get teacherID from context set by middleware
	teacherID, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	lessons, err := h.service.GetTeacherLessons(teacherID)
	if err != nil {
		http.Error(w, "Failed to fetch lessons", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lessons)
}

// Student handlers
func (h *LessonHandler) GetStudentLessons(w http.ResponseWriter, r *http.Request) {
	studentID, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	lessons, err := h.service.GetStudentLessons(studentID)
	if err != nil {
		http.Error(w, "Failed to fetch lessons", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lessons)
}

// Admin handlers for lesson management
func (h *LessonHandler) AssignTeacher(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	var req struct {
		TeacherID uint `json:"teacher_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.AssignTeacherToLesson(lessonID, req.TeacherID)
	if err != nil {
		http.Error(w, "Failed to assign teacher", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LessonHandler) EnrollStudent(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	var req struct {
		StudentID uint `json:"student_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.EnrollStudentInLesson(lessonID, req.StudentID)
	if err != nil {
		http.Error(w, "Failed to enroll student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Teacher handlers for student management
func (h *LessonHandler) GetLessonStudents(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	teacherID, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	students, err := h.service.GetLessonStudents(lessonID, teacherID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

func (h *LessonHandler) AddStudentToLesson(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	var req struct {
		StudentID uint `json:"student_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.EnrollStudentInLesson(lessonID, req.StudentID)
	if err != nil {
		http.Error(w, "Failed to add student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LessonHandler) RemoveStudentFromLesson(w http.ResponseWriter, r *http.Request) {
	lessonIDStr := mux.Vars(r)["lessonID"]
	lessonID, err := strconv.ParseUint(lessonIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lesson ID", http.StatusBadRequest)
		return
	}

	studentIDStr := mux.Vars(r)["studentID"]
	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	err = h.service.RemoveStudentFromLesson(lessonID, uint(studentID))
	if err != nil {
		http.Error(w, "Failed to remove student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
