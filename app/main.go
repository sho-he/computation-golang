package main

import "golang_computation/expression"

func main () {
		result := simple.Add{Left: simple.Number{Value: 1}, Right: simple.Number{Value: 2}}
		result.To_s()
		result.Inspect()
}
