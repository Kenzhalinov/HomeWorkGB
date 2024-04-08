package main

import (
	"fmt"
	"math/rand"
)

func main() {

	arr := [6]string{"one", "two", "three", "foo", "five", "six"}

	randomInt := rand.Intn(3) + 1

	newArr := make([]string, len(arr))
	copy(newArr[:], arr[:])

	for i := range newArr {
		j := rand.Intn(i + 1)
		newArr[i], newArr[j] = newArr[j], newArr[i]
	}

	result := make([]string, 0, randomInt)

	for i := 0; i < randomInt && i < len(newArr); i++ {

		result = append(result, newArr[i])
	}
	fmt.Println(result)
}
