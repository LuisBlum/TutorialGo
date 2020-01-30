package main

import "fmt"

func main() {
	//DECLARACIONES
	var Titulo string
	var PesoKg float32
	var Estatura float32
	var IMC float32

	//ASIGNACIONES
	Titulo = "Calculadora IMC"

	fmt.Println("*********************************************")
	fmt.Println(" ** " + Titulo + " **")
	fmt.Println("*********************************************")
	fmt.Println(" ")

	fmt.Printf("Cual es su peso en kg? ")
	fmt.Scanf("%v\n", &PesoKg)
	fmt.Printf("Cual es su estatura en metros? ")
	fmt.Scanf("%v\n", &Estatura)
	fmt.Println(" ")

	IMC = PesoKg / (Estatura * Estatura)
	fmt.Printf("Su IMC es %v \n", IMC)
}
