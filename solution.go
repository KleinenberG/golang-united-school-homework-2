package main

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type myType int

const SidesSquare = 4
const SidesTriangle = 3
const SidesCircle = 0

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func main() {
	CalcSquare(10.0, SidesCircle)
}

func CalcSquare(sideLen float64, sidesNum myType) float64 {
	if sidesNum == SidesSquare {
		p := float64(sideLen) * float64(SidesSquare) / 2
		S := math.Sqrt(p * (p - sideLen) * (p - sideLen) * (p - sideLen))
		fmt.Println(S)
		return S
	} else if sidesNum == SidesTriangle {
		fmt.Println(math.Pow(sideLen, 2))
		return (math.Pow(sideLen, 2))
	} else if sidesNum == SidesCircle {
		r := sideLen / (2 * math.Pi)
		s := math.Pi * r * r
		fmt.Println(s)
		return s
	} else {
		return 0
	}
}
