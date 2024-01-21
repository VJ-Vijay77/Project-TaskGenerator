package taskgenerator

import (
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
)

// Task represents the structure of a task
type Task struct {
	ID      uint32
	Message string
}

var (
	TaskQueue  = make(chan Task) //unbuffered channel
	taskID     = uuid.New().ID()
	taskNumber int
)

func RunGenerateTasks(ctx context.Context) {
	// creating an infinite loop to generate tasks
	for {
		select{
		case <- ctx.Done():
			fmt.Println("Task Generator received Cancellation Signal!")
			return
		default:
			task := generateTask()
			sendTaskToQueue(task)
			time.Sleep(3 * time.Second)
		}
	}
}

func generateTask() Task {
	// incrementing TaskID so IDs of each task will a chronolgical order
	taskID++
	taskNumber++
	return Task{ID: taskID, Message: fmt.Sprintf("Task No : %d\tTask ID : %v", taskNumber, taskID)}
}

func sendTaskToQueue(task Task) {
	TaskQueue <- task
	fmt.Printf("\nTask send to the queue : %s\n", task.Message)
}
