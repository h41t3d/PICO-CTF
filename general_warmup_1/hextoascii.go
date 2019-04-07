package main

import (
    "encoding/hex"
    "fmt"
)

func main() {
    a := "41"
    bs,err := hex.DecodeString(a)
	if err != nil {
        panic(err)
    }
    fmt.Println(string(bs))
}