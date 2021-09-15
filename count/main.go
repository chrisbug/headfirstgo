// count tallies the number of times each line
// occures within a file
package main

import (
	"fmt"
	"log"

	"github.com/chrisbug/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("/home/buggch01/mycode/headfirstgo/count/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
