package main

import (
	"fmt"
	// "math"
	"strconv"
)

func main() {
	hmm := "3D"
	n,_ := strconv.ParseUint(hmm, 16, 32)
	fmt.Println(n)
}