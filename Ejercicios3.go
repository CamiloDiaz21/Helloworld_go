package main

import "fmt"
import "strings"

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

// Ejercicio 32

func operaciones_aricmeticas(){
	x, y := -5, 10
    
	if x > 0 || y > 0{
		fmt.Println("ambas funciones son verdaderas")
	}else{
       fmt.Println("no se cumple el AND")
	}
}
// func main(){
// 	operaciones_aricmeticas()
// }

func contraseña(){
	var contraseña string
	fmt.Print("Por favor, ingrese su contraseña: ")
	fmt.Scanln(&contraseña)

	
	if !strings.ContainsAny(contraseña, "abcdefghijklmnopqrstuvwxyz") || !strings.ContainsAny(contraseña, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") || !strings.ContainsAny(contraseña, "0123456789")|| !strings.ContainsAny(contraseña, "°!#$%&/=?¡*@.,+-") || len(contraseña) <=8 {
  	   fmt.Println("La contraseña no cumple con los requisitos.")
} else {
  fmt.Println("La contraseña es válida.")
}
}

// func main() {
// 	contraseña()
// }

func switche(){
var juguete string
    fmt.Println("Elija que tipo de juguete agregar? persona, animal o cosa")
    fmt.Scanln(&juguete)
    switch juguete {
    case "persona":
        fmt.Println("El juguete es una figura de accion")
    case "cosa":
        fmt.Println("El juguete es una cosa")
    case "animal":
        fmt.Println("El juguete es una mascota")
    default:
        fmt.Println("El juguete es otra categoria")
    }
}
// func main() {
// 	switche()
// }

// Ejercicio 36 

func desconectar() {
    fmt.Println("Se ha desconectado de la base de datos")
}

func leer() {
    fmt.Println("Se han leido los registros de la base de datos")
}

func actualizar() {
    fmt.Println("Se han actalizado registros de la base de datos")
}

func conectar() {
    fmt.Println("Se ha conectado a la base de datos")
}

// func main() {
//     conectar()
//     defer desconectar()
//     leer()
//     actualizar()
// }

// Ejercicio 38 

// func main() {
//     var marcasDeCoches = make([]string, 2)
//     marcasDeCoches[0] = "Mazda"
//     marcasDeCoches[1] = "Toyota"
//     fmt.Println(marcasDeCoches) // [Mazda Toyota]
//     nuevoSlice := append(marcasDeCoches, "Nissan")
//     fmt.Println(nuevoSlice) // [Mazda Toyota Nissan]
// }

// Ejercicio 39 

// func main() {
//     razasDePerros := []string{"labrador", "poodle", "doberman", "shitzu", "beagle"}
//     fmt.Println(razasDePerros)
//     razasDePerros = append(razasDePerros[:2], razasDePerros[2+1:]...)
//     fmt.Println(razasDePerros)
// }

// Ejercicio 40

// func main() {
//     figuras := []string{"circulo", "cuadrado", "triangulo", "rombo", "trapecio", "heptagono"}
//     var figuras2 = make([]string, len(figuras))
//     fmt.Println(figuras)
//     copy(figuras2, figuras)
//     figuras = append(figuras[:1], figuras[2:]...)
//     fmt.Println(figuras2)
//     fmt.Println(figuras)
// }

// Ejercicio 41 
// func main() {
//     var diasDeLaSemana = make(map[string]int)
//     diasDeLaSemana["lunes"] = 1
//     diasDeLaSemana["martes"] = 2
//     diasDeLaSemana["miercoles"] = 3
//     diasDeLaSemana["jueves"] = 4
//     diasDeLaSemana["viernes"] = 5
//     diasDeLaSemana["sabado"] = 6
//     diasDeLaSemana["domingo"] = 7
//     fmt.Println(diasDeLaSemana)
// }

// Ejercicio 42 

// func main() {
//     diasDeLaSemanaEnIngles := map[string]string{
//         "lunes":     "Monday",
//         "martes":    "Tuesday",
//         "miercoles": "Wednesday",
//         "jueves":    "Thursday",
//         "sabado":    "Saturday",
//         "domingo":   "Sunday",
//     }
//     fmt.Println(diasDeLaSemanaEnIngles)
//     delete(diasDeLaSemanaEnIngles, "domingo")
//     fmt.Println(diasDeLaSemanaEnIngles)
// }