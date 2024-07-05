package transport

import (
	"net/url"
	"strings"
	"time_tracker/internal/models"
)

func (h *HTTPUserHandle) GetUserInfo(s string) (user *models.Employee, err error) {

	passport := strings.Fields(s)

	params := url.Values{}
	params.Add("passportSerie", passport[0])
	params.Add("passportNumber", passport[1])

	//url := "https://example.com/info?" + params.Encode()

	//ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	//defer cancel()
	//req, err := http.NewRequestWithContext(
	//	ctx,
	//	"GET",
	//	url,
	//	nil,
	//)
	//FailOnErrors(err, "error with NewRequestWithContext")
	//
	//_, err = http.DefaultClient.Do(req)
	//FailOnErrors(err, "error when executing the request to API")

	//content, err := io.ReadAll(response.Body)
	//FailOnErrors(err, "fail to read response")

	//err = json.Unmarshal(content, &user)
	//FailOnErrors(err, "corrupt json data")

	// TODO
	//user.PassportSerie = passport[0]
	//user.PassportNumber = passport[1]

	return &models.Employee{
			Surname:        "Иванов",
			Name:           "Иван",
			Patronymic:     "Иванович",
			Address:        "г. Москва, Кремль, стр. 1",
			PassportSerie:  passport[0],
			PassportNumber: passport[1],
		},
		nil
}
