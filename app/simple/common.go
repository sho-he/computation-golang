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