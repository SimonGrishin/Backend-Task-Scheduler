package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

type TaskType int

const (
	PrintTask TaskType = iota
	ComputeTask
	Other
)

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
	}

	return nil
}

// type Task struct {
// 	ID     TaskID   `json:"_id" bson:"_id"`
// 	Data   string   `json:"data" bson:"data"`
// 	Type   TaskType `json:"type" bson:"type"`
// 	Status string   `json:"status" bson:"status"`
// }

type Task struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Data   string             `json:"data" bson:"data"`
	Type   string             `json:"type" bson:"type"`
	Status string             `json:"status" bson:"status"`
}

var tasks []Task

type Config struct {
	Port             int    `yaml:"port"`
	ConnectionString string `yaml:"connection_string"`
	DatabaseName     string `yaml:"database_name"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
}

var tasksCollection *mongo.Collection

func main() {
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	uri := fmt.Sprintf(config.ConnectionString, config.Username, config.Password)

	fmt.Println(config.DatabaseName)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to client ")
	defer client.Disconnect(context.TODO())

	tasksCollection = client.Database("task_scheduler").Collection("tasks")

	router := gin.Default()
	router.POST("/tasks", addTask)
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	addr := fmt.Sprintf("localhost:%d", config.Port)
	router.Run(addr)

}

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

// unused function
func connectDB(uri string) (*mongo.Client, error) {

	// Connect to your Atlas cluster
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to client... ")
	defer client.Disconnect(context.TODO())

	return client, nil
}

func addTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = primitive.NewObjectID()
	newTask.Status = "201 Created"

	_, err := tasksCollection.InsertOne(context.Background(), newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating task"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// Getting All tasks

func getTasks(c *gin.Context) {

	cursor, err := tasksCollection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tasks"})
		return
	}
	defer cursor.Close(context.Background())

	var tasks []Task
	err = cursor.All(context.Background(), &tasks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding tasks"})
		return
	}

	// for _, task := range tasks {
	// 	err := executeTask(task)
	// 	if err != nil {
	// 		unknownTasksErrors = append(unknownTasksErrors, err)
	// 	}
	// }

	c.IndentedJSON(http.StatusOK, tasks)
}

// Getting Task by ID

func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var task Task
	err = tasksCollection.FindOne(context.Background(), bson.D{{"_id", objectID}}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = tasksCollection.UpdateOne(context.Background(), bson.D{{"_id", objectID}}, bson.D{{"$set", updatedTask}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating task"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = tasksCollection.DeleteOne(context.Background(), bson.D{{"_id", objectID}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func executeTask(task Task) error {

	switch task.Type {
	case "PrintTask":
		fmt.Println("PrintTask:", task.Data)
	case "ComputeTask":
		result, err := computeExpression(task.Data)
		if err != nil {
			return fmt.Errorf("Error computing expression: %w", err)
		}
		fmt.Println("ComputeTask result:", result)
	case "Other":
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
