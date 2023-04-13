package simple

import "fmt"

type variable struct {
	name string
}

func Variable() *variable {
	variable := &variable{}
	return variable
}

func (v *variable) To_s() string {
	expression := fmt.Sprintf("x")
	return expression
}

func (v *variable) Inspect() {
	t := fmt.Sprintf("<< Variable >>")
	fmt.Println(t)
}

func (v *variable) Expression() {
	fmt.Printf("x\n")
}

func (v *variable) IsReducible() bool {
	return true
}

func (v *variable) Reduce(env environment) numberInterface {
	return env[v.name]
}