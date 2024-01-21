package processingnode

import (
	"context"
	"fmt"
	"sync"
	"taskgenerator/taskgenerator"
	"time"
)

func RunProcessingNode(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Processing Node received cancellation signal!")
			return
		default:
			task := receiveFromQueue()
			go processTask(task)
		}
	}
}

func receiveFromQueue() taskgenerator.Task {
	task := <-taskgenerator.TaskQueue
	fmt.Printf("Task received from Queue: Task : %s\n", task.Message)
	return taskgenerator.Task{ID: task.ID, Message: task.Message}
}

func processTask(task taskgenerator.Task) {
	fmt.Printf("Processing Task : %v\n", task.Message)
	time.Sleep(2 * time.Second)
	fmt.Printf("Task Processed Successfully - Details:%v\n", task)
	time.Sleep(3 * time.Second)
}
