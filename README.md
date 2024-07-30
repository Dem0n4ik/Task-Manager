# Task Manager

A command-line tool written in Go for managing tasks. This tool allows you to add, view, edit, delete, filter, and check overdue tasks. Additionally, it provides functionality to sync tasks with a server and save them to a local file.

## Features

- **Add a New Task**: Create a new task with title, priority, due date, category, description, and tags.
- **View All Tasks**: Display a list of all tasks with their details.
- **Edit a Task**: Modify existing tasks' details.
- **Delete a Task**: Remove a task from the list.
- **Filter Tasks**: Filter tasks by priority, due date, or keyword.
- **Check Overdue Tasks**: View tasks that are overdue.
- **Sync with Server**: Synchronize tasks with a remote server.
- **Save Tasks**: Save tasks to a local JSON file.
- **Load Tasks**: Load tasks from a local JSON file.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Dem0n4ik/task-manager.git
   cd task-manager
   ```

2. Initialize Go modules:
   ```bash
   go mod init task-manager
   ```

3. Build the application:
   ```bash
   go build -o task_manager
   ```

## Usage

1. Run the application:
   ```bash
   ./task_manager
   ```

2. You will see a menu with available commands:
   ```
   Commands:
   1. ADD - Add a new task
   2. VIEW - View all tasks
   3. EDIT - Edit a task
   4. DELETE - Delete a task
   5. FILTER - Filter tasks
   6. OVERDUE - Check overdue tasks
   7. SYNC - Sync tasks with server
   8. SAVE - Save tasks
   9. EXIT - Exit the program
   ```

3. Enter the number corresponding to the desired command and follow the prompts.

### Example

```
Commands:
1. ADD - Add a new task
2. VIEW - View all tasks
3. EDIT - Edit a task
4. DELETE - Delete a task
5. FILTER - Filter tasks
6. OVERDUE - Check overdue tasks
7. SYNC - Sync tasks with server
8. SAVE - Save tasks
9. EXIT - Exit the program
Enter command number: 1

Adding a new task
Task title: Buy groceries
Priority: High
Due date (format yyyy-mm-dd): 2024-08-01
Category: Personal
Task description: Buy milk, eggs, and bread
Tags (comma separated): shopping,urgent
Task added successfully.
```

## Dependencies

- **Go 1.16+**: Make sure Go is installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

- **net/http**: Used for synchronizing tasks with a remote server.

## Configuration

- **Server URL**: In the `syncTasksWithServer` function, replace the `url` variable with your server's URL.

## File Storage

- **Data File**: The tasks are stored in a file named `tasks.json`. This file will be created in the same directory as the executable.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
