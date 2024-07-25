package function1

import (
	"errors"
	"fmt"
)

//import "fmt"

func main() {
	fmt.Println(soma_com_mais_de_um_retorno(3, 5))

	soma_processamento()

	fmt.Println("\nSoma variatica: ", soma_variatica(1, 2, 3, 4))
	fmt.Println("\nSoma closures: ", soma_closures())

}

// função com retormo
func Soma(a int, b int) int {
	return a + b
}

func soma_2_forma_declaracao(a, b int) int {
	return a + b
}

func soma_com_mais_de_um_retorno(a, b int) (int, bool) {
	if a+b > 30 {
		return a + b, true
	} else {
		return a + b, false
	}
}

// formas de usar exceptions. Não tem try/catch
func soma_e_exception(a, b int) (int, error) {
	if a+b > 30 {
		return a + b, errors.New("Erro na operação matemica!")
	} else {
		return a + b, nil
	}
}

func soma_processamento() {
	soma, erro := soma_e_exception(12, 23)

	if soma > 4 {
		fmt.Println("Soma bem sucedida!")

	} else {
		fmt.Println(erro)
	}

}

// -----------------------------------

//FUNÇÕES VARIATICAS

func soma_variatica(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

//FUNÇÕES anonimas ou closures

func soma_closures() int {
	soma := func() int {
		return soma_variatica(100, 200) * 2
	}()
	return soma
}
