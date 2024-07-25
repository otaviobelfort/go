package main

const a = "Olá!"

var b bool

// pode ter declaração e atribuição
var (
	d bool
	e string
	f int
	g float64
	h int = 3
)

func main() {
	b = true
	println(b)
}

// a var i foi declarada e nao utilizada. O compilador acusa erro
// com package pode acontecer o mesmo erro
func declaracaoEnaoUsoDeVar() {
	//var i bool = true
	println(a)

	// declaração e atribuição dinamica short
	// na primeira vez usa-se `:`
	// nas proximas atribuições da mesma variavel é necessario usar o `:`
	j := "Short"
	println(j)
}
