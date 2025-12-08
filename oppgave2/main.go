package main

import "fmt"


// lag en funksjon som sorterer en liste
func main() {
	const MAX_NUMBERS_TERMS_OF_FABONACCI_NUMBERS = 4000000
	a, b := 0, 1

	sumEventFabonacciNumber := 0; 

	for a < MAX_NUMBERS_TERMS_OF_FABONACCI_NUMBERS {

		if a % 2 == 0 {
			fmt.Println(a)
			sumEventFabonacciNumber = sumEventFabonacciNumber + a
		}

		a,b = b,a+b
	}

	fmt.Println("Sum Partall Fabonacci tall opp til 4 millioner :%d \n", sumEventFabonacciNumber)
}