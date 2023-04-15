package simple

import "fmt"

type assign struct {
	name string
	expression numberInterface
}

func Assign(name string, expression numberInterface) *assign {
	return &assign{name: name, expression: expression}
}

func (a *assign) To_s() string {
	return a.name + " = " + a.expression.To_s()
}

func (a *assign) Inspect() {
	t := "<< Assign >>"
	fmt.Println(t)
}

func (a *assign) IsReducible() bool {
	return true
}

func (a *assign) Expression(env environment) {
	fmt.Printf("%v = %v %v\n", a.name, a.expression.To_s(), env)
}

func (a *assign) Reduce(e environment) stateInterface {
	if a.expression.IsReducible() {
		return Assign(a.name, a.expression.Reduce(e))
	} else {
		number := a.expression
		fmt.Printf("%v", number)
		e[a.name] = number
		return DoNothing(e)
	}
}
