package main

import (
	"fmt"
)

func quicksort(array []int) []int {
	
	if len(array) < 2 {
		return array
	}

	pivo := array[int(len(array) / 2)] // pivo como elemento central do slice

	menoresQuePivo := CriaArrayComMenoresQuePivo(array, pivo)
	maioresQuePivo := CriarArrayComMaioresQueOpivo(array, pivo)

	menoresOrdenados := quicksort(menoresQuePivo)
	maioresOrdenados := quicksort(maioresQuePivo)


	menoresOrdenados = append(menoresOrdenados, pivo)

	return append(menoresOrdenados, maioresOrdenados...)
	
	
}

func CriaArrayComMenoresQuePivo(slice []int, pivo int) []int {
	var values []int
	for _, v := range slice {
		if v < pivo {
			values = append(values, v)
		}
	}

	return values
}

func CriarArrayComMaioresQueOpivo(slice []int, pivo int) []int {
	var values []int
	for _, v := range slice {
		if v > pivo {
			values = append(values, v)
		}
	}
	return values 
}

func main() {
	slice := []int{5,2,56,765,78,2541,6,3,23,8,7}

	fmt.Println(quicksort(slice))
}