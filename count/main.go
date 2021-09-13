// count tallies the number of times each line
// occures within a file
package main

import (
	"fmt"
	"log"

	"github.com/chrisbug/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
