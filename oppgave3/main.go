package main

import "fmt"


// lag en funksjon som sorterer en liste
func main() {
	
	fmt.Println("Hello ")
	
	gdc()
}

func gdc(){

	const number = 600851475143
	
	primes := FindPrimFactorsUpTo(number)
	primeFactors := []int{} 
	tall := number

	for len(primes) > 0 {
		
		primeNumber := primes[0]

		if (tall % primeNumber == 0) {
			tall = tall / primeNumber

			primeFactors = append(primeFactors, primeNumber)
			
		} else {
			primes = primes[1:]
		}
	}

	fmt.Println("**************")
	fmt.Println("List of primes factors", primeFactors)
	sum := MultiplyFactors(primeFactors)
	fmt.Println("Product of prims", sum )

	fmt.Println("Larges factor", primeFactors[len(primeFactors)-1])
}

func MultiplyFactors(numbers []int) int{
	product := 1

	for i := 0; i < len(numbers); i += 1 {
		product = product * numbers[i]
	}

	return product

}



func FindPrimFactorsUpTo(number int) []int {
	var primeFactors []int

	for i := 1; i*i <= number; i += 1 {
		fmt.Println("Numbers tested",i," . Primenumbers", len(primeFactors))
		if IsAPrime(i) {
			primeFactors = append(primeFactors, i)
		}
	}
	
	fmt.Println(primeFactors)
	return primeFactors

}

func IsAPrime(i int) bool {
	returnValue := true

	if (i == 1) { return false}
		 
	for j := 2; j*j < i; j += 1 {		
		if(i % j == 0) {
			// har ingen rest, altså ikke et primtall
			returnValue = false
			break
		} else {
			// har en rest 
			returnValue = true			
		}
	} 
    return returnValue
}