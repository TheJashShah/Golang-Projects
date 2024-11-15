package main

import (
	"fmt"
	"math/rand"
)

type task struct {
	content     string
	id          int
	priority    int
	is_complete int
}

func generateRandom() int {

	id := rand.Intn(10000) + 1

	return id
}

func CreateTask(content string, priority int) task {

	T := task{
		content:     content,
		id:          generateRandom(),
		priority:    priority,
		is_complete: 0,
	}

	return T
}

func (Task *task) returnString() string {

	isComplete := ""

	if Task.is_complete == 1 {
		isComplete = "True"
	} else {
		isComplete = "False"
	}

	str := fmt.Sprintf("| %-25v | %-5v | %-10v | %-6v | \n", Task.content, Task.id, Task.priority, isComplete)

	return str

}

func (Task *task) UpdateTask(new string) {
	Task.content = new
}

func (Task *task) UpdatePriority(new int) {
	Task.priority = new
}
func (Task *task) updateStatus(new int) {
	Task.is_complete = new
}
