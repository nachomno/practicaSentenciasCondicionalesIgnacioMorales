package main

import "fmt"

func main() {
	// ------------1---------
	fmt.Print("-------- Ejercicio 1 -----------")
	var numero int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Println("El número es mayor que cero.")
	} else if numero == 0 {
		fmt.Println("El número es igual a cero.")
	} else {
		fmt.Println("El número es menor que cero.")
	}
	// ------------2---------
	fmt.Print("-------- Ejercicio 2 -----------")
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("El número es par.")
	} else {
		fmt.Println("El número es impar.")
	}
	// ------------3---------
	fmt.Print("-------- Ejercicio 3 -----------")
	var temp1, temp2, temp3 float64
	fmt.Print("Ingrese la primera temperatura: ")
	fmt.Scan(&temp1)
	fmt.Print("Ingrese la segunda temperatura: ")
	fmt.Scan(&temp2)
	fmt.Print("Ingrese la tercera temperatura: ")
	fmt.Scan(&temp3)

	suma := temp1 + temp2 + temp3
	promedio := suma / 3

	fmt.Printf("La suma de las temperaturas es: %.2f\n", suma)
	fmt.Printf("El promedio de las temperaturas es: %.2f\n", promedio)
	
	// ------------4---------
	fmt.Print("-------- Ejercicio 4 -----------")
	var dias, horas, minutos, segundos int
	fmt.Print("Ingrese el número de días: ")
	fmt.Scan(&dias)
	fmt.Print("Ingrese el número de horas: ")
	fmt.Scan(&horas)
	fmt.Print("Ingrese el número de minutos: ")
	fmt.Scan(&minutos)
	fmt.Print("Ingrese el número de segundos: ")
	fmt.Scan(&segundos)

	totalSegundos := dias*86400 + horas*3600 + minutos*60 + segundos

	fmt.Printf("El periodo en segundos es: %d\n", totalSegundos)
	// ----------5--------------
	fmt.Print("-------- Ejercicio 5 -----------")
	var X int
	fmt.Print("Ingrese el número X: ")
	fmt.Scan(&X)
	S1 := min(X, 50)
	X -= S1
	S2 := min(X, 50)
	X -= S2
	S3 := min(X, 600)
	X -= S3
	S4 := min(X, 800)
	X -= S4
	S5 := X
	fmt.Printf("S1 = %d, S2 = %d, S3 = %d, S4 = %d, S5 = %d\n", S1, S2, S3, S4, S5)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}