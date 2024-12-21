package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func checkDistance(arr []int) string {
	if len(arr) < 2 {
		return "Data berjarak tidak tetap"
	}
	distance := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != distance {
			return "Data berjarak tidak tetap"
		}
	}
	return fmt.Sprintf("Data berjarak %d", distance)
}

func main() {
	var input int
	var numbers []int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		numbers = append(numbers, input)
	}

	insertionSort(numbers)

	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Println()

	result := checkDistance(numbers)
	fmt.Println(result)
}
