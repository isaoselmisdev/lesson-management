package lessons

import (
	"lesson-management/pkg/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(api *mux.Router, handler handler.Handler) {
	api.HandleFunc("/lessons/{lessonID:[0-9]+}", handler.Get).Methods(http.MethodGet)
	api.HandleFunc("/lessons/list", handler.List).Methods(http.MethodGet)
	api.HandleFunc("/lessons", handler.Create).Methods(http.MethodPost)
	api.HandleFunc("/lessons/{lessonID}", handler.Update).Methods(http.MethodPut)
}
