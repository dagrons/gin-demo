package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x%11 == 0 {
			fmt.Println("yes")
			continue
		}
		cntOne := 0
		for x != 0 {
			if x%10 == 1 {
				cntOne++
			}
			x /= 10
		}
		if cntOne >= 2 {
			fmt.Println("yes")
		} else {
			fmt.Printf("no")
		}
	}
}
