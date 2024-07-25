package main

//import "fmt"
import (
	"fmt"
)

func main() {

	//percorrerArray()
	//slice()
	maps()
}

// array
func percorrerArray() {

	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println("O tamanho do array é: ", len(array))

	var arrayNumeros [3]int
	arrayNumeros[0] = 123
	arrayNumeros[1] = 23
	arrayNumeros[2] = 34
	for i, v := range arrayNumeros {
		fmt.Printf("O valor é %d e o indice é %d \n", i, v)
	}

}

func slice() {
	// é um array com capacidade dinamica
	slice := []int{12, 22, 12, 33, 445, 5543}

	fmt.Printf("Tamanho: %d , Cap: %d %v", len(slice), cap(slice), slice)
	fmt.Printf("Tamanho: %d , Cap: %d %v", len(slice[:2]), cap(slice[:2]), slice[2:])
	// para cada append ele dobra a capácidade
	slice = append(slice, 100)
	fmt.Printf("\n Tamanho: %d , Cap: %d %v", len(slice), cap(slice), slice)

}

func maps() {
	// inicializacao
	//pessoas := map[string]int{}
	//pessoas_fisica := make(map[string]int)

	salarios := map[string]int{"Otávio": 100000, "João": 20000}
	fmt.Println("Salários: ", salarios["Otávio"])

	delete(salarios, "Otávio")
	salarios["Jose"] = 5000

	for i, v := range salarios {
		fmt.Println(i, v)

	}

	// ignorar chave ou valor `_`
	for i, _ := range salarios {
		fmt.Println(i)

	}

	// ignorar chave ou valor `_`
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}
}
