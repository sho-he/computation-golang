package main

import "golang_computation/simple"

func main () {
		simple.Machine(simple.Multiply(simple.Add(simple.Variable("x"), simple.Number(2)),simple.Number(3)),simple.Environment("x", simple.Number(2))).Run()
		simple.Machine(simple.LessThan(simple.Add(simple.Number(1), simple.Number(2)), simple.Number(3)),simple.Environment("x", simple.Number(2))).Run()

		statement := simple.Assign("x", simple.Add(simple.Variable("x"), simple.Number(1)))
		statement.Inspect()
}
