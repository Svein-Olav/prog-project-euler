package main

import "fmt"

func main() {
    slice_steg_3 := AlleTallMultiplumAv(3, 1000)
	slice_steg_5 := AlleTallMultiplumAv(5, 1000)
	slice_disjunkte := ReturnDisjunkteMengde(slice_steg_3, slice_steg_5)		

	sum := SumAvElementerIMengdene(slice_disjunkte)
	fmt.Println("\nSummen av alle elementene i mengdene er:", sum)

}

func AlleTallMultiplumAv(steg int, grense int) []int {	
	retur_liste := []int{}

	for i := steg; i < grense; i += steg {
		retur_liste = append(retur_liste, i)
	}
	return retur_liste
}

func ReturnDisjunkteMengde(set1 []int, set2 []int) []int {
	returnVerdi := []int{}
	// Lag hashmap med key value par
	mappe := make(map[int]bool)
	
	// Fyll ut hashmap ved nøkkel verdi og default true value
	for _, verdi := range set1 {
		mappe[verdi] = true
	}
	
	//Oppdatert verdi på alle nøkler som eksisterer i begge mengdene til false
	for _, verdi := range set2 {
		value, exists := mappe[verdi]
		if exists && value {
			mappe[verdi] = false
		} 
	}
	
	// Legg til alle nøkler med true value til retur liste
	for key, value := range mappe {
		if value {
			returnVerdi = append(returnVerdi, key)
		}
	}

	// slå sammen mengdene til 1 siden vi nå vet at de er disjunkte
	returVerdi := append(returnVerdi, set2...)
	
	return returVerdi
}

func SumAvElementerIMengdene(set1 []int) int {
	sum := 0
	
	for _, verdi := range set1 {
		sum += verdi
	}
		
	return sum
}	


