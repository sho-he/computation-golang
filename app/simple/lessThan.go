package simple

import "fmt"

type lessThan struct {
	Left numberInterface
	Right numberInterface
}

func LessThan(l numberInterface, r numberInterface) *lessThan {
	lessThan := &lessThan{Left: l, Right: r}
	return lessThan
}

func (g *lessThan) To_s() string {
	expression := fmt.Sprintf("(%v < %v)", g.Left.To_s(), g.Right.To_s())
	return expression
}

func (g *lessThan) Inspect() {
	t := fmt.Sprintf("<< lessThan >>")
	fmt.Println(t)
}

func (g *lessThan) Expression() {
	fmt.Printf("%v > %v\n", g.Left.To_s(), g.Right.To_s())
}

func (g *lessThan) IsReducible() bool {
	return true
}

func (g *lessThan) Reduce(env environment) numberInterface {
	if g.Left.IsReducible() {
		result := LessThan(g.Left.Reduce(env), g.Right)
		return result
	} else if g.Right.IsReducible() {
		result := LessThan(g.Left, g.Right.Reduce(env))
		return result
	} else {
		l := g.Left.getNumber()
		r := g.Right.getNumber()
		return Boolean(l < r)
	}
}

func (g *lessThan) getNumber() int {
	return g.Left.getNumber() + g.Right.getNumber()
}