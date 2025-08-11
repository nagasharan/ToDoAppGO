package main

import "fmt"

func main() {

	var taskOne = "Apply for one company"
	var taskTwo = "Update the resume and analyze which experience to keep or remove"
	var taskThree = "Learn system design basics for 1 hr"
	var taskFour = "Solve 2 leetcode problems"

	//Slice is a variable size but array is fixed size
	var listOfItems = []string{taskOne, taskTwo, taskThree, taskFour}
	fmt.Println("###Hello, Welcome to ToDoList app!!!###")
	printAll(listOfItems)
}

func printAll(listOfItems []string) {
	fmt.Println("*List of all Todos*")

	for index, item := range listOfItems {
		//fmt.Println(index+1, ". "+item)
		fmt.Printf("%d. %s\n", index+1, item)
	}
}
