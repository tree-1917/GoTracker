# Task Tracker CLI

**A simple command-line interface (CLI) application to track and manage your tasks.**

This CLI tool helps you organize your to-do list by allowing you to add, update, delete, and mark tasks as in-progress or done. Tasks are stored in a JSON file in the current directory, making it persistent and lightweight.

---

## **Features**

* **Add tasks**: Quickly add new tasks with descriptions.
* **Update tasks**: Edit task descriptions or change their status.
* **Delete tasks**: Remove tasks by ID.
* **Mark tasks**: Mark a task as in-progress or done.
* **List tasks**:

  * List all tasks
  * List tasks by status (`todo`, `in-progress`, `done`)
* **Persistent storage**: All tasks are saved in a JSON file, created automatically if it does not exist.

---

## **Task Properties**

Each task has the following properties:

| Property      | Type     | Description                                   |
| ------------- | -------- | --------------------------------------------- |
| `id`          | int      | Unique identifier for the task                |
| `description` | string   | Short description of the task                 |
| `status`      | string   | Task status: `todo`, `in-progress`, or `done` |
| `createdAt`   | datetime | Date and time the task was created            |
| `updatedAt`   | datetime | Date and time the task was last updated       |

---

## **Installation & Setup**

1. Clone the repository:

```bash
git clone https://github.com/tree-1917/GoTracker.git
cd task-tracker-cli
```

2. Ensure you have a compatible runtime for your chosen language (e.g., Go, Python, Node.js).

3. Run the CLI commands directly in your terminal. The JSON file (`tasks.json`) will be automatically created in the project directory if it does not exist.

---

## **Usage**

### **Adding a Task**

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### **Updating a Task**

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### **Deleting a Task**

```bash
task-cli delete 1
```

### **Marking Task Status**

```bash
task-cli mark-in-progress 2
task-cli mark-done 3
```

### **Listing Tasks**

```bash
# List all tasks
task-cli list

# List tasks by status
task-cli list todo
task-cli list in-progress
task-cli list done
```

---

## **Getting Started**

1. **Create a project directory** and initialize a version control system:

```bash
mkdir task-tracker-cli
cd task-tracker-cli
git init
```

2. **Implement the CLI structure** for handling user inputs.

3. **Implement features step by step**:

   * Add task → List tasks → Update task → Mark task → Delete task

4. **Test each feature** thoroughly before moving to the next.

5. **Check the JSON file** to ensure tasks are stored correctly.

---

## **Contributing**

Contributions are welcome! You can:

* Submit bug reports
* Suggest improvements
* Add features or enhance existing functionality

Please make sure all contributions maintain consistency with the project structure and CLI usage.

---

## **License**

This project is licensed under the MIT License.

---
