package main

import (
	"flag"
	"fmt"
)

func main() {
	c := flag.String("c", "No key set", "Your pocket consumer key")
	flag.Parse()

	fmt.Println(*c)
}
