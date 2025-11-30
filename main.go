package main

// package
import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// ------------------
// Task Schema
// ------------------
type Task struct {
	ID        int       `json:"id"`
	Desc      string    `json:"desc"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

var tasks Tasks
var nextID int = 1
var taskPath = "tasks.json"

// ------------------
// InitNextID
// ------------------
func initNextID() {
	maxID := 0
	for _, task := range tasks.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	nextID = maxID + 1
}

// ------------------
// Insert Task
// ------------------
func insert_task(desc string) {
	newTask := Task{
		ID:        nextID,
		Desc:      desc,
		Status:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tasks.Tasks = append(tasks.Tasks, newTask)
	nextID++
}

// -------------------
// Load Tasks
// -------------------
func load_tasks(taskFilePath string) {
	taskFile, err := os.Open(taskFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return

	}
	defer taskFile.Close()
	fmt.Println("Successfully Opened users.json")

	byteValue, err := io.ReadAll(taskFile)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if err := json.Unmarshal(byteValue, &tasks); err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	initNextID() // find the nextID
}

// ---------------------
// Save On Json
// ---------------------
func saveTasks() {
	// 0. Create the json file
	file, err := os.Create(taskPath)
	if err != nil {
		fmt.Println("Error Create file:", err)
		return
	}
	defer file.Close()
	// 1. Convert from go object to json
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	// 2. save to file
	if _, err := file.Write(jsonData); err != nil {
		fmt.Println("Error Writing to file: ", err)
	} else {
		fmt.Println("Tasks saved Successfully")
	}

}

// ---------------------
// Read Task
// ---------------------
func readTasks() {

	fmt.Println("----------------------")
	for _, task := range tasks.Tasks {
		fmt.Println("Task ID: ", task.ID)
		fmt.Println("Task Desc: ", task.Desc)
		fmt.Println("Task Status: ", task.Status)
		fmt.Println("----------------------")
	}
	fmt.Println("----------------------")
}

// -----------------------
// Delete Eelement By ID
// -----------------------
func deleteByID(targetID int) {
	newTasks := tasks.Tasks[:0]
	for _, task := range tasks.Tasks {
		if task.ID != targetID {
			newTasks = append(newTasks, task)
		}
	}
	tasks.Tasks = newTasks
}

// ---------------------
// Update Task
// ---------------------
func udpateTask(targetID int, mode int, value interface{}) {
	// find task first
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == targetID {
			switch mode {
			case 0: // update description
				decs, ok := value.(string)
				if ok {
					tasks.Tasks[i].Desc = decs
					tasks.Tasks[i].UpdatedAt = time.Now()
				} else {
					fmt.Println("Invaild value for description")
				}
			case 1: // update status
				status, ok := value.(int)
				if ok {
					if status == 0 {
						tasks.Tasks[i].Status = false
					} else {
						tasks.Tasks[i].Status = true
					}
					tasks.Tasks[i].UpdatedAt = time.Now()
				} else {
					fmt.Println("Invaild value for status")
				}
			default:
				fmt.Println("Unknown update mode")
			}
			return
		}
	}
}

// ---------------------
// Main Function
// ---------------------
func main() {
	// 1. load tasks
	load_tasks(taskPath)

	// 2. Insert New Task
	//insert_task("Go To Gym")
	//insert_task("Play Football")

	// 3. Delete Task
	//deleteByID(1)

	// 4. Update Task
	//udpateTask(2, 0, "Write integeration tests")

	// 5. Read Tasks
	readTasks()

	// 6. Save Tasks
	saveTasks()
}
