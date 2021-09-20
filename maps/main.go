package main

import (
	"fmt"
)

func main() {
	// var ranks map[string]int
	// ranks = make(map[string]int)
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	fmt.Println(ranks)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
}
