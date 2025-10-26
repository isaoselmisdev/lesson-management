package students

import (
	"encoding/json"
	"fmt"
	"lesson-management/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	service IStudentService
}

func NewStudentHandler(service IStudentService) *StudentHandler {
	return &StudentHandler{
		service: service,
	}
}

func (h *StudentHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Implement GET logic for lessons
	studentIDstr := mux.Vars(r)["studentID"]
	studentID, err := strconv.ParseUint(studentIDstr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	// Use studentID to fetch student details
	student, err := h.service.GetStudentByID(studentID)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		fmt.Println("Error while fetching student: ", err)
		return
	}

	// Return student details as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, "Failed to encode student data", http.StatusInternalServerError)
		return
	}
}

func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get students by lesson
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Implement POST logic for lessons
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	studentIDstr := mux.Vars(r)["studentID"]
	studentID, err := strconv.ParseUint(studentIDstr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	_, err = h.service.GetStudentByID(studentID)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		fmt.Println("Error while fetching student: ", err)
		return
	}

	var requestBody models.PatchStudentRequest
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	student, err := h.service.UpdateStudent(&requestBody, studentID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error while updating student: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, "Failed to encode student data", http.StatusInternalServerError)
		return
	}
}
