package simple

import "fmt"

type add struct {
	Left numberInterface
	Right numberInterface
}

func Add(l numberInterface, r numberInterface) *add {
	add := &add{Left: l, Right: r}
	return add
}

func (a *add) To_s() string {
	expression := fmt.Sprintf("%v + %v", a.Left.To_s(), a.Right.To_s())
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