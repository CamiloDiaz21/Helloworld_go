package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// import "fmt"

// func imprimir() {
// 	fmt.Println("Texto desde la funcion imprimir")

// }
// // func main() {
// 	fmt.Println("Texto desde la funcion main")
// 	imprimir()
// }

// func Hola_string(s string) string {
// 	return "Hola" + s
// }

// func main() {
// 	fmt.Println(Hola_string(" Esteban"))
// }

// func sumar(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	fmt.Println(sumar(3, 5))
// }

// func comparar_bool() {
// 	var a bool
// 	var b bool
// 	a = true
// 	b = false

// 	if a == b {
// 		fmt.Println("A y B son iguales")
// 	} else {
// 		fmt.Println("A y B no son iguales")
// 	}
// }

// func main() {
// 	comparar_bool()

// }

// func arreglo() {
// 	var aprendices [5]string
// 	aprendices[0] = "Leonadardo"
// 	aprendices[1] = "Juan Sebastian"
// 	aprendices[2] = "Frank"
// 	aprendices[3] = "Juan Jose"
// 	aprendices[4] = "Jaider"

// 	fmt.Println(aprendices[1])
// }

// func main() {
// 	arreglo()
// }

// func tipos_datos() {
// 	var texto string = "Esteban"
// 	var numero int = 1000
// 	var decimal float64 = 0.0001

// 	fmt.Println(reflect.TypeOf(texto))
// 	fmt.Println(reflect.TypeOf(numero))
// 	fmt.Println(reflect.TypeOf(decimal))
// }

// func main() {
// 	// tipos_datos()
// }

func convertir_string_boolean() {
	var palabra string = "false"
	boolean, err := strconv.ParseBool(palabra)
	if err != nil {
		fmt.Println("Este es el error ", err)
	} else {
		fmt.Println(boolean, reflect.TypeOf(boolean))
	}
}

// func main() {
	// convertir_string_boolean()
// }

// func covertir_boolean_string() {
// 	var boolean bool = true
// 	strbool := strconv.FormatBool(boolean)
// 	fmt.Println(strbool, reflect.TypeOf(strbool))
// }

// func main() {
// 	covertir_boolean_string()
// }

// Ejercicio 15
// func variables(){
// 	var (
// 	nombre	 string = "Camilo"
// 	apellido string = "Diaz"
// 	años	 int     = 18
// 	)

// 	fmt.Println("Nombre:", nombre)
// 	fmt.Println("Apellido:", apellido)
//  	fmt.Println("Edad:", años)
// }

//  func main() {
//     variables()
//  }

// ejercicio 16
func valor_cero() {
	var nombre string 
	var edad int 
	var peso float64 
	var estudiate bool 
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Peso:", peso)
	fmt.Println("Estudiante:", estudiate)
}

//  func main() {
//     valor_cero()
//  }

// Ejercicio 17

func Declaracion_variables() {
    nombre := "Cristian Camilo"
    edad := 18
    peso := 65
    estudiante := true
    fmt.Println("Nombre: ", nombre)
    fmt.Println("Edad: ", edad)
    fmt.Println("Peso: ", peso)
    fmt.Println("Estudiante: ", estudiante)
}

  func main() {
	 Declaracion_variables()
}
// ejercicio 18 
var var1 = "Nivel 1"
 
func scope(){
	var var2 = "Nivel 2"
	{
		var var3 = "Nivel 3"
		fmt.Println(var3)
	}
    fmt.Println(var1)
    fmt.Println(var2)
}
//   func main() {
 	// scope()
//   }

// ejercicio 19 
func punteros() {

    Manzana1 := "roja"
    fmt.Println("la manzana 1 es: ", Manzana1)

    sandia := "varde"
    fmt.Println("la sandia es: ", sandia)

    manzana2 := &Manzana1
    fmt.Println("la manzana 2 es: ", *manzana2)

    Manzana1 = "Roja"

    fmt.Println("La manzana 1 es:", Manzana1, &Manzana1)
    fmt.Println("La sandia es:", sandia, &sandia)
    fmt.Println("La manzana 2 es:", *manzana2, manzana2)

}
//  func main() {
	// punteros()
//  }

func copia_punteros(altura float32) float32 {
	altura = altura  * 3.78 
	return altura 
}

var altura float32 = 1.70
// func main(){
	//  copia_punteros(altura)
    
//   fmt.Println("La altura es:", altura, " mts")
//   fmt.Println("la altura es:", copia_punteros(altura), " pies")
//   fmt.Println("La altura nueva es: ", altura, " mts")
//   }
// ejercicio 21


// 	const pi float64 = 3.14159
  func constante(){
  fmt.Println("El valor de pi es:", pi, "mts")
  }
// func main (){
//      constante()
// }


//  const pi = 3.14159
func uso_funciones(radio float64) float64{
	 return pi * radio * radio 
 }
 func radio(){
 	fmt.Println("El area de un circulo cuyo radio es 3 es: ", uso_funciones(5))
}

// func main(){
//      radio()
// 	}

// Ejercicio 23

const pi = 3.14159
func circulo(radio float64)(area float64, perimetro float64){
area = pi * radio * radio 
perimetro = 2 * pi * radio 
return area, perimetro 
}
func multiples_valores(){
	area, perimetro:= circulo(8)
	fmt.Println("el area del circulo es:", area)
	fmt.Println("El perimetro del circulo es:", perimetro)

}

// func main(){
// 	multiples_valores()
// }