package main

import "fmt"


// lag en funksjon som sorterer en liste
func main() {
	
	fmt.Println("Hello, World!")

	// prime numbers : 2, 3, 5, 7, 11, 13, 17, 19

	for i :=1; i <= 20; i += 1 {
		isPrime := IsAPrime(i)
		if isPrime {
			fmt.Println("is prime",i)
		} else {
			fmt.Println("is not prime",i)
		}
    }

	FindPrimFactorsUpTo(500)

}

func gdc(){
	primeFactors :=  FindPrimFactorsUpTo(500)
	tall := 6000

	for len(primeFactors) > 0 {
		primeNumber := primeFactors[0]

		if (tall % primeNumber == 0) {
			
		}


	}



}



func FindPrimFactorsUpTo(number int) []int {
	var primeFactors []int

	for i := 1; i <= number; i += 1 {
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
		 
	for j := 2; j < i; j += 1 {		
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