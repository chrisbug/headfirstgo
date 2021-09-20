// count tallies the number of times each line
// occures within a file
package main

import (
	"fmt"
	"log"

	"github.com/chrisbug/headfirstgo/datafile"
)

func main() {
	// lines, err := datafile.GetStrings("/home/buggch01/mycode/headfirstgo/count/votes.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var names []string
	// var counts []int
	// for _, line := range lines {
	// 	matched := false
	// 	for i, name := range names {
	// 		if name == line {
	// 			counts[i]++
	// 			matched = true
	// 		}
	// 	}
	// 	if matched == false {
	// 		names = append(names, line)
	// 		counts = append(counts, 1)
	// 	}
	// }
	// for i, name := range names {
	// 	fmt.Printf("%s: %d\n", name, counts[i])
	// }
	var counts map[string]int = make(map[string]int)
	var winner string
	lines, err := datafile.GetStrings("/Users/chris/code/headfirstgo/count/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
		if count > counts[winner] {
			winner = name
		}
	}
	fmt.Printf("Your winner is %s", winner)
}
