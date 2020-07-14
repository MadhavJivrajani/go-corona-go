package utils

import (
	"fmt"
	"sort"
)

// Beautify prints a lsit of strings alphabetically and numerically indexed.
func Beautify(items []string) {
	sort.Strings(items)

	for index, item := range items {
		if index+1 < 10 {
			fmt.Printf("%d.  %s\n", index+1, item)
		} else {
			fmt.Printf("%d. %s\n", index+1, item)
		}
	}
}
