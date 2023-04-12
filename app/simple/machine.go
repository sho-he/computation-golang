package simple

import "fmt"

type machine struct {
	expression numberInterface
}

func Machine(e numberInterface) *machine {
	machine := &machine{expression: e}
	return machine
}

func (m *machine) Step(expression numberInterface) numberInterface {
	return expression.Reduce()
}

func (m *machine) Run() {
	for m.expression.IsReducible() {
		fmt.Println(m.expression.To_s())
		m.expression = m.Step(m.expression)
	}
	fmt.Println(m.expression.To_s())
}