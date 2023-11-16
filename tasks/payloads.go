package tasks

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

const (
	// TypeWelcomeEmail is a name of the task type for sending a
	//  welcome email
	TypeWelcomeEmail = "email:welcome"
	// TypeReminderEmail is a name of the task type for sending a
	//  reminder email
	TypeReminderEmail = "email:reminder"
)

type WelcomeEmail struct {
	userID int
}

type ReminderEmail struct {
	userID int
	sentIn time.Time
}

// NewWelcomeEmailTask task payload for a new welcome email
func NewWelcomeEmailTask(id int) (*asynq.Task, error) {
	// Specify task load
	payload, err := json.Marshal(WelcomeEmail{userID: id})
	if err != nil {
		return nil, err
	}

	// Return a new task with given type and payload
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}

// NewWelcomeEmailTask task payload for a new welcome email
func NewReminderEmailTask(id int, ts time.Time) (*asynq.Task, error) {
	// Specify task load
	payload, err := json.Marshal(ReminderEmail{userID: id, sentIn: ts})
	if err != nil {
		return nil, err
	}

	// Return a new task with given type and payload
	return asynq.NewTask(TypeReminderEmail, payload), nil
}
