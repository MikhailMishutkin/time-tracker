package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time_tracker/internal/models"
)

// ...
func (s *UserService) CreateUser(ctx context.Context, user *models.Employee) error {

	ser, num, err := convertPassport(user.PassportSerie, user.PassportNumber)

	err = s.ur.CreateUser(ctx, &models.EmployeeStore{
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		PassportSerie:  ser,
		PassportNumber: num})

	return err

}

// ...
func (s *UserService) UpdateUser(ctx context.Context, user *models.Employee) error {

	exst, err := s.ur.CheckExist(ctx, user.UUID)
	if err != nil {
		return fmt.Errorf("something wrong with check user existence: %v", err)
	} else if !exst {
		return fmt.Errorf("no such user uuid: %v", err)
	}

	if user.Surname != "" && user.Name != "" && user.Patronymic != "" && user.Address != "" && user.PassportSerie != "" && user.PassportNumber != "" {
		userFromStore, err := convertUser(user)

		err = s.ur.UpdateUser(ctx, userFromStore)
		return err
	} else {
		return errors.New("empty fields in update data")
	}

	return err
}

// ...
func (s *UserService) DeleteUser(ctx context.Context, user *models.Employee) error {
	err := s.ur.DeleteUser(ctx, user)
	return err
}

// ...
func (s *UserService) GetAllUsersInfo(
	ctx context.Context,
	fp *models.FiltAndPagin,
) (
	users []*models.EmployeeStore,
	err error,
) {
	var vInt int
	for key, value := range fp.FilterMap {
		fpStr := &models.FiltAndPagin{Key: key, ValueString: value, Limit: fp.Limit, Offset: fp.Offset}
		fpInt := &models.FiltAndPagin{Key: key, ValueInt: vInt, Limit: fp.Limit, Offset: fp.Offset}
		switch {
		case key == "surname":
			users, err = s.ur.GetAllUsersFilter(ctx, fpStr)
			return users, err
		case key == "name":
			users, err = s.ur.GetAllUsersFilter(ctx, fpStr)
			return users, err
		case key == "patronymic":
			users, err = s.ur.GetAllUsersFilter(ctx, fpStr)
			return users, err
		case key == "address":
			users, err = s.ur.GetAllUsersFilter(ctx, fpStr)
			return users, err
		case key == "passport_serie":
			vInt, err = strconv.Atoi(value)
			// TODO errors
			if err != nil {
				return nil, err
			}
			fpInt.ValueInt = vInt
			users, err = s.ur.GetAllUsersFilter(ctx, fpInt)
			return users, err
		case key == "passport_number":
			vInt, err = strconv.Atoi(value)
			// TODO errors
			if err != nil {
				return nil, err
			}
			fpInt.ValueInt = vInt
			users, err = s.ur.GetAllUsersFilter(ctx, fpInt)
			return users, err
		}
	}

	users, err = s.ur.GetAllUsers(ctx, fp)
	return users, err
}
