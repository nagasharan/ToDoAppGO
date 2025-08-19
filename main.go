package main

import (
	"fmt"
	"net/http"
)

var taskOne = "Apply for one company"
var taskTwo = "Update the resume and analyze which experience to keep or remove"
var taskThree = "Learn system design basics for 1 hr"
var taskFour = "Solve 2 leetcode problems"

// Slice is a variable size but array is fixed size
var listOfItems = []string{taskOne, taskTwo, taskThree, taskFour}

func main() {

	//var taskOne = "Apply for one company"
	//var taskTwo = "Update the resume and analyze which experience to keep or remove"
	//var taskThree = "Learn system design basics for 1 hr"
	//var taskFour = "Solve 2 leetcode problems"

	//Slice is a variable size but array is fixed size
	//var listOfItems = []string{taskOne, taskTwo, taskThree, taskFour}
	fmt.Println("###Hello, Welcome to ToDoList app!!!###")
	fmt.Printf("test")
	//printAll(listOfItems)
	//fmt.Println()
	//listOfItems = addTask(listOfItems, "Go for run")
	//fmt.Println()
	//printAll(listOfItems)

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/tasks", showTasks)
	http.ListenAndServe("localhost:8080", nil)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {

	for _, task := range listOfItems {
		fmt.Fprintln(writer, task)
	}

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello! How are you doing."
	fmt.Fprintln(writer, greeting)
}

func printAll(listOfItems []string) {
	fmt.Println("*List of all Todos*")

	for index, item := range listOfItems {
		//fmt.Println(index+1, ". "+item)
		fmt.Printf("%d. %s\n", index+1, item)
	}
}

func addTask(listOfItems []string, newTask string) []string {
	listOfItems = append(listOfItems, newTask)
	return listOfItems
}
