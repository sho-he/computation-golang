package simple

import "fmt"

type multiply struct {
	Left numberInterface
	Right numberInterface
}

func Multiply(l numberInterface, r numberInterface) *multiply {
	multiply := &multiply{Left: l, Right: r}
	return multiply
}

func (m *multiply) To_s() string {
	expression := fmt.Sprintf("%v + %v", m.Left.To_s(), m.Right.To_s())
	return expression
}

func (m *multiply) Inspect() {
	t := fmt.Sprintf("<< Multiply >>")
	fmt.Println(t)
}

func (m *multiply) Expression() {
	fmt.Println(m.Left.To_s())
}

func (m multiply) IsReducible() bool {
	res := true
	fmt.Println(res)
	return res
}