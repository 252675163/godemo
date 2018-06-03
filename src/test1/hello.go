package main

import (
	"fmt"
	"stringutil"
	"math"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH")+"\n")
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}