package main

import (
	"log"
	"math/rand"
	"time"

	"golangtasks/tasks"

	"github.com/hibiken/asynq"
)

func main() {
	// create a new redis connection for the client
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379", // Redis server address
	}

	//  Create a new Asynq client
	client := asynq.NewClient(redisConnection)
	defer client.Close()

	// infinite loop to create tasks as Asynq client.
	for {
		// Generate a random user ID
		userID := rand.Intn(1000) + 10

		// Set a delay duration to 2 minutes.
		delay := 2 * time.Minute

		//Define tasks
		task1, err := tasks.NewWelcomeEmailTask(userID)
		if err != nil {
			log.Fatalf("Could not create task: %s", err)
		}
		task2, err := tasks.NewReminderEmailTask(userID, time.Now().Add(delay))
		if err != nil {
			log.Fatalf("Could not create task: %s", err)
		}

		// Process the task immediately in critical queue
		if _, err := client.Enqueue(
			task1,                   // task payload
			asynq.Queue("critical"), // set queue for the task
		); err != nil {
			log.Fatal(err)
		}

		// Process the task 2 minutes later in low queue
		if _, err := client.Enqueue(
			task2,                  // task payload
			asynq.Queue("low"),     // set queue for the task
			asynq.ProcessIn(delay), // set time to process task
		); err != nil {
			log.Fatal(err)
		}
	}

}
