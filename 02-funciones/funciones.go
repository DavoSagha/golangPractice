package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func main() {
	var nombre string
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)
	saludar(nombre)

}
