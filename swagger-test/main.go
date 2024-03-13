package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "friends2meet/swagger-test/docs"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

// @title 			swagger-task-scheduler-api
// @version 		1.0
// @description 	This is a simple API documentation example that takes a GIN API using mongoDB in GO to handle API calls
// @termsOfService 	http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					Security feature to test API-keys

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

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

	r := gin.Default()

	// add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.POST("tasks", addTask)
		v1.GET("tasks", getTasks)
		v1.GET("tasks/:id", getTaskByID)
		v1.PUT("tasks/:id", updateTask)
		v1.DELETE("tasks/:id", deleteTask)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}

// loadConfig loads configuration settings from the specified file.
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

// @Summary Add a new task
// @Description Add a new task to the database
// @Accept json
// @Produce json
// @Param task body Task true "Task object"
// @Security ApiKeyAuth
// @Tags Tasks
// @Success 201 {object} Task
// @Failure 404 "task not found"
// @Router /tasks [post]
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

// @Summary Get all tasks
// @Description Retrieve all tasks from the database
// @Produce json
// @Tags Tasks
// @Success 200 {array} Task
// @Router /tasks [get]
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

// @Summary Get a task by ID
// @Description Retrieve a task from the database by its ID
// @Produce json
// @Param id path string true "Task ID"
// @Tags Tasks
// @Success 200 {object} Task
// @Router /tasks/{id} [get]
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

// @Summary Update a task
// @Description Update a task in the database by its ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body Task true "Updated task object"
// @Tags Tasks
// @Success 200 {object} Task
// @Router /tasks/{id} [put]
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

// @Summary Delete a task
// @Description Delete a task from the database by its ID
// @Produce json
// @Param id path string true "Task ID"
// @Tags Tasks
// @Success 200 {string} string "Task deleted successfully"
// @Router /tasks/{id} [delete]
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
