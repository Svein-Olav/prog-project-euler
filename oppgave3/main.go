package main

import "fmt"


// lag en funksjon som sorterer en liste
func main() {
	
	fmt.Println("Hello, World!")
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
	returnValue := false

	if i == 2 {
		returnValue = true;
	 }

	 /*
	 if(i % 2 == 0 && i > 2) {
		returnValue = false
	 }
	 */
	 for j := 3; j <= i; j += 2 {
		fmt.Printf("i = %v: j = %v \n",i,j)
		if(i % j == 0) {
			fmt.Printf("har ingen rest")
			returnValue = false
			break
		} else {
			returnValue = true
			fmt.Printf("har rest")
		}
	} 
    return returnValue
}