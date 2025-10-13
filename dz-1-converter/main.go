package main

import (
	"errors"
	"fmt"
	"strings"
)

type tCourse = map[string]float64

const EMPTY_CURRENCY_ERROR_CODE = "EMPTY_VALUE"
const UNKNOWN_CURRENCY_ERROR_CODE = "UNKNOWN_VALUE"
const EXCLUDE_CURRENCY_ERROR_CODE = "EXCLUDE_VALUE"

func main() {
	sourceCurrecyCode := getCurrency(false, "")
	value := getUserInput()
	targerCurrecyCode := getCurrency(true, sourceCurrecyCode)
	convertedValue, _ := convert(value, sourceCurrecyCode, targerCurrecyCode)
	fmt.Printf("%.2f", convertedValue)
}

func getUserInput() float64 {
	var value float64
	fmt.Print("Введите число: ")
	for {
		fmt.Scan(&value)
		if value <= 0 {
			fmt.Println("Введите нормальное число")
			continue
		}
		break
	}

	return value
}

func getCurrency(isTarget bool, excludeCurrencyCode string) string {
	currencies := make([]string, 0, 3)
	if excludeCurrencyCode != "u" {
		currencies = append(currencies, "USD (U)")
	}
	if excludeCurrencyCode != "e" {
		currencies = append(currencies, "EUR (E)")
	}
	if excludeCurrencyCode != "r" {
		currencies = append(currencies, "RUB (R)")
	}

	currencyType := "исходную"
	if isTarget {
		currencyType = "конечную"
	}
	fmt.Println("Выберите", currencyType, "валюту:", strings.Join(currencies, ", "))

	for {
		currencyCode, err := getUserCurrency(excludeCurrencyCode)
		if err != nil {
			errorCode := err.Error()
			if errorCode == EMPTY_CURRENCY_ERROR_CODE {
				fmt.Println("Необходимо выбрать валюту")
				continue
			}
			if errorCode == UNKNOWN_CURRENCY_ERROR_CODE {
				fmt.Println("Такой валюты нет")
				continue
			}
			if errorCode == EXCLUDE_CURRENCY_ERROR_CODE {
				fmt.Println("Нельзя выбирать ту же валюту")
				continue
			}
		}
		return currencyCode
	}
}

func getUserCurrency(excludeCurrencyCode string) (string, error) {
	var currencyCode string
	fmt.Scan(&currencyCode)
	if currencyCode == "" {
		return "", errors.New(EMPTY_CURRENCY_ERROR_CODE)
	}
	currencyCode = strings.ToLower(currencyCode)
	if currencyCode != "u" && currencyCode != "e" && currencyCode != "r" {
		return "", errors.New(UNKNOWN_CURRENCY_ERROR_CODE)
	}
	if currencyCode == excludeCurrencyCode {
		return "", errors.New(EXCLUDE_CURRENCY_ERROR_CODE)
	}
	return currencyCode, nil
}

func convert(value float64, source string, target string) (float64, error) {
	const eurToUsd float64 = 1.4
	const usdToRub float64 = 94
	const eurToRub = eurToUsd * usdToRub
	courses := map[string]tCourse{
		"e": {
			"u": eurToUsd,
			"r": eurToRub,
		},
		"r": {
			"e": 1 / eurToRub,
			"u": 1 / usdToRub,
		},
		"u": {
			"e": 1 / eurToUsd,
			"r": usdToRub,
		},
	}
	sourceM, ok := courses[source]
	if !ok {
		return 0, errors.New("UNKNOWN_CURRENCIES")
	}
	course, ok := sourceM[target]
	if !ok {
		return 0, errors.New("UNKNOWN_CURRENCIES")
	}
	return value * course, nil
}
