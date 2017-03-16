package main

// go get github.com/mitsuse/matrix-go
import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"sync"
	"time"

	"math"
	"math/rand"

	"image"
	"image/color"
	"image/png"

	"encoding/base64"
	"encoding/json"
)

type model struct {
	PI  float64 `json:"pi"`
	PNG string  `json:"png"`
}

func main() {
	// Use `Carlc -h` for generatd help.
	grid := flag.Int("grid", 1000, "The edge size of the plane.")
	circ := flag.Int("circ", 900, "The diameter of the simulated circle.")
	pts := flag.Int("pts", 100, "The number of random points.")
	its := flag.Int("its", 10, "The number of iterations.")
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Println(`Example: ./compute -grid=1000 -circ=900 -pts=100 -its=10
		Expect a result in JSON: {PI:3.14, PNG:*big blob of png data*}
		The result can be parsed using the unix command line tool jq.
		This module is intended for use with a frontend to allow scalable, distributed computation.`)
	}
	flag.Parse()

	// Strictly ensure valid inputs, awkwardly multiplying dereferences
	if extras := len(flag.Args()); extras > 0 || (*grid**circ**pts**its == 0) {
		fmt.Println("Usage: ./compute -grid=1000 -circ=900 -pts=100 -its=10")
		log.Fatal("Invalid arguments were provided. See ./compute -h for more.")
	}

	rand.Seed(time.Now().UTC().UnixNano())

	result := EstimateπByMonteCarlo(*grid, *circ, *pts, *its)

	output, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(output))
}

// EstimateπByMonteCarlo calculates iterations of MonteCarlo and averages them in parallel.
func EstimateπByMonteCarlo(gridSize, circleDiameter, points, iterations int) (m model) {
	var wg sync.WaitGroup
	results := make(chan float64)
	img := image.NewNRGBA(image.Rect(0, 0, gridSize, gridSize))

	go drawCircle(gridSize, circleDiameter, img)

	go func() {
		for result := range results {
			m.PI += result
		}
	}()

	// Add iterations workers to calculate on the same image
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			results <- MonteCarlo(gridSize, circleDiameter, points, img)
			return
		}()
	}

	// Wait for all threads to complete
	wg.Wait()
	close(results)

	// Encode the image as png
	bufPNG := new(bytes.Buffer)
	bufB64 := new(bytes.Buffer)
	if err := png.Encode(bufPNG, img); err != nil {
		log.Fatalf("PNG Encoding error: %s", err.Error())
	}

	// Encode PNG as B64
	enc := base64.NewEncoder(base64.StdEncoding, bufB64)
	if _, err := enc.Write(bufPNG.Bytes()); err != nil {
		log.Fatalf("B64 Encoding error: %s", err.Error())
	}
	enc.Close()

	// Return the result model
	m.PI = m.PI / float64(iterations)
	m.PNG = bufB64.String()
	return m
}

// MonteCarlo calculates a MonteCarlo simulation of π.
func MonteCarlo(gridSize, circleDiameter, points int, img *image.NRGBA) (result float64) {
	center := gridSize / 2
	radius := circleDiameter / 2
	probabilityIn := 0

	// Draw random points, calculate percent of points inside circle
	for i := 0; i < points; i++ {
		x := rand.Intn(gridSize)
		y := rand.Intn(gridSize)

		xoff := float64(x - center)
		yoff := float64(y - center)
		if xoff < 0 {
			xoff = -1 * xoff
		}
		if yoff < 0 {
			yoff = -1 * yoff
		}

		in := math.Sqrt(math.Pow(xoff, 2)+math.Pow(yoff, 2)) < float64(radius)
		if in {
			img.Set(x, y, color.NRGBA{192, 0, 0, 255})
			probabilityIn++
		} else {
			img.Set(x, y, color.NRGBA{64, 64, 192, 255})
		}
	}

	// Calculate estimate
	πAprox := float64(4*probabilityIn) / float64(points)
	return πAprox
}

// drawCircle will render the circle edge, boldly assuming 2π ≈ 6.283
func drawCircle(gridSize, circleDiameter int, img *image.NRGBA) {
	π2 := 6.283
	center := gridSize / 2
	radius64 := float64(circleDiameter) / 2
	circleResolution := float64(1000)

	for i := float64(0); i < circleResolution; i++ {
		rads := i / circleResolution * π2
		x := int(math.Floor(math.Sin(rads)*radius64)) + center
		y := int(math.Floor(math.Cos(rads)*radius64)) + center
		img.Set(x, y, color.RGBA{0, 0, 0, 255})
	}
}
