package repository

import (
	"context"
	"fmt"
	"log"
	"time_tracker/internal/models"
)

// ...
func (r *Repo) GetTasKCost(ctx context.Context, user *models.Employee) (*models.Task, error) {

	return nil, nil
}

// ...
func (r *Repo) StartTask(ctx context.Context, start *models.Employee) error {
	const query = `INSERT INTO worktime (start, task_id, user_uuid) VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(ctx, query, start.Task.WorkTime.Start, start.Task.UUID, start.UUID)
	if err != nil {
		fmt.Errorf("something wrong with write start task data to db: %v", err)
	}
	return err
}

// ...
func (r *Repo) FinishTask(ctx context.Context, user *models.Employee) error {
	const query = `UPDATE worktime SET finish = $1
				WHERE task_id = $2 AND user_uuid =$3`

	_, err := r.DB.Exec(ctx, query, user.Task.WorkTime.Finish, user.Task.UUID, user.UUID)
	if err != nil {
		fmt.Errorf("something wrong with write finish task data to db: %v", err)
	}
	return err
}

func (r *Repo) CreateTask(ctx context.Context, title string) (id int, err error) {
	const query = `WITH t AS(
		INSERT INTO task (title, active)
			VALUES ($1, $2)
    	ON CONFLICT (title) DO NOTHING  
   		 RETURNING id)
		SELECT * FROM t
		UNION
		SELECT id FROM task WHERE title=$1`
	var active bool = true
	err = r.DB.QueryRow(ctx, query, title, active).Scan(&id)

	return id, err
}

func (r *Repo) FindActiveTask(ctx context.Context, userUUID int) (*models.Task, error) {
	const query = `SELECT task_id FROM task, worktime
					WHERE worktime.user_uuid = $1 AND task.active = true`
	var taskId int
	err := r.DB.QueryRow(ctx, query, userUUID).Scan(&taskId)

	return &models.Task{UUID: taskId}, err
}

func (r *Repo) SetActive(ctx context.Context, task *models.Task) error {
	log.Println("SetActive invoked")
	const query = `UPDATE task SET active = $1
				WHERE id = $2`
	log.Println(task.Active, task.UUID)
	tag, err := r.DB.Exec(ctx, query, task.Active, task.UUID)
	log.Println(tag)
	if err != nil {
		fmt.Errorf("can't set active: %v", err)
	}
	return err
}
