package main

import (
	"2nd_task/collection"
	"fmt"
)

func main() {
	var list collection.List
	list.Add("Mike")
	list.Add("Bill")
	list.Add("Lolly")
	list.Add("Kate")
	list.Add("Helen")

	list.Print()
	fmt.Print("\n")
	list.Remove(3)
	fmt.Print("\nAfter remove: ")
	list.Print()
	fmt.Print("\nGet(2): ",list.Get(2))
	fmt.Print("\nFirst: ",list.First())
	fmt.Print("\nLast: ",list.Last())
	fmt.Print("\nLength: ",list.Length())
}