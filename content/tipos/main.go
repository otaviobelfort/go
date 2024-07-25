package main

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
	println(identificador)
}
