package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time_tracker/internal/models"
)

// ...
func (r *Repo) CreateUser(ctx context.Context, user *models.EmployeeStore) error {
	const query = `
INSERT INTO users (surname, name, patronymic, address, passportSerie, passportNumber) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid
`
	tag, err := r.DB.Exec(
		ctx,
		query,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.Address,
		user.PassportSerie,
		user.PassportNumber,
	)
	// TODO
	fmt.Printf("createUser repo tag: %s", tag.String())

	return err
}

// ...
func (r *Repo) UpdateUser(ctx context.Context, user *models.EmployeeStore) error {
	log.Println("UpateUser ivoked", user)
	const query = `
UPDATE users 
SET surname= $2, name = $3, patronymic = $4, address = $5, passportSerie = $6, passportNumber = $7 
WHERE uuid = $1
`
	_, err := r.DB.Exec(
		ctx,
		query,
		user.UUID,
		user.Surname,
		user.Name,
		user.Patronymic,
		user.Address,
		user.PassportSerie,
		user.PassportNumber,
	)
	return err
}

// ...
func (r *Repo) DeleteUser(ctx context.Context, user *models.Employee) error {
	return nil
}

// ...
func (r *Repo) GetAllUsersFilter(ctx context.Context, fp *models.FiltAndPagin) (users []*models.EmployeeStore, err error) {

	query := `SELECT * FROM users` + fmt.Sprintf(" WHERE %s ", fp.Key) +
		`= $1
			LIMIT $2
			OFFSET $3
			`
	var rows pgx.Rows

	if fp.ValueInt == 0 {
		rows, err = r.DB.Query(ctx, query, fp.ValueString, fp.Limit, fp.Offset)
	} else {
		rows, err = r.DB.Query(ctx, query, fp.ValueInt, fp.Limit, fp.Offset)
	}
	// TODO
	if err != nil {
		return nil, fmt.Errorf("something wrong with get users info: ", err)
	}

	for rows.Next() {
		user := &models.EmployeeStore{}
		if err = rows.Scan(&user.UUID, &user.Surname, &user.Name, &user.Patronymic,
			&user.Address, &user.PassportSerie, &user.PassportNumber); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next: %s", err)
		}
		// TODO
		fmt.Println(user)

		users = append(users, user)
	}

	return users, err
}

// ...
func (r *Repo) GetAllUsers(ctx context.Context, fp *models.FiltAndPagin) (users []*models.EmployeeStore, err error) {
	log.Println("GetAllUsers invoked")
	query := `SELECT * FROM users 
			LIMIT $1
			OFFSET $2
			`
	rows, err := r.DB.Query(ctx, query, fp.Limit, fp.Offset)
	// TODO
	fmt.Println(rows)
	if err != nil {
		return nil, fmt.Errorf("something wrong with get users info: ", err)
	}

	for rows.Next() {
		user := &models.EmployeeStore{}
		if err = rows.Scan(&user.UUID, &user.Surname, &user.Name, &user.Patronymic,
			&user.Address, &user.PassportSerie, &user.PassportNumber); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next: %s", err)
		}
		// TODO
		fmt.Println(user)

		users = append(users, user)
	}

	// TODO
	fmt.Println(users)
	return users, err
}

func (r *Repo) CheckExist(ctx context.Context, uuid int) (exst bool, err error) {

	const query = `SELECT EXISTS (SELECT true FROM users where uuid = $1) `

	err = r.DB.QueryRow(ctx, query, uuid).Scan(&exst)

	return exst, err

}
