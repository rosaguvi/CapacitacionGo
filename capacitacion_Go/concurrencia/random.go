package main

import (
	"fmt"
	"math/rand"
)

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func main() {
	r := rand.New(rand.NewSource(55))
	ir := IntRange{-10, 10}
	for i := 0; i < 100; i++ {
		fmt.Println(ir.NextRandom(r))
	}
}
