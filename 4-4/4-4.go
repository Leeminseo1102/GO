// 4-4
package main

import (
	"fmt"
)

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for value := range a {
		fmt.Println(value)
	}
}
