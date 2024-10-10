package utils

import (
	"fmt"
	"log"
)

func TakeInput(methodType string) string {

	var input string
	if methodType == "accept/reject" {
		fmt.Print("Enter 1/0 to accept/reject: ")
	}
	if methodType == "complete" {
		fmt.Print("Enter 1 to complete: ")
	}
	n, err := fmt.Scanf("%v", &input)
	if err != nil || n != 1 {
		log.Fatalf("scanf error %v", err)
	}
	return input
}
