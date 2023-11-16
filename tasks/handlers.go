package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

// HandleWelcomeEmailTask  handler for welcome email task
func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var p WelcomeEmail
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("Json unmarshal failed: %v, %w", err, asynq.SkipRetry)
	}

	// Dummy message to the workers output
	log.Printf("Send welcome email to User ID %d", p.userID)
	return nil
}

// HandleWelcomeEmailTask  handler for welcome email task
func HandleReminderEmailTask(ctx context.Context, t *asynq.Task) error {
	var p ReminderEmail
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("Json unmarshal failed: %v, %w", err, asynq.SkipRetry)
	}

	// Dummy message to the workers output
	log.Printf("Send welcome email to User ID %d", p.userID)
	log.Printf("Reason: time is up (%v)\n", p.sentIn)
	return nil
}
