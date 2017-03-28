package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(100)
	fmt.Println("Test rand int", rand.Intn(10))
}
