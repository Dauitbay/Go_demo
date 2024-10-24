package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower = 2
	userHeight := 1.8
	userKilogram := 100.0
	fmt.Println("__ Body Index Mass Calculator __ ")
	fmt.Print("Enter your height in meter: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kg: ")
	fmt.Scan(&userKilogram)
	BMI := userKilogram / math.Pow(userHeight, BMIPower)
	fmt.Printf("Your BMI is: %.0f \n", BMI)
	fmt.Print(BMI)

}
