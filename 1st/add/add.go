package add

import (
	"fmt"
)

func PrintRow(m []int) {
	for _, arr := range m {
		fmt.Printf("%6d", arr)
	}
	fmt.Println("")
}
