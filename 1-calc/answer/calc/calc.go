package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("This program expects 3 arguments - two numbers and an operator")
	}

	num1, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	num2, err := strconv.ParseInt(os.Args[2], 10, 64)
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
