package main

import (
	"fmt"

	"github.com/xyproto/makeflags"
)

func main() {
	config := makeflags.New()
	config.Handle()

	fmt.Println(config)
}
