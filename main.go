package main

import (
	"fmt"
)

func main() {
	var inputuser int
	fmt.Println("masukkan 1/2: ")
	fmt.Scanln(&inputuser)

	if inputuser == 1 {
		var rows int = 5
		var k int

		for i := 1; i <= rows; i++ {
			k = 0
			for space := 1; space <= rows-i; space++ {
				fmt.Print("  ")
			}
			for {
				fmt.Print("* ")
				k++
				if k == 2*i-1 {
					break
				}
			}
			fmt.Println("")
		}
	} else {
		var rows2 int = 5
		var a, b int
		for a = rows2; a >= 1; a-- {
			for space := 1; space <= rows2-a; space++ {
				fmt.Print("  ")
			}
			for b = a; b <= 2*a-1; b++ {
				fmt.Printf("* ")
			}
			for b = 0; b < a-1; b++ {
				fmt.Printf("* ")
			}
			fmt.Println("")
		}
	}
}
