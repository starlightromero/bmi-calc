package info

import "fmt"

const (
	mainTitle    = "BMI Calculator"
	separator    = "------------------"
	WeightPrompt = "Please enter your weight (lbs): "
	HeightPrompt = "Please enter your height (in): "
)

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
