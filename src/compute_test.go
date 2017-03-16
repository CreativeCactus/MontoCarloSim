package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleUse() {
	fmt.Printf("Calculating pairs for {2, 4}:\n")
	input := []int{2, 4}
	result := EstimateπByMonteCarlo(10, 10, 10, 10)
	fmt.Printf("Estimate: %d\n", result.PI)
	// Output:
	// Estimate: 3.0
}

func TestKnown(t *testing.T) {
	//Known pass
	EstimateπByMonteCarlo(2, 2, 2, 2)
	EstimateπByMonteCarlo(4, 4, 4, 4)
	EstimateπByMonteCarlo(8, 8, 8, 8)
	// Nothing to test, all outputs nondeterministic.

	//Known fail
	EstimateπByMonteCarlo(0, 1, 1, 1)
	EstimateπByMonteCarlo(-1, 1, 1, 1)

}

func TestRandom(t *testing.T) {
	//Set up PRNG
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	//Run with small random inputs and expect no crash
	for i := 0; i < 1<<15; i++ {
		EstimateπByMonteCarlo(r.Intn(10), r.Intn(10), r.Intn(10), r.Intn(10))
	}
}

func BenchmarkPairs(b *testing.B) {
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		EstimateπByMonteCarlo(32, 32, 32, 32)
	})
	b.StopTimer()
}
