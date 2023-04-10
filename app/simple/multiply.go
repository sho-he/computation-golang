package simple

import "fmt"

type multiply struct {
	Left number
	Right number
}

func Multiply(l number, r number) *multiply {
	multiply := &multiply{Left: l, Right: r}
	return multiply
}

func (m *multiply) to_s() string {
	expression := fmt.Sprintf("%v * %v", m.Left.Value, m.Right.Value)
	fmt.Println(expression)
	return expression
}

func (m *multiply) Inspect() string {
	t := fmt.Sprintf("<< Multiply >>")
	fmt.Println(t)
	return t
}

func (m multiply) IsReducible() bool {
	res := true
	fmt.Println(res)
	return res
}