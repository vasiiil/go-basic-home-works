package main

import "fmt"

func main() {

	eur := getUserInput()
	rub := convert(eur, "eur", "rub")
	fmt.Printf("%.2f", rub)
}

func getUserInput() float64 {
	var value float64
	fmt.Print("Введите число: ")
	fmt.Scan(&value)

	return value
}

func convert(value float64, source string, target string) float64 {
	const usdToEur float64 = 1.4
	const usdToRub float64 = 94
	const eurToRub = usdToRub * usdToEur
	return value * eurToRub
}
