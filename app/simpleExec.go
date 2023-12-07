package main

import "golang_computation/simple"

func main () {
		simple.Machine(simple.Multiply(simple.Add(simple.Variable("x"), simple.Number(2)),simple.Number(3)),simple.Environment("x", simple.Number(2))).Run()
		simple.Machine(simple.LessThan(simple.Add(simple.Number(1), simple.Number(2)), simple.Number(3)),simple.Environment("x", simple.Number(2))).Run()

		simple.SMachine(simple.Assign("x", simple.Add(simple.Number(1), simple.Number(2))), simple.Environment("x", simple.Number(2))).Run()
		simple.SMachine(simple.Sequence(simple.Assign("x", simple.Add(simple.Number(1), simple.Number(2))),simple.Assign("x", simple.Add(simple.Number(1), simple.Number(2)))), simple.Environment("x", simple.Number(2))).Run()
		simple.SMachine(simple.If(simple.Variable("x"), simple.Assign("x", simple.Add(simple.Number(1), simple.Number(2))), simple.Assign("x", simple.Add(simple.Number(4), simple.Number(5)))), simple.Environment("x", simple.Boolean(true), )).Run()
}
