package students

import (
	"lesson-management/pkg/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(api *mux.Router, handler handler.Handler) {
	api.HandleFunc("/students/{studentID}", handler.Get).Methods(http.MethodGet)
	api.HandleFunc("/students/list", handler.List).Methods(http.MethodGet)
	api.HandleFunc("/students", handler.Create).Methods(http.MethodPost)
	api.HandleFunc("/students", handler.Update).Methods(http.MethodPut)
}
