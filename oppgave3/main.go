package main

import "fmt"


// Solution for Project Euler problem 3
func main() {
	const numberToFactor = 600851475143

	fmt.Println("Hello, this is my anwser. The number to factor is ", numberToFactor)

	fmt.Println("Finding prime factors...")
	
	listOfFactors := GreatestDivisorCalculator(numberToFactor)

	fmt.Println("**************")
	fmt.Println("List of primes factors", listOfFactors)
	sum := MultiplyFactors(listOfFactors)
	fmt.Println("Product of prims", sum, " should be equal to ", numberToFactor)

	fmt.Println("Larges factor", listOfFactors[len(listOfFactors)-1])

}

// Greates Divisor Calculator. Algritme to find prime factors
// returns a list of prime factors for the number 600851475143
func GreatestDivisorCalculator(numberToFactor int) []int {

	listOfPrimes := FindListOfPrimNumbersUpTo(numberToFactor)
	collectedPrimeFactors := []int{} 
	tall := numberToFactor

	for len(listOfPrimes) > 0 {
		
		primeNumber := listOfPrimes[0]

		if (tall % primeNumber == 0) {
			tall = tall / primeNumber

			collectedPrimeFactors = append(collectedPrimeFactors, primeNumber)
			
		} else {
			listOfPrimes = listOfPrimes[1:]
		}
	}	

	return collectedPrimeFactors
}

// Multiplies all numbers in a list and returns the product
func MultiplyFactors(numbers []int) int{
	product := 1

	for i := 0; i < len(numbers); i += 1 {
		product = product * numbers[i]
	}

	return product

}


// Finds all prime numbers up to a given number
func FindListOfPrimNumbersUpTo(number int) []int {
	var primeFactors []int

	for counter := 1; counter*counter <= number; counter += 1 {		
		if IsAPrime(counter) {
			primeFactors = append(primeFactors, counter)
		}
	}
		
	return primeFactors

}

// Checks if a number is a prime number
func IsAPrime(i int) bool {
	returnValue := true

	if (i == 1) { return false}
		 
	for j := 2; j*j < i; j += 1 {		
		if(i % j == 0) {			
			// has a divisor
			returnValue = false
			break
		} else {
			// does not have a divisor 
			returnValue = true			
		}
	} 
    return returnValue
}