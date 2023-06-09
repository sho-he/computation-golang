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
	expression := fmt.Sprintf("(%v * %v)", m.Left.To_s(), m.Right.To_s())
	return expression
}

func (m *multiply) Inspect() {
	t := fmt.Sprintf("<< Multiply >>")
	fmt.Println(t)
}

func (m *multiply) Expression(e environment) {
	fmt.Printf("(%v * %v)\n", m.Left.To_s(), m.Right.To_s())
}

func (m *multiply) IsReducible() bool {
	return true
}

func (m *multiply) Reduce(env environment) numberInterface {
	if m.Left.IsReducible() {
		result := Multiply(m.Left.Reduce(env), m.Right)
		return result
	} else if m.Right.IsReducible() {
		result := Multiply(m.Left, m.Right.Reduce(env))
		return result
	} else {
		l := m.Left.getNumber()
		r := m.Right.getNumber()
		return Number(l * r)
	}
}

func (m *multiply) getNumber() int {
	return m.Left.getNumber() + m.Right.getNumber()
}