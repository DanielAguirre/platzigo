// Imprimer un arreglo y un mapa usando range para iterar el arreglo y el mapa
// 
package main

import "fmt"

func main() {
	numeros := []int{2,4,6}
	suma := 0
	for _, num := range numeros {
		suma += num
	}

	fmt.Println("suma:", suma)

	for i, numero := range numeros {
		if numero == 3{
			fmt.Println("index:", i)
		}
	}

	algo := map[string]string{"a":"auto", "b": "bebé"}

	// Itera en un mapa en imprime la llaves y valores del mapa
	for k, v:= range algo {
		fmt.Printf("Key: %s Value: %s,\n", k, v)
	}
}