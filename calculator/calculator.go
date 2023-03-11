package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Run into an error")
			main()
		}
	}()

	type operation struct {
		StartNumber  float64
		FunctionName string
		NextNumber   float64
		Result       float64
	}

	functions := map[string]func(float64, float64) float64{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":   divide,
		"power":    power,
		"sqrt":     sqrt,
	}

	var currentNumber float64
	for {
		var input string
		fmt.Print("Enter the start number : ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue // continue the loop to re-prompt the input
		}

		currentNumber, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue // continue the loop to re-prompt the input
		}
		break
	}

	history := []operation{}

	running := true //jika variabel running false, maka program berhenti saat pengguna memasukkan exit.
	for running {
		var functionName string
		fmt.Print("What operator do you want to use\n(type 'done' to finish, type 'exit' to quit): ")
		fmt.Scanln(&functionName)

		switch functionName {
		case "done":
			running = false
		case "exit":
			fmt.Println("Exiting...")
			running = false
			return
		case "sqrt":
			result := functions[functionName](currentNumber, 0)
			history = append(history, operation{
				StartNumber:  currentNumber,
				FunctionName: functionName,
				NextNumber:   0,
				Result:       result,
			})
			currentNumber = result
		case "power":
			var nextNumber float64
			for {
				var inputNumber string
				fmt.Print("Enter the power : ")
				fmt.Scanln(&inputNumber)

				number, err := strconv.ParseFloat(inputNumber, 64)
				if err != nil {
					fmt.Println("Invalid input! Please enter a number.")
					continue
				}

				nextNumber = number
				break
			}

			result := functions[functionName](currentNumber, nextNumber)
			history = append(history, operation{
				StartNumber:  currentNumber,
				FunctionName: functionName,
				NextNumber:   nextNumber,
				Result:       result,
			})
			currentNumber = result
		default:
			var nextNumber float64
			for {
				var inputNumber string
				fmt.Print("Enter the next number : ")
				fmt.Scanln(&inputNumber)

				number, err := strconv.ParseFloat(inputNumber, 64)
				if err != nil {
					fmt.Println("Invalid input! Please enter a number.")
					continue
				}

				nextNumber = number
				break
			}

			result := functions[functionName](currentNumber, nextNumber)
			history = append(history, operation{
				StartNumber:  currentNumber,
				FunctionName: functionName,
				NextNumber:   nextNumber,
				Result:       result,
			})
			currentNumber = result
		}
	}

	fmt.Println("History:")
	if len(history) > 0 {
		fmt.Printf("start number = %v\n", history[0].StartNumber)
		for _, op := range history {
			fmt.Printf("%v %v = %v\n", op.FunctionName, op.NextNumber, op.Result)
		}
	}
	fmt.Println("Result:")
	fmt.Println(currentNumber)
}

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if x == 0 || y == 0{
		panic("divide by zero, error!")
	}
	return x / y
}

func power(x, y float64) float64 {
	return math.Pow(x, y)
}

func sqrt(x, y float64) float64 {
	return math.Sqrt(x)
}