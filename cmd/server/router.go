package main

import (
	"lesson-management/internal/modules/lessons"
	"lesson-management/internal/modules/students"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	api := mux.NewRouter().StrictSlash(false).PathPrefix("/").Subrouter()

	lessonRepo := lessons.NewLessonRepository()
	lessonService := lessons.NewLessonService(lessonRepo)
	lessonHandler := lessons.NewLessonHandler(lessonService)

	studentRepo := students.NewStudentRepository()
	studentService := students.NewStudentService(studentRepo)
	studentHandler := students.NewStudentHandler(studentService)

	// Call init routes for the modules
	lessons.InitRoutes(api, lessonHandler)
	students.InitRoutes(api, studentHandler)

	return api
}
