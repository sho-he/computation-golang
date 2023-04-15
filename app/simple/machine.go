package simple

type machine struct {
	expression numberInterface
	environment environment
}

func Machine(e numberInterface, env environment) *machine {
	machine := &machine{expression: e, environment: env}
	return machine
}

func (m *machine) Step(expression numberInterface, e environment) numberInterface {
	return expression.Reduce(e)
}

func (m *machine) Run() {
	for m.expression.IsReducible() {
		m.expression.Expression(m.environment)
		m.expression = m.Step(m.expression, m.environment)
	}
	m.expression.Expression(m.environment)
}

type sMachine struct {
	expression stateInterface
	environment environment
}

func SMachine(e stateInterface, env environment) *sMachine {
	machine := &sMachine{expression: e, environment: env}
	return machine
}

func (m *sMachine) Step(expression stateInterface, e environment) stateInterface {
	return expression.Reduce(e)
}

func (m *sMachine) Run() {
	for m.expression.IsReducible() {
		m.expression.Expression(m.environment)
		m.expression = m.Step(m.expression, m.environment)
	}
	m.expression.Expression(m.environment)
}