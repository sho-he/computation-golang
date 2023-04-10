package main

import "golang_computation/simple"

func main () {
		result := simple.Add(simple.Number(1), simple.Number(2))
		result.To_s()
		result.Inspect()
		result.IsReducible()
}
