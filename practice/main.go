package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

type TaskType int

const (

	// Set the PrintTask constant to be referenced
	PrintTask TaskType = iota

	// Set the ComputeTask constant referenced by incrementing PrintTask
	ComputeTask
)

type Task struct {
	Type TaskType `yaml:"type"`
	Data string   `yaml:"data"`
}

//Configuration represents the YAML configuration

type Config struct {
	Tasks []Task `yaml:"tasks"`
}

func (tl *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Define an intermediate struct to capture the tasks array
	var intermediate struct {
		Tasks []Task `yaml:"tasks"`
	}

	// Unmarshal into the intermediate struct
	if err := unmarshal(&intermediate); err != nil {
		return err
	}

	// Assign the tasks to the TaskList
	tl.Tasks = intermediate.Tasks

	return nil
}

func readConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error while reading %v", err)
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal([]byte(file), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// func customUnmarshal(unmarshal func(interface{}) error) error {

// 	var typeString string
// 	if err := unmarshal(&typeString); err != nil {
// 		return err
// 	}

// 	switch typeString {
// 	case "PrintTask":
// 		return setValue(PrintTask, node)
// 	case "ComputeTask":
// 		return setValue(ComputeTask, node)
// 	default:
// 		return fmt.Errorf("unknown task type: %s", typeString)
// 	}
// }

// func setValue(value TaskType, node *yaml.Node) error {
// 	node.Kind = yaml.Sc
// 	node.Style = yaml.DoubleQuotedStyle
// 	node.Tag = "!!str"
// 	node.Value = fmt.Sprintf("%d", value)
// 	return nil
// }

func main() {

	// p := Person{"John", 30}
	// y, err := yaml.Marshal(p)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// fmt.Println(y)
	// fmt.Println(string(y))
	// // name: John
	// // age: 30

	// err = os.WriteFile("file1.yaml", y, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// arithmetic := "1+1"

	// fmt.Println(arithmetic)

	// reading yaml file and calling readConfig
	filename := "tasks.yaml"

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error while reading %v", err)
	}

	fmt.Printf("The file data is: %v\n", string(file))

	// Unmarshal YAML into TaskList
	var taskList Config
	err = yaml.Unmarshal(file, &taskList)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}

	// Print tasks
	for _, task := range taskList.Tasks {
		fmt.Printf("Task Name: %b\n", task.Type)
		// Print other task fields as needed
	}

	var config Config

	// var intermediate struct {
	// 	Tasks []Task `'yaml:"tasks"`
	// }

	// if err = unmarshal(&intermediate); err != nil {
	// 	fmt.Println("Error Unmarshalling:", err)
	// }

	// config, err := readConfig(configFileName)

	// if err != nil {
	// 	log.Fatalf("Error reading configuration: %v", err)
	// }

	fmt.Println("the data is: ", &config)
	//fmt.Printf("\nThe  data is %v and the type is %T\n", )

}

// Implement custom UnmarshalYAML method for TaskList
