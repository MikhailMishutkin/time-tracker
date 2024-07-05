package transport

import (
	"context"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"time_tracker/internal/models"
)

type HTTPUserHandle struct {
	hum HTTPUserManager
}

type HTTPTaskHandle struct {
	htm HTTPTaskManager
}

func NewHTTPUserHandle(hum HTTPUserManager) *HTTPUserHandle {
	return &HTTPUserHandle{
		hum: hum,
	}
}

func NewHTTPTaskHandle(htm HTTPTaskManager) *HTTPTaskHandle {
	return &HTTPTaskHandle{htm: htm}
}

type HTTPUserManager interface {
	CreateUser(context.Context, *models.Employee) error
	UpdateUser(context.Context, *models.Employee) error
	DeleteUser(context.Context, *models.Employee) error
	GetAllUsersInfo(context.Context, *models.FiltAndPagin) ([]*models.EmployeeStore, error)
}
type HTTPTaskManager interface {
	GetTasKCost(context.Context, *models.Employee) (*models.Task, error)
	StartTask(context.Context, *models.Employee) error
	FinishTask(context.Context, *models.Employee) error
}

func (h *HTTPUserHandle) RegisterUser(router *mux.Router) {
	router.HandleFunc("/users/create", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/edit/{uuid}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/delete/{uuid}", h.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", h.GetUsers).Methods("GET")

}

func (h *HTTPTaskHandle) RegisterTask(router *mux.Router) {
	router.HandleFunc("/task/time-cost", h.GetTaskCost).Methods(
		"Get",
	).Queries("uuid", "{uuid}", "surname", "{surname}", "firstdate", "{firstdate}", "seconddate", "{seconddate}")
	router.HandleFunc("/task/start/{uuid}", h.StartTask).Methods("POST")
	router.HandleFunc("/task/finish/{uuid}", h.FinishTask).Methods("POST")
}
