package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2}
	output := robHouses(input)

	fmt.Println("Output:", output)

}

func robHouses(houses []int) int {

	if len(houses) == 0 {
		return 0
	}
	if len(houses) == 1 {
		return houses[0]
	}
	if len(houses) == 2 {
		return getMax(houses[0], houses[1])
	}

	tempArray := make([]int, len(houses))
	tempArray[0] = houses[0]
	tempArray[1] = getMax(houses[0], houses[1])

	for i := 2; i < len(houses); i++ {
		tempArray[i] = getMax(houses[i]+tempArray[i-2], tempArray[i-1])
	}

	return tempArray[len(houses)-1]
}

func getMax(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
}
