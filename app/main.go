package main

import "golang_computation/expression"

func main () {
		a := simple.Add{Left: simple.Number{Value: 1}, Right: simple.Number{Value: 2}}
		a.To_s()
		n := simple.Number{Value: 1}
		n.Inspect()
}
