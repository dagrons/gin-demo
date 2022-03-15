package cm

import (
	"fmt"
)

func PP(a [][]byte) {
	for _, s := range a {
		for _, c := range s {
			fmt.Print(string(c), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
