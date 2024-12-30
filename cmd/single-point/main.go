package main

import (
	"fmt"
	"gorf/pkg/propagation"
	"log"
	"math"
	"os"
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

func main() {
	foo1()
}
