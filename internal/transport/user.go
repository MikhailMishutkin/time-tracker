package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	"io"
	"net/http"
	"strconv"
	"time_tracker/internal/models"
)

// CreateUser
//
//	@Summary		CreateUser
//	@Description	create a user
//	@Tags			create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Employee	true	"user info"
//	@Success		200		{integer} integer 1
//	@Failure		400,404	{integer}	integer 1
//	@Failure		500		{integer}	integer 1
//	@Router			/users/create [post]
func (h *HTTPUserHandle) CreateUser(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	var user *models.Employee
	err = json.Unmarshal(content, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("corrupt json data" + err.Error()))
	}

	user, err = h.GetUserInfo(user.PassportNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("error with user's info request" + err.Error()))
	}

	err = h.hum.CreateUser(context.Background(), user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
	}
}

// ...
func (h *HTTPUserHandle) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["uuid"]
	uuid, err := strconv.Atoi(id)
	FailOnErrorsHttp(w, err, "can't convert UUID", http.StatusBadRequest)

	user := &models.Employee{UUID: uuid}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(content, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("corrupt json data" + err.Error()))
	}

	err = h.hum.UpdateUser(context.Background(), user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

	}
}

func (h *HTTPUserHandle) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["uuid"]
	uuid, err := strconv.Atoi(id)
	FailOnErrors(err, "can't convert UUID")

	user := &models.Employee{UUID: uuid}

	err = h.hum.DeleteUser(context.Background(), user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

	}
}

// Get All Users info
//
//	@Summary		GetUserInfo
//	@Description	get user info
//	@Tags			get info
//	@Accept			json
//	@Produce		json
//	@Success		200		{object} models.GetAllUsersResponse
//	@Failure		400,404	{integer}	integer 1
//	@Failure		500		{integer}	integer 1
//	@Router			/users [get]
func (h *HTTPUserHandle) GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	// TODO Errors!!
	strLimit := r.URL.Query().Get("limit")
	limit := 0
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit < 0 {
			http.Error(w, "limit query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}
	strOffset := r.URL.Query().Get("offset")
	offset := 0
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset < 0 {
			http.Error(w, "offset query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}

	filter := r.URL.Query().Get("filter")
	filterMap := map[string]string{}
	if filter != "" {
		filterMap, err = validateAndReturnFilterMap(filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	response, err := h.hum.GetAllUsersInfo(
		context.Background(),
		&models.FiltAndPagin{FilterMap: filterMap, Limit: limit, Offset: offset},
	)

	respJson := &models.GetAllUsersResponse{Response: response}

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Connection:", "close")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := json.NewEncoder(w).Encode(respJson); err != nil {
			fmt.Println(err)
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		}

	}
}
