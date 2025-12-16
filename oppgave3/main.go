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

}

func FinnAllePrimtallOpptil(number int) []int {
	var primtallsFaktorer []int
	
	for i := 1; i <= number; i += 1 {
		if IsAPrime(i) {
			primtallsFaktorer = append(primtallsFaktorer, i)
		}
	}
	
	fmt.Println(primtallsFaktorer)
	return primtallsFaktorer
}

func IsAPrime(i int) bool {
	returnValue := true

	//if i == 2 {
	//	returnValue = true;
	// }
	 
	 for j := 2; j < i; j += 1 {
		//fmt.Printf("i = %v: j = %v \n",i,j)
		if(i % j == 0) {
			//fmt.Println("har ingen rest")
			returnValue = false
			break
		} else {
			returnValue = true
			//fmt.Println("har rest")
		}
	} 
    return returnValue
}