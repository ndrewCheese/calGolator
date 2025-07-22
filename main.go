package main

import (
	"fmt"
	"math"
	"strings"
)

func main(){
	var op string
	var firstNum, secondNum float64
	var result float64
	var decision string
	validOps := map[string]bool{
		"+":true,
		"-":true,
		"*":true,
		"/":true,
		"^":true,
		"sqrt":true,
	}

	fmt.Println("calGolator v0.2")
	for {
	for {
		fmt.Println("Please choose your action. + - * / ^ sqrt")
		fmt.Scan(&op)
		op = strings.ToLower(strings.TrimSpace(op))
		if !validOps[op] { // if op isn't in validOps
			fmt.Println("Invalid operator:",op)
		} else {
			break
		}
	}

	for {
		fmt.Print("First number: ")
		if _, err := fmt.Scan(&firstNum); err != nil {
			fmt.Println("Invalid input for first number. Try again")
		} else {
			break
		}
	}
	
	if op == "sqrt"{ // we only need one number for sqrt
		fmt.Printf("√%.2f = %.0f\n",firstNum,math.Sqrt(firstNum))
		fmt.Println("Continue? y n")
		fmt.Scan(&decision)
		if decision == "n"{
			break
		}

	} else {
	for {
			fmt.Print("Second number: ")
			if _, err := fmt.Scan(&secondNum); err != nil {
				fmt.Println("Invalid input for second number. Try again")
			} else {
				break
			}
		}

		switch op {
			case "+":
				result = (firstNum + secondNum)
			case "-":
				result = (firstNum - secondNum)
			case "*":
				result = (firstNum * secondNum)
			case "/":
				if secondNum == 0 {
					fmt.Println("Can't divide by zero")
					return
				}
				result = (firstNum / secondNum)
			case "^":
				result = math.Pow(firstNum, secondNum)
			default:
				return
		}

		if result == math.Trunc(result){
			fmt.Printf("%.2f %s %.2f = %.0f\n",firstNum, op, secondNum, result)
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f\n",firstNum, op, secondNum, result)
		}

		fmt.Println("Continue? y n")
		fmt.Scan(&decision)
		decision = strings.ToLower(strings.TrimSpace(decision))
		if decision == "n"{
			break
		}
	}
	}
	fmt.Println("Exiting...")		
}
