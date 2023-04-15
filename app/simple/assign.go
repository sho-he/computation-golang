package simple

import "fmt"

type assign struct {
	name string
	expression numberInterface
}

func Assign(name string, expression numberInterface) *assign {
	return &assign{name: name, expression: expression}
}

func (a *assign) To_s(name string, expression numberInterface) string {
	return name + " = " + expression.To_s()
}

func (a *assign) Inspect() {
	t := "<< Assign >>"
	fmt.Println(t)
}

func (a *assign) IsReducible() bool {
	return true
}

func (a *assign) Reduce(e environment) (stateInterface, environment) {
	if a.expression.IsReducible() {
		return Assign(a.name, a.expression.Reduce(e)), e
	} else {
		return DoNothing(), Environment(a.name, a.expression)
	}
}
