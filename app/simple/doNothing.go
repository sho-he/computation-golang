package simple

import "fmt"

type doNothing struct {
}

func DoNothing() *doNothing {
	return &doNothing{}
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
