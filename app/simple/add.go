package simple

import "fmt"

type Add struct {
	Left Number
	Right Number
}

func (a *Add) To_s() string {
	expression := fmt.Sprintf("%d + %d", a.Left.Value, a.Right.Value)
	fmt.Println(expression)
	return expression
}

func (a *Add) Inspect() string {
	t := fmt.Sprintf("<< %T >>", a)
	fmt.Println(t)
	return t
}