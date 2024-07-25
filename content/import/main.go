package main

//import "fmt"
import (
	"fmt"
)

const a = "Otávio!"

var b bool

// criando um type
type ID int64

// pode ter declaração e atribuição
var (
	d             bool
	e             string
	f             int
	g             float64
	h             int = 3
	identificador ID  = 10000001
)

func main() {
	fmt.Printf("O tipo de <identificador> é %t", identificador)
	fmt.Println()
	fmt.Printf("O valor de <identificador> é %v", identificador)
}
