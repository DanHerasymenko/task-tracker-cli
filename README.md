
# Simple CLI Task Tracker


A simple command-line interface (CLI) application to manage tasks. This task tracker allows you to add, update, delete, and manage tasks by their statuses.



## Getting Started

To build and run the project:

```golang
go build
```

## Commands and Usage
Adding a New Task
```golang
task-tracker-cli add "Buy groceries"
```
Updating and Deleting Tasks
```golang
task-tracker-cli update 1 "Buy groceries and cook dinner"
task-tracker-cli delete 1
```
Marking a Task as In Progress or Done
```golang
task-tracker-cli mark-in-progress 1
task-tracker-cli mark-done 1
```
Listing All Tasks
```golang
task-tracker-cli list
```
Listing Tasks by Status
```golang
task-tracker-cli list done
task-tracker-cli list todo
task-tracker-cli list in-progress
```

## Project Inspiration

This project idea was inspired by [Roadmap.sh Task Tracker Project.](https://roadmap.sh/projects/task-tracker)

