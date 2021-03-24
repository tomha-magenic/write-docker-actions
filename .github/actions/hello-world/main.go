package main

import (
	"fmt"
	"os"
)

func main() {

	// Access Inputs as environment vars
	firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
	secondGreeting := os.Getenv("INPUT_SECONDGREETING")
	thirdGreeting := os.Getenv("INPUT_THIRDGREETING")

	// Use those inputs in the action logic

	fmt.Println("Hello " + firstGreeting)
	fmt.Println("Hello " + secondGreeting)

	if thirdGreeting != "" {
		fmt.Println("Hello " + thirdGreeting)
	}
}
