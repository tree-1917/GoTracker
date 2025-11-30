# 📝 Task Tracker CLI

**A fun and simple command-line app to manage your tasks and to-do list!**

Keep track of what you need to do, what you are currently working on, and what you have completed — all from your terminal.

---

## **✨ Features**

* ➕ **Add tasks**: Quickly add new tasks with descriptions.
* ✏️ **Update tasks**: Edit task descriptions or status.
* ❌ **Delete tasks**: Remove tasks by ID.
* ✅ **Mark tasks**: Mark as `in-progress` or `done`.
* 📋 **List tasks**:

  * List **all tasks**
  * List tasks by **status** (`todo`, `in-progress`, `done`)
* 💾 **Persistent storage**: Tasks are saved in a JSON file automatically.

---

## **📦 Task Properties**

| Property      | Type     | Description                                   |
| ------------- | -------- | --------------------------------------------- |
| `id`          | int      | Unique identifier for the task                |
| `description` | string   | Short description of the task                 |
| `status`      | string   | Task status: `todo`, `in-progress`, or `done` |
| `createdAt`   | datetime | When the task was created                     |
| `updatedAt`   | datetime | Last time the task was updated                |

---

## **⚡ Installation & Setup**

1. Clone the repository:

```bash
git clone https://github.com/tree-1917/GoTracker.git
cd task-tracker-cli
```

2. Ensure your runtime is installed (Go, Python, Node.js, etc.)

3. Run the CLI commands in your terminal. The JSON file (`tasks.json`) will be created automatically.

---

## **💻 Usage Examples**

### Add a Task

```bash
task-cli add "Buy groceries" 
# ➕ Task added successfully (ID: 1)
```

### Update a Task

```bash
task-cli update 1 "Buy groceries and cook dinner" 
# ✏️ Task updated successfully
```

### Delete a Task

```bash
task-cli delete 1
# ❌ Task deleted successfully
```

### Mark Task Status

```bash
task-cli mark-in-progress 2
# 🔄 Task marked as in-progress

task-cli mark-done 3
# ✅ Task marked as done
```

### List Tasks

```bash
task-cli list           # List all tasks
task-cli list todo      # List tasks that are todo
task-cli list in-progress  # List tasks in progress
task-cli list done      # List completed tasks
```

---

## **🚀 Future Features (TUI with Bubble Tea)**

We plan to enhance the CLI experience using [**Bubble Tea**](https://github.com/charmbracelet/bubbletea), a **stateful and fun Go framework** for terminal apps.

### What’s coming:

* 🖥️ **Interactive terminal UI**: Navigate tasks with arrow keys.
* 🎨 **Colorful interface**: Highlight tasks by status.
* 🔄 **Real-time updates**: Toggle status or edit descriptions without typing commands.
* 📊 **Progress indicators**: Track tasks visually using bars or spinners.

> Bubble Tea will turn the Task Tracker CLI into a **modern, fun, and fully interactive terminal application**.

---

## **🛠️ Getting Started**

1. Create a project directory and initialize Git:

```bash
mkdir task-tracker-cli
cd task-tracker-cli
git init
```

2. Implement CLI commands: add, update, delete, mark, list.

3. Test each feature and check `tasks.json` to ensure tasks are stored correctly.

4. Later, integrate Bubble Tea for a TUI experience.

---

## **🤝 Contributing**

Contributions are welcome! You can:

* Report bugs 🐞
* Suggest new features 💡
* Improve code or add documentation 📝

Please make sure all contributions maintain CLI functionality and code consistency.

---

## **📄 License**

This project is licensed under the **MIT License**.

---

## **🎉 Happy Coding!**

Track your tasks, stay organized, and enjoy a modern terminal experience!

---
