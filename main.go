package main

import (
	"fmt"
	"math"
)

func main(){
	var answer string
	var firstNum, secondNum float64
	var result float64
	var decision string

	fmt.Println("calGolator ver.0.0.1")
	for {
	
	fmt.Println("Please choose your action. + - * / ^ sqrt")
	
	fmt.Scan(&answer)
	fmt.Println("You chose",answer)

	
	fmt.Print("First number: ")
	if _, err := fmt.Scan(&firstNum); err != nil {
		fmt.Println("Invalid input for first number.")
		return
	}

	if answer == "sqrt"{ // We only need one number for sqrt
		fmt.Printf("√%.2f = %.2f\n",firstNum,math.Sqrt(firstNum))
		fmt.Println("Continue? y n")
		fmt.Scan(&decision)
		if decision == "n"{
			break
		}
	} else {
		fmt.Print("Second number: ")
			if _, err := fmt.Scan(&secondNum); err != nil {
				fmt.Println("Invalid input for second number.")
				return
			}
		
			switch answer {
				case "+":
					result = (firstNum + secondNum)
				case "-":
					result = (firstNum - secondNum)
				case "*":
					result = (firstNum * secondNum)
				case "/":
					if secondNum == 0 {
						fmt.Println("division by 0 err")
						return
					}
					result = (firstNum / secondNum)
				case "^":
					result = math.Pow(firstNum, secondNum)
				default:
					fmt.Println("operator err")
					return
			}
			
			fmt.Printf("%.2f %s %.2f = %.2f\n",firstNum, answer, secondNum, result)
			
			fmt.Println("Continue? y n")
			fmt.Scan(&decision)
			if decision == "n"{
				break
			}
	}
	}
	fmt.Println("Exiting...")		
}
