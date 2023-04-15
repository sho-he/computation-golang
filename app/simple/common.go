package simple

type numberInterface interface {
	To_s() string
	Inspect() 
	IsReducible() bool
	Expression()
	Reduce(env environment) numberInterface
	getNumber() int
}

type environment map[string]numberInterface

func Environment(variable string, value numberInterface) environment { 
	env := environment{variable: value}
	return env
}

type stateInterface interface {
}