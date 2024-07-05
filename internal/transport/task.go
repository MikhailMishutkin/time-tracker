package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"time_tracker/internal/models"
)

// TODO обработка ошибок!
func (h *HTTPTaskHandle) StartTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["uuid"]
	uuid, err := strconv.Atoi(id)
	FailOnErrorsHttp(w, err, "can't convert UUID", http.StatusBadRequest)

	content, err := io.ReadAll(r.Body)
	FailOnErrorsHttp(w, err, "can't convert UUID", http.StatusBadRequest)

	start := &models.Employee{
		UUID: uuid,
	}

	err = json.Unmarshal(content, &start)
	FailOnErrorsHttp(w, err, "corrupt json data", http.StatusBadRequest)

	err = h.htm.StartTask(context.Background(), start)
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
func (h *HTTPTaskHandle) FinishTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["uuid"]
	uuid, err := strconv.Atoi(id)
	FailOnErrorsHttp(w, err, "can't convert UUID", http.StatusBadRequest)

	content, err := io.ReadAll(r.Body)
	FailOnErrorsHttp(w, err, "can't convert UUID", http.StatusBadRequest)

	finish := &models.Employee{
		UUID: uuid,
	}

	err = json.Unmarshal(content, &finish)
	FailOnErrorsHttp(w, err, "corrupt json data", http.StatusBadRequest)

	err = h.htm.FinishTask(context.Background(), finish)
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
func (h *HTTPTaskHandle) GetTaskCost(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTaskCost ivoked")
	req := mux.Vars(r)

	UUID, err := strconv.Atoi(req["UUID"])
	FailOnErrorsHttp(w, err, "can't convert uuid", http.StatusBadRequest)

	dateString := req["firstdate"] // must be in format "2021-11-22"
	firstDate, err := time.Parse("2006-01-02", dateString)
	FailOnErrorsHttp(w, err, "can't convert parse first time value", http.StatusBadRequest)

	dateString = req["seconddate"]
	secondDate, err := time.Parse("2006-01-02", dateString)
	FailOnErrorsHttp(w, err, "can't convert parse second time value", http.StatusBadRequest)

	user := &models.Employee{
		UUID:    UUID,
		Surname: req["surname"],
		Task: &models.Task{
			Period: &models.Period{
				FirstDate:  firstDate,
				SecondDate: secondDate,
			},
		},
	}

	a, err := h.htm.GetTasKCost(context.Background(), user)
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
		if err := json.NewEncoder(w).Encode(a); err != nil {
			fmt.Println(err)
			http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		}
	}
}
