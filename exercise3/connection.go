package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ConnectionString string `yaml:"connectionString"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	Port             int    `yaml:"port"`
}

type TaskType int

const (
	// PrintTask represents a task to print a message
	PrintTask TaskType = iota
	// ComputeTask represents a task to compute a simple arithmetic expression
	ComputeTask
	// Other represents any unknown task
	Other
)

type Task struct {
	ID     primitive.ObjectID
	Type   TaskType
	Data   string
	Status string
}

// Task struct represents a task
// type Task struct {
// 	ID     primitive.ObjectID `json:"ID" bson:"_id"`
// 	Type   TaskType           `json:"type" bson:"type"`
// 	Data   string             `json:"data" bson:"data"`
// 	Status string             `json:"status" bson:"status"`
// }

var tasks []Task

func main() {

	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	uri := fmt.Sprintf(config.ConnectionString, config.Username, config.Password)

	fmt.Println("Config: ", config.Username)
	fmt.Println("Password: ", config.Password)

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection := client.Database("task_scheduler").Collection("tasks")

	newTask := Task{
		ID:     primitive.NewObjectID(),
		Type:   TaskType(PrintTask),
		Data:   "Jelly",
		Status: "200 OK",
	}
	// Insert the document into the specified collection
	collection.InsertOne(context.TODO(), newTask)
	// Find and return the document
	collection = client.Database("task_scheduler").Collection("tasks")
	filter := bson.D{{"_id", primitive.ObjectID("65bc31969981f64f4f66cbb1")}}
	var result Task
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document Found:\n%+v\n", result)
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

func connectDB(connectionString, dbName string) (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}
