package main

import "fmt"

// Ejercicio 23

const Pi = 3.1416

func circulo2(radio float64) (area float64, perimetro float64) {
	area = Pi * radio * radio
	perimetro = 2 * Pi * radio
	return area, perimetro
}

// func main() {
// 	a, p := circulo2(8)
// 	fmt.Println("El area del circulo es: ", a)
// 	fmt.Println("El perimetro del circulo es: ", p)
// }

// Ejercicio 24
func sumar(numeros ...int) int {
	// el total inicial es 0
	total := 3000
	// recorrer todos los numeros
	for _, numero := range numeros {
		// en cada iteración sumar al total el valor del numero
		total = numero * total
	}
	// retornar el valor total
	return total
}

func variadicas() {
	fmt.Println(sumar(55))
	fmt.Println(sumar(55))
	fmt.Println(sumar(78))
}

// func main() {
// 	variadicas()
// }

// Ejercicio 25

func ecoDeLaMontana(mensaje string, iteraciones int) {
	if iteraciones > 1 {
		ecoDeLaMontana(mensaje, iteraciones-1)
	}
	fmt.Println(mensaje, iteraciones)
}

// func main() {
// 	ecoDeLaMontana("Hola", 5)
// }

func circulo1(radio float64) (area func() float64, perimetro func() float64) {

	area = func() float64 {
		return 3.1416 * radio * radio
	}

	perimetro = func() float64 {
		return 2 * 3.1416 * radio
	}

	return
}

// func main() {

// 	area, perimetro := circulo1(8)
// 	fmt.Println("El area del circulo es", area())
// 	fmt.Println("El perimetro del circulo es", perimetro())

// }