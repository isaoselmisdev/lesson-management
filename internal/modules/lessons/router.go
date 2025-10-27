package lessons

import (
	"lesson-management/internal/modules/auth"
	"lesson-management/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, handler *LessonHandler, authService auth.IAuthService) {
	// Authentication middleware
	authMiddleware := middleware.AuthMiddleware(authService)

	// Public endpoints (no auth required)
	router.HandleFunc("/api/lessons/{lessonID:[0-9]+}", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/lessons", handler.List).Methods(http.MethodGet)

	// Admin-only endpoints
	adminRoutes := router.PathPrefix("/api/lessons").Subrouter()
	adminRoutes.Use(authMiddleware)
	adminRoutes.Use(middleware.RequireRole("admin"))
	adminRoutes.HandleFunc("", handler.Create).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/{lessonID:[0-9]+}", handler.Update).Methods(http.MethodPut)
	adminRoutes.HandleFunc("/{lessonID:[0-9]+}", handler.Delete).Methods(http.MethodDelete)
	adminRoutes.HandleFunc("/{lessonID:[0-9]+}/assign-teacher", handler.AssignTeacher).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/{lessonID:[0-9]+}/enroll-student", handler.EnrollStudent).Methods(http.MethodPost)

	// Teacher-only endpoints
	teacherRoutes := router.PathPrefix("/api/teacher").Subrouter()
	teacherRoutes.Use(authMiddleware)
	teacherRoutes.Use(middleware.RequireRole("teacher"))
	teacherRoutes.HandleFunc("/lessons", handler.GetTeacherLessons).Methods(http.MethodGet)

	lessonStudentRoutes := router.PathPrefix("/api/lessons/{lessonID:[0-9]+}/students").Subrouter()
	lessonStudentRoutes.Use(authMiddleware)
	lessonStudentRoutes.Use(middleware.RequireRole("teacher"))
	lessonStudentRoutes.HandleFunc("", handler.GetLessonStudents).Methods(http.MethodGet)
	lessonStudentRoutes.HandleFunc("", handler.AddStudentToLesson).Methods(http.MethodPost)
	lessonStudentRoutes.HandleFunc("/{studentID:[0-9]+}", handler.RemoveStudentFromLesson).Methods(http.MethodDelete)

	// Student-only endpoints
	studentRoutes := router.PathPrefix("/api/student").Subrouter()
	studentRoutes.Use(authMiddleware)
	studentRoutes.Use(middleware.RequireRole("student"))
	studentRoutes.HandleFunc("/lessons", handler.GetStudentLessons).Methods(http.MethodGet)
}
