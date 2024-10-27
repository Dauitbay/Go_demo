package main

import (
	"fmt"
	"math"
	"os/user"
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
	BMI := calculateBMI(userKilogram, userHeight)
	outputResult(BMI)
}


func outputResult(bmi float64) {
	result := fmt.Sprintf("Your BMI is: %.0f \n", bmi)
	fmt.Print(result)
}

func calculateBMI(userKilogram float64, userHeight float64) float64 {
	const BMIPower = 2
	BMI := userKilogram / math.Pow(userHeight/100, BMIPower)
	return BMI
}
