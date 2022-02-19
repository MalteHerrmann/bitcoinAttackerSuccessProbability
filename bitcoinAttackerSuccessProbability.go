package main

import (
	"fmt"
	"math"
)

// Calculates the success probability for an attacker trying to harm
// the blockchain, depending on the probability q and number of transactions z
// in the chain.
func AttackerSuccessProbability(q float64, z int) float64 {
	p := 1.0 - q
	lambda := float64(z) * (q / p)
	sum := 1.0

	var poisson float64

	for k := 0; k <= z; k++ {
		poisson = poissonDensity(lambda, k)
		sum = sum - poisson*(1-math.Pow(q/p, float64(z-k)))
	}

	return sum
}

// Calculates the poisson density for the expected value lambda
// and the number of events k.
func poissonDensity(lambda float64, k int) float64 {
	poisson := math.Exp(-lambda)
	for i := 1; i <= k; i++ {
		poisson = poisson * (lambda / float64(i))
	}

	return poisson
}

func main() {
	q := 0.1

	for z := 0; z <= 3; z++ {
		res := AttackerSuccessProbability(q, z)
		fmt.Printf("q=%v, z=%v: %v\n", q, z, res)
	}
}
