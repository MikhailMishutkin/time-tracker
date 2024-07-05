package service

import (
	"context"

	"time_tracker/internal/models"
)

type UserService struct {
	ur UserRepositorier
}

type TaskService struct {
	tr TaskRepositorier
}

func NewUserService(ur UserRepositorier) *UserService {
	return &UserService{ur: ur}
}

func NewTaskService(tr TaskRepositorier) *TaskService {
	return &TaskService{tr: tr}
}

type UserRepositorier interface {
	CreateUser(context.Context, *models.EmployeeStore) error
	UpdateUser(context.Context, *models.EmployeeStore) error
	DeleteUser(context.Context, *models.Employee) error
	GetAllUsersFilter(context.Context, *models.FiltAndPagin) ([]*models.EmployeeStore, error)
	GetAllUsers(context.Context, *models.FiltAndPagin) ([]*models.EmployeeStore, error)
	CheckExist(context.Context, int) (bool, error)
}

type TaskRepositorier interface {
	GetTasKCost(context.Context, *models.Employee) (*models.Task, error)
	StartTask(context.Context, *models.Employee) error
	FinishTask(context.Context, *models.Employee) error
	CreateTask(context.Context, string) (int, error)
	FindActiveTask(context.Context, int) (*models.Task, error)
	SetActive(ctx context.Context, task *models.Task) error
}
