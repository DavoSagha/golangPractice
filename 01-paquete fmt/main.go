package main

import "fmt"

func main() {
	fmt.Println("Hola mundo!")
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Alex"
	edad := 26

	fmt.Printf("Hola %s, vos tenes %d \n", nombre, edad)
	fmt.Printf("Hola %v, vos tenes %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola %v, vos tenes %v", nombre, edad)

	fmt.Println(mensaje)
	fmt.Printf("Nombre: %T \n", nombre) //me dice qué es nombre, en este caso string

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("El otro nombre es: ", nombre)
}
