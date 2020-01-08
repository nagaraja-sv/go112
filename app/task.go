package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/julienschmidt/httprouter"
)

type Task struct {
	Description string
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "simplystpharma"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	//datastore.Key("Taskkey")

	// Sets the kind for the new entity.
	kind := "Task"
	// Sets the name/ID for the new entity.
	name := "sampletask12"
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)
	taskKey.Namespace = "narendar"

	// Creates a Task instance.
	task := Task{
		Description: "Hello World!",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Printf("Saved %v: %v\n", taskKey, task.Description)
}
