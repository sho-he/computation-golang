package simple

import "fmt"

type doNothing struct {
	environment environment
}

func DoNothing(env environment) *doNothing {
	return &doNothing{environment: env}
}

func (d *doNothing) To_s() string {
	return "do-nothing"
}

func (d *doNothing) Inspect() {
	t := "<< DoNothing >>"
	fmt.Println(t)
}

func (d *doNothing) IsReducible() bool {
	return false
}

func (d *doNothing) Expression(env environment) {
	fmt.Printf(d.To_s() + " %v\n", env)
}

func (d *doNothing) Reduce(env environment) stateInterface {
	return d
}