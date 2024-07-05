package service

import (
	"log"
	"strconv"
	"time_tracker/internal/models"
)

func FailOnErrors(err error, message string) {
	if err != nil {
		log.Printf("%s: %s", err, message)
	}
}

func convertPassport(serField, numField string) (ser, num int, err error) {
	ser, err = strconv.Atoi(serField)
	FailOnErrors(err, "can't convert passportSerie")
	num, err = strconv.Atoi(numField)
	FailOnErrors(err, "can't convert passportNumber")

	return ser, num, err
}

func convertUser(user *models.Employee) (userFromStore *models.EmployeeStore, err error) {
	ser, num, err := convertPassport(user.PassportSerie, user.PassportNumber)

	userFromStore = &models.EmployeeStore{
		UUID:           user.UUID,
		Surname:        user.Surname,
		Name:           user.Name,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
		PassportSerie:  ser,
		PassportNumber: num,
	}

	return userFromStore, err
}
