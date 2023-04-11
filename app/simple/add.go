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
	expression := fmt.Sprintf("( %v + %v )", a.Left.To_s(), a.Right.To_s())
	return expression
}

func (a *add) Inspect() {
	t := fmt.Sprintf("<< Add >>")
	fmt.Println(t)
}

func (a *add) Expression() {
	fmt.Printf("%v + %v\n", a.Left.To_s(), a.Right.To_s())
}

func (a *add) IsReducible() bool {
	return true
}

func (a *add) Reduce() numberInterface {
	if a.Left.IsReducible() {
		result := Add(a.Left.Reduce(), a.Right)
		return result
	} else if a.Right.IsReducible() {
		result := Add(a.Left, a.Right.Reduce())
		return result
	} else {
		l := a.Left.getNumber()
		r := a.Right.getNumber()
		return Number(l + r)
	}
}

func (a *add) getNumber() int {
	return a.Left.getNumber() + a.Right.getNumber()
}