package main

import (
	"fmt"
)

func main() {
	for {
		var num1, num2 float64
		var action string

		fmt.Println("Enter first num:")
		_, err1 := fmt.Scanln(&num1)
		if err1 != nil {
			fmt.Println("Error: invalid number!")
			var dump string
			fmt.Scanln(&dump) 
			continue
		}

		fmt.Println("Enter second num:")
		_, err2 := fmt.Scanln(&num2)
		if err2 != nil {
			fmt.Println("Error: invalid number!")
			var dump string
			fmt.Scanln(&dump)
			continue
		}

		fmt.Println("select action: mul, add, sub, div, nil, help:")
		fmt.Scanln(&action)

		if action == "mul" {
			fmt.Println(num1 * num2)

		} else if action == "add" {
			fmt.Println(num1 + num2)

		} else if action == "sub" {
			fmt.Println(num1 - num2)

		} else if action == "div" {
			if num2 == 0 {
				fmt.Println("Error: cannot divide by zero!")
			} else {
				fmt.Println(num1 / num2)
			}

		} else if action == "nil" {
			fmt.Println("Stopping program...")
			return

		} else if action == "help"{
		fmt.Println("enter first number, press enter then enter the second number ")
		fmt.Println("then press enter a list of mathematical operation will appear")
		fmt.Println("choose by typing the word not the number on the new line press enter to see the anwer")
		fmt.Println("if you want to quit type nil and press enter")
	}else{
			fmt.Println("Unknown action!")
		}
	}
}
