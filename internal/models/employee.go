package models

type Employee struct {
	UUID           int    `json:"uuid,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Name           string `json:"name,omitempty"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"address,omitempty"`
	PassportSerie  string `json:"passportSerie,omitempty"`
	PassportNumber string `json:"passportNumber,omitempty"`
	Task           *Task  `json:"task,omitempty"`
}

type EmployeeStore struct {
	UUID           int    `json:"uuid,omitempty"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	PassportSerie  int    `json:"passportSerie"`
	PassportNumber int    `json:"passportNumber"`
}

type GetAllUsersResponse struct {
	Response []*EmployeeStore `json:"response"`
}
