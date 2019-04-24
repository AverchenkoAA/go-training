package main

import (
	"1st/add"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 10000
	arrayNum := make([]int, 11)
	var arrayPer []int
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < n; i++ {
		d1 := rand.Intn(6) + 1
		d2 := rand.Intn(6) + 1
		arrayNum[d1+d2-2]++
	}
	for _, arr := range arrayNum {
		arrayPer = append(arrayPer, arr*100/n)
	}
	for i := 0; i < 11; i++ {
		fmt.Printf("%6d", i+2)
	}
	fmt.Println("")
	add.PrintRow(arrayNum)
	add.PrintRow(arrayPer)

}
