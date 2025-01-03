package main

import (
	"fmt"
	"gorf/pkg/propagation"
	"log"
	"math"
	"os"

	"github.com/go-gl/mathgl/mgl64"
)

func foo1() {
	f, err := os.Create("dBi.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var foo propagation.Pattern = propagation.SincPattern{Factor: 0}

	elevations := make([]float64, 360/5+1)
	azimuths := make([]float64, 360/5+1)

	for i := 0; i < len(elevations); i++ {
		elevations[i] = 5 * float64(i) * (math.Pi / 180)
		azimuths[i] = 5 * float64(i) * (math.Pi / 180)
	}

	f.WriteString("azimuth, elevation, dBi\n")
	for i := 0; i < len(elevations); i++ {
		for j := 0; j < len(azimuths); j++ {
			az, el := azimuths[j], elevations[i]
			dBi := foo.CalcGainAE(az, el)
			f.WriteString(fmt.Sprintf("%v, %v, %v\n", az*180/math.Pi, el*180/math.Pi, dBi))
		}
	}
}

func foo2() {
	testLoc := mgl64.Vec3{1, 0, 0}

	sigLoc := mgl64.Vec3{0, 0, 0}
	sigDir := mgl64.Vec3{1, 0, 0}
	sig := propagation.Signal{
		Pattern:   propagation.SincPattern{Factor: 1},
		Watts:     100,
		Frequency: 2_800_000_000,
		Direction: sigDir,
		Location:  sigLoc,
	}

	fmt.Printf("--- %v ---\n", propagation.PowerAtPosition(sig, testLoc))
}

func main() {
	foo2()
}
