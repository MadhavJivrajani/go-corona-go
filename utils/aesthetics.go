package utils

import (
	"fmt"
	"sort"
)

// BeautifyIndexed prints a lsit of strings alphabetically and numerically indexed.
func BeautifyIndexed(items []string) {
	sort.Strings(items)

	for index, item := range items {
		if index+1 < 10 {
			fmt.Printf("%d.  %s\n", index+1, item)
		} else {
			fmt.Printf("%d. %s\n", index+1, item)
		}
	}
}
