package simple

type numberInterface interface {
	To_s() string
	Inspect() 
	IsReducible() bool
	Expression()
	Reduce() numberInterface
	getNumber() int
}