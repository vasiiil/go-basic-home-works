package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var exit string
	for {
		operation := getOperation()
		if operation == "n" {
			return
		}
		numbers := getNumbers()
		if len(numbers) == 0 {
			return
		}

		result, _ := calculate(operation, numbers)
		fmt.Printf("Operation: %v, result: %.2f\n", operation, result)

		fmt.Println("Хотите продолжить? (y/n)")
		fmt.Scan(&exit)
		if exit == "n" || exit == "N" {
			return
		}
	}
}

func getOperation() string {
	fmt.Println("Выберите операцию AVG (a), SUM (s), MEDIAN (m)")
	fmt.Println("Для выхода введите n")
	var operation string
	allowedOperations := map[string]bool{
		"a": true,
		"s": true,
		"m": true,
		"n": true,
	}

	for {
		fmt.Scan(&operation)
		_, isAllowed := allowedOperations[operation]
		if !isAllowed {
			fmt.Println("Такой операции нет")
			continue
		}
		return operation
	}
}

func getNumbers() []int {
	fmt.Println("Введите целые числа через запятую")
	fmt.Println("Для выхода введите n")
	var s string
	fmt.Scan(&s)
	if s == "n" {
		return make([]int, 0)
	}

	stringSlice := strings.Split(s, ",")
	result := make([]int, 0, len(stringSlice))
	for _, s := range stringSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error converting '%s' to integer: %v\n", s, err)
			// Handle the error appropriately, e.g., skip, return an error, or use a default value
			continue
		}
		result = append(result, num)
	}

	return result
}

func calculate(operation string, numbers []int) (float64, error) {
	sum := 0.0
	for _, value := range numbers {
		sum += float64(value)
	}

	operations := map[string]func()float64{
		"a": func() float64 {
			return sum / float64(len(numbers))
		},
		"s": func() float64 {
			return sum
		},
		"m": func() float64 {
			return median(numbers)
		},
	}
	f := operations[operation]
	if f != nil {
		return f(), nil
	}

	return 0.0, errors.New("INVALID_OPERATION")
}

func median(data []int) float64 {
	// Create a copy of the data to avoid modifying the original slice.
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)

	// Sort the copied data.
	sort.Ints(dataCopy)

	l := len(dataCopy)

	// Handle empty slice case.
	if l == 0 {
		return 0
	}

	// Calculate median based on whether the length is even or odd.
	if l%2 == 0 {
		// Even length: average of the two middle elements.
		return float64(dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		// Odd length: the middle element.
		return float64(dataCopy[l/2])
	}
}
