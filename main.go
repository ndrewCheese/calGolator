package main

import (
	"fmt"
	"math"
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
	
	fmt.Println("calGolator v0.1")
	for {
	
	fmt.Println("Please choose your action. + - * / ^ sqrt")
	fmt.Scan(&op)

	if !validOps[op] { // if op isn't in validOps
		fmt.Println("Invalid operator:",op)
		return
	}

	fmt.Print("First number: ")
	if _, err := fmt.Scan(&firstNum); err != nil {
		fmt.Println("Invalid input for first number")
		return
	}

	if op == "sqrt"{ // we only need one number for sqrt
		fmt.Printf("√%.2f = %.2f\n",firstNum,math.Sqrt(firstNum))
		fmt.Println("Continue? y n")
		fmt.Scan(&decision)
		if decision == "n"{
			break
		}

	} else {
		fmt.Print("Second number: ")
			if _, err := fmt.Scan(&secondNum); err != nil {
				fmt.Println("Invalid input for second number")
				return
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
			
			fmt.Printf("%.2f %s %.2f = %.2f\n",firstNum, op, secondNum, result)
			
			fmt.Println("Continue? y n")
			fmt.Scan(&decision)
			if decision == "n"{
				break
			}
	}
	}
	fmt.Println("Exiting...")		
}
