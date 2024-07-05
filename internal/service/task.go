package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"time_tracker/internal/models"
)

//сделать цикл, где по значениям дат в периоде, извлекать старт и финиш, плюсовать

// ...
func (s *TaskService) GetTasKCost(ctx context.Context, user *models.Employee) (*models.Task, error) {
	return nil, nil
}

// ...
func (s *TaskService) StartTask(ctx context.Context, user *models.Employee) (err error) {
	task, err := s.tr.FindActiveTask(ctx, user.UUID)
	log.Println(err, task)

	var ErrNoRows = errors.New("no rows in result set")

	if err != nil && err.Error() == ErrNoRows.Error() {
		id, err := s.tr.CreateTask(ctx, user.Task.Title)

		start := time.Now()
		//time.Sleep(50 * time.Millisecond)
		FailOnErrors(err, "can't start task")

		err = s.tr.StartTask(
			ctx,
			&models.Employee{
				UUID: user.UUID,
				Task: &models.Task{
					UUID: id,
					WorkTime: &models.WorkTime{
						Start: start.Round(time.Minute),
					}}})
		if err != nil {
			return fmt.Errorf("can't start task: %v", err)
		}

		err = s.tr.SetActive(ctx, &models.Task{
			UUID:   id,
			Active: true,
		})
		if err != nil {
			return fmt.Errorf("can't set active to true: %v", err)
		}

	} else {
		return fmt.Errorf("user has an active task %v, finish it", task.UUID)
	}

	return nil
}

// ...
func (s *TaskService) FinishTask(ctx context.Context, user *models.Employee) error {
	task, err := s.tr.FindActiveTask(ctx, user.UUID)
	if err != nil {
		return fmt.Errorf("user has no active task, nothing to finish")
	}

	finish := time.Now()

	err = s.tr.FinishTask(ctx, &models.Employee{
		UUID: user.UUID,
		Task: &models.Task{
			UUID: task.UUID,
			WorkTime: &models.WorkTime{
				Finish: finish.Round(time.Minute),
			}}})
	if err != nil {
		return fmt.Errorf("can't finish task: %v", err)
	}

	err = s.tr.SetActive(ctx, &models.Task{
		UUID:   task.UUID,
		Active: false,
	})
	if err != nil {
		return fmt.Errorf("can't set active to false: %v", err)
	}

	return err

}
