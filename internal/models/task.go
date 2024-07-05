package models

import "time"

type Task struct {
	UUID     int       `json:"UUID,omitempty"`
	Title    string    `json:"title"`
	Period   *Period   `json:"period,omitempty"`
	TimeCost time.Time `json:"timeCost,omitempty"`
	WorkTime *WorkTime `json:"workTime"`
	Active   bool      `json:"active"`
}

type WorkTime struct {
	Start  time.Time `json:"start,omitempty"`
	Finish time.Time `json:"finish,omitempty"`
}

type Period struct {
	FirstDate  time.Time `json:"firstDate"`
	SecondDate time.Time `json:"secondDate"`
}
