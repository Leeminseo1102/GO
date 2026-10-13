// 3-16
package main

import (
	"fmt"
)

func main() {
	i := 2
	switch i {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 3, 5, 7:
		fmt.Println("홀수")
	}
}
