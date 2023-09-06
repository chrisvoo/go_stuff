package main

import (
	"fmt"
)

type Subtask struct {
	Param1 string
	Param2 string
	Status string
}

func tasks() {
	allTasks := []Subtask{{Status: "completed"}, {Status: "incompleted"}}
	// the _ stands for the index, in case of maps, for their keys
	for idx, x := range allTasks {
		if x.Status != "completed" {
			fmt.Printf("Task %d is still incomplete\n", idx)
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	fmt.Println()
}
