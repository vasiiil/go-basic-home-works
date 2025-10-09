package main

import "fmt"

func main() {
	const usdToEur float64 = 1.4
	const usdToRub float64 = 94

	eur := 10.0
	usd := eur * usdToEur
	rub := usd * usdToRub
	fmt.Printf("%.2f", rub)
}