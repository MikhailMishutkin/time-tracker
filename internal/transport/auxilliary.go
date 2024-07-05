package transport

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time_tracker/internal/models"
)

// TODO: delete?
func FailOnErrors(err error, message string) {
	if err != nil {
		log.Printf("%s: %s", err, message)
	}
}

func FailOnErrorsHttp(w http.ResponseWriter, err error, message string, status int) {
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(message + err.Error()))
	}
}

var userFields = getUserFields()

func getUserFields() []string {
	var field []string

	v := reflect.ValueOf(models.Employee{})
	for i := 0; i < v.Type().NumField(); i++ {
		value := v.Type().Field(i).Tag.Get("json")
		splits := strings.Split(value, ",")
		field = append(field, splits[0])

	}

	return field
}

func validateAndReturnFilterMap(filter string) (map[string]string, error) {
	splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return nil, errors.New("malformed sortBy query parameter, should be field.value")
	}
	field, value := splits[0], splits[1]
	if !stringInSlice(userFields, field) {
		return nil, errors.New("unknown field in filter query parameter")
	}
	return map[string]string{field: value}, nil
}

func stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}

	return false
}
