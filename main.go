package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Complete    bool
	Priority    string
	DueDate     time.Time
	Tags        []string
	Category    string
	Description string
}

var tasks []Task
const dataFile = "tasks.json"

func main() {
	loadTasksFromFile()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nCommands:")
		fmt.Println("1. ADD - Add a new task")
		fmt.Println("2. VIEW - View all tasks")
		fmt.Println("3. EDIT - Edit a task")
		fmt.Println("4. DELETE - Delete a task")
		fmt.Println("5. FILTER - Filter tasks")
		fmt.Println("6. OVERDUE - Check overdue tasks")
		fmt.Println("7. SYNC - Sync tasks with server")
		fmt.Println("8. SAVE - Save tasks")
		fmt.Println("9. EXIT - Exit the program")
		fmt.Print("Enter command number: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "1":
			addTask(reader)
		case "2":
			viewTasks()
		case "3":
			editTask(reader)
		case "4":
			deleteTask(reader)
		case "5":
			filterTasks(reader)
		case "6":
			checkOverdueTasks()
		case "7":
			syncTasksWithServer()
		case "8":
			saveTasksToFile()
		case "9":
			fmt.Println("Exiting...")
			saveTasksToFile()
			return
		default:
			fmt.Println("Invalid command number. Please try again.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Println("\nAdding a new task")
	fmt.Print("Task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Priority: ")
	priority, _ := reader.ReadString('\n')
	priority = strings.TrimSpace(priority)

	fmt.Print("Due date (format yyyy-mm-dd): ")
	dueDateStr, _ := reader.ReadString('\n')
	dueDateStr = strings.TrimSpace(dueDateStr)
	dueDate, _ := time.Parse("2006-01-02", dueDateStr)

	fmt.Print("Category: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Print("Task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Tags (comma separated): ")
	tagsStr, _ := reader.ReadString('\n')
	tagsStr = strings.TrimSpace(tagsStr)
	tags := strings.Split(tagsStr, ",")

	task := Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Priority:    priority,
		DueDate:     dueDate,
		Complete:    false,
		Tags:        tags,
		Category:    category,
		Description: description,
	}

	tasks = append(tasks, task)
	fmt.Println("Task added successfully.")
}

func viewTasks() {
	fmt.Println("\nTasks:")
	for _, task := range tasks {
		status := "Incomplete"
		if task.Complete {
			status = "Complete"
		}
		fmt.Printf("ID: %d\nTitle: %s\nStatus: %s\nPriority: %s\nDue Date: %s\nCategory: %s\nDescription: %s\nTags: %s\n\n",
			task.ID, task.Title, status, task.Priority, task.DueDate.Format("2006-01-02"), task.Category, task.Description, strings.Join(task.Tags, ", "))
	}
}

func editTask(reader *bufio.Reader) {
	fmt.Print("Enter the ID of the task to edit: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	for i, task := range tasks {
		if task.ID == id {
			fmt.Println("\nEditing task")
			
			fmt.Print("New title (leave empty to keep current): ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title != "" {
				tasks[i].Title = title
			}

			fmt.Print("New priority (leave empty to keep current): ")
			priority, _ := reader.ReadString('\n')
			priority = strings.TrimSpace(priority)
			if priority != "" {
				tasks[i].Priority = priority
			}

			fmt.Print("New due date (format yyyy-mm-dd) (leave empty to keep current): ")
			dueDateStr, _ := reader.ReadString('\n')
			dueDateStr = strings.TrimSpace(dueDateStr)
			if dueDateStr != "" {
				dueDate, _ := time.Parse("2006-01-02", dueDateStr)
				tasks[i].DueDate = dueDate
			}

			fmt.Print("New category (leave empty to keep current): ")
			category, _ := reader.ReadString('\n')
			category = strings.TrimSpace(category)
			if category != "" {
				tasks[i].Category = category
			}

			fmt.Print("New description (leave empty to keep current): ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			if description != "" {
				tasks[i].Description = description
			}

			fmt.Print("New tags (comma separated) (leave empty to keep current): ")
			tagsStr, _ := reader.ReadString('\n')
			tagsStr = strings.TrimSpace(tagsStr)
			if tagsStr != "" {
				tasks[i].Tags = strings.Split(tagsStr, ",")
			}

			fmt.Println("Task edited successfully.")
			return
		}
	}
	fmt.Println("Task with this ID not found.")
}

func deleteTask(reader *bufio.Reader) {
	fmt.Print("Enter the ID of the task to delete: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task with this ID not found.")
}

func checkOverdueTasks() {
	now := time.Now()
	fmt.Println("\nOverdue tasks:")
	for _, task := range tasks {
		if !task.Complete && task.DueDate.Before(now) {
			fmt.Printf("ID: %d\nTitle: %s\nDue Date: %s\n\n",
				task.ID, task.Title, task.DueDate.Format("2006-01-02"))
		}
	}
}

func filterTasks(reader *bufio.Reader) {
	fmt.Print("Enter filter (priority, due date, or keyword): ")
	filter, _ := reader.ReadString('\n')
	filter = strings.TrimSpace(filter)

	fmt.Println("\nFilter results:")
	for _, task := range tasks {
		if strings.Contains(task.Title, filter) ||
			strings.EqualFold(task.Priority, filter) ||
			task.DueDate.Format("2006-01-02") == filter {
			status := "Incomplete"
			if task.Complete {
				status = "Complete"
			}
			fmt.Printf("ID: %d\nTitle: %s\nStatus: %s\nPriority: %s\nDue Date: %s\nCategory: %s\nDescription: %s\nTags: %s\n\n",
				task.ID, task.Title, status, task.Priority, task.DueDate.Format("2006-01-02"), task.Category, task.Description, strings.Join(task.Tags, ", "))
		}
	}
}

func syncTasksWithServer() {
	// Replace URL with your server URL
	url := "http://localhost:8080/sync"
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error converting data to JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error syncing with server:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Tasks successfully synced with the server.")
	} else {
		fmt.Println("Error syncing with server:", resp.Status)
	}
}

func saveTasksToFile() {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Tasks successfully saved.")
	}
}

func loadTasksFromFile() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, create a new one
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
