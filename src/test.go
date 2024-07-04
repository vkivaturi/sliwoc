package main

import "fmt"
import "sliwoc/core/cachewrite"

func main() {
	fmt.Println("Getting started with a light weight" +
		"cache that is short lived and write optimized")

	cachewrite.AddEntity()
	
}
