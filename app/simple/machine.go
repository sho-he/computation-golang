package simple

import "fmt"

type machine struct {
	expression numberInterface
	environment environment
}

func Machine(e numberInterface, env environment) *machine {
	machine := &machine{expression: e, environment: env}
	return machine
}

func (m *machine) Step(expression numberInterface) numberInterface {
	return expression.Reduce(m.environment)
}

func (m *machine) Run() {
	for m.expression.IsReducible() {
		fmt.Println(m.expression.To_s())
		m.expression = m.Step(m.expression)
	}
	fmt.Println(m.expression.To_s())
}