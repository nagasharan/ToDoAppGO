package main

import "fmt"

func main() {

	var taskOne = "Apply for one company"
	var taskTwo = "Update the resume and analyze which experience to keep or remove"
	var taskThree = "Learn system design basics for 1 hr"
	var taskFour = "Solve 2 leetcode problems"
	fmt.Println("###Hello, Welcome to ToDoList app!!!###")

	fmt.Println("*List of all Todos*")
	fmt.Println(taskOne)
	fmt.Println(taskTwo)
	fmt.Println(taskThree)
	fmt.Println(taskFour)

	fmt.Println()
	fmt.Println("*TUTORIALS*")
	fmt.Println(taskThree)
	fmt.Println(taskFour)

	fmt.Println()
	fmt.Println("*JOB APPLICATION*")
	fmt.Println(taskOne)
	fmt.Println(taskTwo)
}
