package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	userHeight := 1.8
	userKilogram := 100.0
	BMI := userKilogram / math.Pow(userHeight, BMIPower)
	fmt.Print(BMI)

}
