package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Task struct {
	Id              int       `json:"id"`
	TaskName        string    `json:"taskname"`
	TaskDescription string    `json:"taskdescription"`
	TaskStatus      string    `json:"taskstatus"`
	TaskStartDate   time.Time `json:"startdate"`
	TaskEndDate     time.Time `json:"enddate"`
}

func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	aux := &struct {
		TaskStartDate string `json:"startdate"`
		TaskEndDate   string `json:"enddate"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the startdate
	startDate, err := time.Parse("2006-01-02", aux.TaskStartDate)
	if err != nil {
		return fmt.Errorf("invalid startdate format: %v", err)
	}
	t.TaskStartDate = startDate

	// Parse the enddate
	endDate, err := time.Parse("2006-01-02", aux.TaskEndDate)
	if err != nil {
		return fmt.Errorf("invalid enddate format: %v", err)
	}
	t.TaskEndDate = endDate

	return nil
}
