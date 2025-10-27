package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, handler *AuthHandler) {
	router.HandleFunc("/api/auth/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/register/admin", handler.RegisterAdmin).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/register/teacher", handler.RegisterTeacher).Methods(http.MethodPost)
}
