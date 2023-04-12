package simple

import "fmt"

type greaterThan struct {
	Left numberInterface
	Right numberInterface
}

func GreaterThan(l numberInterface, r numberInterface) *greaterThan {
	greaterThan := &greaterThan{Left: l, Right: r}
	return greaterThan
}

func (g *greaterThan) To_s() string {
	expression := fmt.Sprintf("( %v > %v )", g.Left.To_s(), g.Right.To_s())
	return expression
}

func (g *greaterThan) Inspect() {
	t := fmt.Sprintf("<< greaterThan >>")
	fmt.Println(t)
}

func (g *greaterThan) Expression() {
	fmt.Printf("%v > %v\n", g.Left.To_s(), g.Right.To_s())
}

func (g *greaterThan) IsReducible() bool {
	return true
}

func (g *greaterThan) Reduce() numberInterface {
	if g.Left.IsReducible() {
		result := Add(g.Left.Reduce(), g.Right)
		return result
	} else if g.Right.IsReducible() {
		result := Add(g.Left, g.Right.Reduce())
		return result
	} else {
		l := g.Left.getNumber()
		r := g.Right.getNumber()
		return Number(l + r)
	}
}

func (g *greaterThan) getNumber() int {
	return g.Left.getNumber() + g.Right.getNumber()
}