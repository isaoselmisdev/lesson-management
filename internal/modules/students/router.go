package students

import (
	"lesson-management/internal/modules/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, handler *StudentHandler, authService auth.IAuthService) {
	// Students routes (if needed with auth later)
	router.HandleFunc("/api/students/{studentID:[0-9]+}", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/students", handler.List).Methods(http.MethodGet)
	router.HandleFunc("/api/students", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/students/{studentID:[0-9]+}", handler.Update).Methods(http.MethodPut)
}
