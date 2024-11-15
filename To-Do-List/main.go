package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(prompt string) string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	i, _ := reader.ReadString('\n')

	input := strings.TrimSpace(i)

	return input
}

func DisplayList(list []task) string {

	task_ := "Task"
	id_ := "ID"
	priority_ := "Priority"
	Status := "Status"

	str := fmt.Sprintf("| %-25v | %-5v | %-10v | %-5v | \n", task_, id_, priority_, Status)

	for _, Task := range list {
		str += Task.returnString()
	}

	return str
}

func SortList(list []task) {

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].priority > list[j].priority {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func WriteToFile(list string) {

	data := []byte(list)

	err := os.WriteFile("To-Do-List.txt", data, 0644)

	if err != nil {
		panic(err)
	}
}

func ReadFromFile(task_list *[]task) {

	file, err := os.Open("To-Do-List.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		for i := 0; i < len(fields); i++ {
			if fields[i] == "|" {
				fields = append(fields[:i], fields[i+1:]...)
			}
		}

		content := ""

		for j := 0; j < len(fields)-3; j++ {
			content += fields[j] + " "
		}

		id, err := strconv.Atoi(fields[len(fields)-3])

		if err != nil {
			fmt.Println(err)
			continue
		}

		priority, err := strconv.Atoi(fields[len(fields)-2])

		if err != nil {
			fmt.Println(err)
			continue
		}

		status := 0

		if fields[len(fields)-1] == "True" {
			status = 1
		}

		*task_list = append(*task_list, task{
			content:     content,
			id:          id,
			priority:    priority,
			is_complete: status,
		})
	}
}

func main() {

	task_list := []task{}

	ReadFromFile(&task_list)

	for {

		fmt.Println(" ")
		fmt.Println("1. Enter a new Task.")
		fmt.Println("2. Update an existing Task.")
		fmt.Println("3. Watch the List.")
		fmt.Println("4. Delete a Task.")
		fmt.Println("5. Save the List.")
		fmt.Println("6. Exit.")
		fmt.Println(" ")

		c := input("Enter you choice: ")

		choice, err := strconv.Atoi(c)

		if err != nil {
			fmt.Println("Enter a valid number.")
			continue
		}

		switch choice {

		case 1:

			content := input("Enter the Task: ")

			p := input("Enter the priority[should be 1 and 5]: ")

			priority, err := strconv.Atoi(p)

			if err != nil {
				fmt.Println("Enter an integer.")
			}

			Task := CreateTask(content, priority)

			task_list = append(task_list, Task)

			fmt.Println("A Task is created.")

		case 2:

			i := input("Enter the ID of the Task: ")

			flag := false

			id, err := strconv.Atoi(i)

			if err != nil {
				fmt.Println("Enter a number.")
			}

			for j := 0; j < len(task_list); j++ {
				if task_list[j].id == id {
					flag = true

					choice := input("Update [C]ontent/[P]riority/[S]tatus: ")

					switch choice {

					case "C":
						new_content := input("Enter the new Task: ")
						task_list[j].UpdateTask(new_content)

						fmt.Println("The Task is updated.")

					case "P":
						new_p := input("Enter the new Prioriy: ")
						new_priority, err := strconv.Atoi(new_p)

						if err != nil {
							fmt.Println("Enter a number.")
						}
						task_list[j].UpdatePriority(new_priority)

						fmt.Println("The Task is updated.")

					case "S":
						new_s := input("Enter the new Status[only 1 as input will count as completion]: ")
						new_status, err := strconv.Atoi(new_s)

						if err != nil {
							fmt.Println("Enter a number.")
						}

						task_list[j].updateStatus(new_status)

						fmt.Println("The Task is updated.")

					default:
						fmt.Println("Enter one of the three options.")
					}
				}
			}

			if !flag {
				fmt.Println("No such Task in the list.")
			}

		case 3:
			SortList(task_list)

			fmt.Println(DisplayList(task_list))

		case 4:

			flag := false

			i := input("Enter the ID of the task: ")
			id, err := strconv.Atoi(i)

			if err != nil {
				fmt.Println("Enter a number.")
			}

			for j := 0; j < len(task_list); j++ {
				if task_list[j].id == id {
					flag = true

					task_list[j] = task_list[len(task_list)-1]

					task_list = task_list[:len(task_list)-1]

					fmt.Println("Task is Deleted.")
				}
			}

			if !flag {
				fmt.Println("No such Task in the list.")
			}

		case 5:
			SortList(task_list)

			WriteToFile(DisplayList(task_list))
			fmt.Println("The List is Saved.")

		case 6:
			SortList(task_list)

			WriteToFile(DisplayList(task_list))

			fmt.Println("The List is Saved.")
			fmt.Println("Exiting program...")

			return

		default:
			fmt.Println("Enter a number between 1 and 6.")

		}
	}
}
