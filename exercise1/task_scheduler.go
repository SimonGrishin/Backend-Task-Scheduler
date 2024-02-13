package main

import (
	"fmt"
	"os"

	"github.com/Knetic/govaluate"
	"gopkg.in/yaml.v2"
)

// TaskType represents the type of a task
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
func (t *TaskType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
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

// Task represents a task to be executed
type Task struct {
	Type TaskType `yaml:"type"`
	Data string   `yaml:"data"`
}

// TasksList represents the YAML structure with a list of tasks
type TasksList struct {
	Tasks []Task `yaml:"tasks"`
}

// Read the YAML file, Unmarshal using the custom UnmarshalYAML function,
// then execute each task and output errors
func main() {
	yamlFile := "tasks.yaml"

	// Read YAML file
	fileContent, err := os.ReadFile(yamlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal YAML into TasksList
	var tasksList TasksList
	var unknownTasksErrors []error
	err = yaml.Unmarshal(fileContent, &tasksList)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
	}

	// Execute tasks
	for _, task := range tasksList.Tasks {
		err := executeTask(task)
		if err != nil {
			unknownTasksErrors = append(unknownTasksErrors, err)
		}
	}

	for _, err := range unknownTasksErrors {
		fmt.Println("Error for unknown task type:", err)
	}
}

// Handle the execution of different tasks. If adding more tasks, call those functions here in the case statement
func executeTask(task Task) error {

	if task.Data == "" {
		fmt.Println()
	}

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
