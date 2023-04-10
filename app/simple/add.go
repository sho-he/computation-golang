package simple

import "fmt"

type add struct {
	Left number
	Right number
}

func Add(l number, r number) *add {
	add := &add{Left: l, Right: r}
	return add
}

func (a *add) To_s() string {
	expression := fmt.Sprintf("%d + %d", a.Left.Value, a.Right.Value)
	fmt.Println(expression)
	return expression
}

func (a *add) Inspect() string {
	t := fmt.Sprintf("<< Add >>")
	fmt.Println(t)
	return t
}

func (a add) IsReducible() bool {
	res := true
	fmt.Println(res)
	return res
}