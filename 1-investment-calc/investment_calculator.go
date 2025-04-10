package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("Enter Years to Invest: ")
	fmt.Scan(&years)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//fmt.Printf("You will earn %.2f \n", futureValue)
	//fmt.Printf("You will earn %.2f with inflation \n", futureRealValue)

	fv, rfv := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", fv)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", rfv)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}
