package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	p, _ := strconv.Atoi(os.Args[1])
	var r int
	for i := 0; i <= p; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			r += i
		}
	}
	fmt.Println(r)
}
