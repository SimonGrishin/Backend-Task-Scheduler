// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type TaskType int

const (
	// PrintTask represents a task to print a message
	PrintTask TaskType = iota
	// ComputeTask represents a task to compute a simple arithmetic expression
	ComputeTask
	// Other represents any unknown task
	Other
)

// UnmarshalYAML custom unmarshaling for TaskType
func (t *TaskType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	switch str {
	case "PrintTask":
		*t = PrintTask
	case "ComputeTask":
		*t = ComputeTask
	default:
		*t = Other
		// return fmt.Errorf("unknown task type: %s", str)
	}

	return nil
}

// Task struct represents a task
type Task struct {
	ID     string   `json:"ID"`
	Type   TaskType `json:"Type"`
	Data   string   `json:"Data"`
	Status string   `json:"Status"`
}

// TasksList represents the YAML structure with a list of tasks
// albums slice to seed record album data.

// var tasks []Task

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// {
// 	{ID: "1", Type: PrintTask, Data: "John Coltrane", Status: "200 Ok"},
// 	{ID: "2", Type: ComputeTask, Data: "1+1", Status: "404 Page not found"},
// 	{ID: "3", Type: PrintTask, Data: "Sarah Vaughan", Status: "202 Accepted"},
// }

// Config represents server configuration
type Config struct {
	Port int `yaml:"port"`
}

var tasks Tasks
var unknownTasksErrors []error

func main() {

	jsonFile := "tasks.json"

	b, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Unable to read file due to %s\n", err)
	}

	err = json.Unmarshal(b, &tasks)
	if err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	// Load configuration from YAML file
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	// fmt.Println("config worked")

	// Set up Gin router
	router := gin.Default()

	// Define routes
	router.POST("/tasks", addTask)
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	// Run the server
	addr := fmt.Sprintf("localhost:%d", config.Port)
	router.Run(addr)

	for _, err := range unknownTasksErrors {
		fmt.Println("Error for unknown task type:", err)
	}
}

// loadConfig reads configuration from a YAML file
func loadConfig(filename string) (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// addTask adds a new task to the in-memory storage
func addTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique ID (you may use a library for this in production)
	newTask.ID = fmt.Sprintf("%d", rand.Intn(99999)+100000)
	newTask.Status = fmt.Sprintf("%d", http.StatusCreated)

	// Perform task operation based on its type (this is a placeholder)
	err := executeTask(newTask)
	if err != nil {
		unknownTasksErrors = append(unknownTasksErrors, err)
	}

	// Add the task to the list
	tasks.Tasks = append(tasks.Tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

// getTasks returns all tasks

// getAlbums responds with the list of all albums as JSON.
func getTasks(c *gin.Context) {

	for _, task := range tasks.Tasks {
		err := executeTask(task)
		if err != nil {
			unknownTasksErrors = append(unknownTasksErrors, err)
		}
	}

	for _, err := range unknownTasksErrors {
		fmt.Println("Error for unknown task type:", err)
	}

	c.IndentedJSON(http.StatusOK, tasks)

}

// func getTasks(c *gin.Context) {
// 	c.JSON(http.StatusOK, tasks)
// }

// getTaskByID returns a task by ID
func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id: ", id)

	task, err := findTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// updateTask updates a task by ID
func updateTask(c *gin.Context) {
	id := c.Param("id")

	task, err := findTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = executeTask(*task)
	if err != nil {
		unknownTasksErrors = append(unknownTasksErrors, err)
	}

	// Perform task operation based on its type (this is a placeholder)
	// performTaskOperation(task)

	c.JSON(http.StatusOK, task)
}

// deleteTask deletes a task by ID
func deleteTask(c *gin.Context) {
	id := c.Param("id")

	task, err := findTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	fmt.Println("the task found by ID is: ", task)

	// Perform task operation based on its type (this is a placeholder)
	// performTaskOperation(task)

	// Remove the task from the list
	tasks.Tasks = removeTaskByID(id)

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// findTaskByID finds a task by ID
func findTaskByID(id string) (*Task, error) {
	for _, task := range tasks.Tasks {
		if task.ID == id {
			return &task, nil
		}
		// fmt.Printf("task id: %s and the id we are looking for: %s", task.ID, id)
	}
	return nil, fmt.Errorf("task not found")
}

// removeTaskByID removes a task by ID
func removeTaskByID(id string) []Task {
	var updatedTasks []Task
	for _, task := range tasks.Tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}
	return updatedTasks
}

// performTaskOperation is a placeholder for the actual task operation based on its type
func performTaskOperation(task Task) {
	// Perform the task operation based on the task type

	err := executeTask(task)
	if err != nil {
		unknownTasksErrors = append(unknownTasksErrors, err)
	}

	fmt.Printf("Performing operation for task with ID %s\n", task.ID)
}

func executeTask(task Task) error {

	switch task.Type {
	case PrintTask:
		fmt.Println("PrintTask:", task.Data)
	case ComputeTask:
		result, err := computeExpression(task.Data)
		if err != nil {
			return fmt.Errorf("Error computing expression: %w", err)
		}
		fmt.Println("ComputeTask result:", result)
	case Other:
		return fmt.Errorf("Unknown task type with data: '%s' ", task.Data)

	default:
		return fmt.Errorf("Unknown task type for data: %s", task.Data)
	}
	return nil
}

// Function takes a string expression and converts to mathematical expression
// and returns the result of computation
func computeExpression(expr string) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return "", err
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return result, nil
}
