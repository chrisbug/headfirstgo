// average calculates the varge of sevral numbers
package main

import (
	"fmt"
	"log"

	"github.com/chrisbug/headfirstgo/datafile"
)

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3]
	fmt.Println(slice)
}

func main2() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average:%0.2f\n", sum/sampleCount)
}
