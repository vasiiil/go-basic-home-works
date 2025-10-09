package main

import "fmt"

func main() {
	const usdToEur float64 = 1.4
	const usdToRub float64 = 94
	const eurToRub = usdToRub * usdToEur

	eur := 10.0
	rub := eur * eurToRub
	fmt.Printf("%.2f", rub)
}