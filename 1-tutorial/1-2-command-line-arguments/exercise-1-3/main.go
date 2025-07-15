package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	begin := time.Now()
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	duration := time.Since(begin).Nanoseconds()
	fmt.Printf("inefficient version took %v\n", duration)

	begin = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	duration = time.Since(begin).Nanoseconds()
	fmt.Printf("efficient version took %v\n", duration)
}
