package main

import "golang_computation/simple"

func main () {
		result := simple.Multiply(simple.Add(simple.Number(1), simple.Number(2)), simple.Number(3))
		result.Expression()
		result.Inspect()
		result2 := result.Reduce()
		result2.Expression()
}
