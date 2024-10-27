package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Body Index Mass Calculator __ ")
	userKilogram, userHeight := getUserInput()
	BMI := calculateBMI(userKilogram, userHeight)
	outputResult(BMI)
	isLean := BMI < 16
	if isLean {
		fmt.Println("You have strong deficit of the weight")
	} else if BMI < 18.5 {
		fmt.Println("You have deficit of the weight")
	} else if BMI < 25 {
		fmt.Println("You have normal weight")
	} else if BMI < 30 {
		fmt.Println("You are overweight")
	} else {
		fmt.Println("You have a degree of obesity")
	}
}


func outputResult(bmi float64) {
	result := fmt.Sprintf("Your BMI is: %.0f \n", bmi)
	fmt.Println(result)
}

func calculateBMI(userKilogram float64, userHeight float64) float64 {
	const BMIPower = 2
	BMI := userKilogram / math.Pow(userHeight/100, BMIPower)
	return BMI
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKilogram float64
	fmt.Print("Enter your height in meter: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight in kg: ")
	fmt.Scan(&userKilogram)
	return userKilogram, userHeight
}
