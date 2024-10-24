package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	userHeight := 1.8
	userKilogram := 100.0
	fmt.Print("__ Body Index Mass Calculator __ \n")
	fmt.Print("Enter your height in meter: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kg: ")
	fmt.Scan(&userKilogram)
	BMI := userKilogram / math.Pow(userHeight, BMIPower)
	fmt.Print(BMI)

}
