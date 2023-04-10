package simple

import "fmt"

type Multiply struct {
	Left Number
	Right Number
}

func (m *Multiply) to_s() string {
	expression := fmt.Sprintf("%v * %v", m.Left.Value, m.Right.Value)
	fmt.Println(expression)
	return expression
}

func (m *Multiply) Inspect() string {
	t := fmt.Sprintf("<< %T >>", m)
	fmt.Println(t)
	return t
}