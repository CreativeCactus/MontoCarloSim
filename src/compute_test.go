package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleUse() {
	result, _ := EstimateπByMonteCarlo(10, 10, 10, 10)
	fmt.Printf("Estimate: %f\n", result.PI)
}

func TestKnown(t *testing.T) {
	//Known pass
	if _, err := EstimateπByMonteCarlo(2, 2, 2, 2); err != nil {
		t.Errorf("Expected success but got error: %s", err.Error())
	}
	if _, err := EstimateπByMonteCarlo(4, 4, 4, 4); err != nil {
		t.Errorf("Expected success but got error: %s", err.Error())
	}
	if _, err := EstimateπByMonteCarlo(8, 8, 8, 8); err != nil {
		t.Errorf("Expected success but got error: %s", err.Error())
	}

	//Known fail
	if result, err := EstimateπByMonteCarlo(0, 1, 1, 1); err == nil {
		t.Errorf("Expected error but got output: %+v", result)
	}
	if result, err := EstimateπByMonteCarlo(-1, 1, 1, 1); err == nil {
		t.Errorf("Expected error but got output: %+v", result)
	}
}

func TestRandom(t *testing.T) {
	//Set up PRNG
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	//Run with small random inputs and expect no crash
	for i := 0; i < 1<<12; i++ {
		if _, err := EstimateπByMonteCarlo(1+r.Intn(10), 1+r.Intn(10), 1+r.Intn(10), 1+r.Intn(10)); err != nil {
			t.Errorf("Expected success but got error: %s", err.Error())
		}
	}
}

func BenchmarkPairs(b *testing.B) {
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		EstimateπByMonteCarlo(32, 32, 32, 32)
	})
	b.StopTimer()
}
