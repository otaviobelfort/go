package math

type Math struct {
	A int
	B int
}

func (soma Math) Soma() int {
	return soma.A + soma.B
}
