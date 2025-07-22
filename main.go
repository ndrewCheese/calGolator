package main

import "fmt"

func main(){
	var answer string
	var firstNum, secondNum float64
	var decision string

	fmt.Println("calGolator ver.0.0.1")
	for {
		
	
	
	
	fmt.Println("Please choose your action. + - * /")
	
	
	fmt.Scanln(&answer)
	fmt.Println("You chose",answer)
	
	fmt.Print("First number: ")
	if _, err := fmt.Scan(&firstNum); err != nil {
		fmt.Println("Invalid input for first number.")
		return
	}

	fmt.Print("Second number: ")
	if _, err := fmt.Scan(&secondNum); err != nil {
		fmt.Println("Invalid input for second number.")
		return
	}

	var result float64
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
		default:
			fmt.Println("operator err")
			return
	}
	
	fmt.Printf("%.2f %s %.2f = %.2f\n",firstNum, answer, secondNum, result)
	
	fmt.Println("Continue? Y N")
	fmt.Scanln(&decision)
	if decision == "N"{
		break
	}
	}
	
	fmt.Println("Exiting...")		
}
