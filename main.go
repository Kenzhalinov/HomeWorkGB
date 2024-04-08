package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println("Введите элементы массива через пробел:")
	_, _ = fmt.Scan(&arr)

	randomInt := rand.Intn(3) + 1
	// newArr := make([]string, len(arr)
	result := make([]string, 0, randomInt)

	for i := 0; i < randomInt && i < len(arr); i++ {
		result = append(result, arr[i])
	}
	fmt.Println(result)
}
