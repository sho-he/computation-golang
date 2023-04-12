package main

import "golang_computation/simple"

func main () {
		simple.Machine(simple.Multiply(simple.Add(simple.Number(1), simple.Number(2)), simple.Number(3))).Run()
}
