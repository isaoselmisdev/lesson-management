package main

import (
	"lesson-management/internal/modules/auth"
	"lesson-management/internal/modules/lessons"
	"lesson-management/internal/modules/students"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Initialize Auth
	authRepo := auth.NewAuthRepository()
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)
	auth.InitRoutes(router, authHandler)

	// Initialize Lessons
	lessonRepo := lessons.NewLessonRepository()
	lessonService := lessons.NewLessonService(lessonRepo)
	lessonHandler := lessons.NewLessonHandler(lessonService)
	lessons.InitRoutes(router, lessonHandler, authService)

	// Initialize Students
	studentRepo := students.NewStudentRepository()
	studentService := students.NewStudentService(studentRepo)
	studentHandler := students.NewStudentHandler(studentService)
	students.InitRoutes(router, studentHandler, authService)

	return router
}
