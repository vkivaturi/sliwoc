package main

import "fmt"

type Val struct {
	key  string
	data string
}

func addToCache(val Val, myMap2 map[string]Val) {
	myMap2[val.key] = val
}

func getFromCache(key string, myMap2 map[string]Val) string {
	return myMap2[key].data
}

func main() {
	fmt.Println("Getting started with a light weight" +
		"cache that is short lived and write optimized")

	cacheMap := make(map[string]Val)
	var value Val
	value.key = "key1"
	value.data = "my data1"
	addToCache(value, cacheMap)
	fmt.Println(getFromCache("key1", cacheMap))
}
