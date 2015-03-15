package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("This program expects 3 arguments - two numbers and an operator")
	}

	var num1, num2 int64
	_, err := fmt.Sscan(os.Args[1], &num1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Sscan(os.Args[2], &num2)
	if err != nil {
		log.Fatal(err)
	}

	var operator string
	operator = os.Args[3]

	if operator == "+" {
		fmt.Println(num1 + num2)
	} else if operator == "*" {
		fmt.Println(num1 * num2)
	} else {
		log.Fatal("Unknown operator: " + operator)
	}
}
