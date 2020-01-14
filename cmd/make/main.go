package main

import (
	"fmt"
	"os"

	"github.com/xyproto/makeflags"
)

func main() {
	config := makeflags.New()

	if config.VersionInfoAndExit {
		fmt.Println(makeflags.Version)
		os.Exit(0)
	}

	if len(config.Targets) == 0 && config.Makefile == "" {
		fmt.Println("make: *** No targets specified and no makefile found.  Stop.")
		os.Exit(2)
	}

	fmt.Printf("[%s] %v\n", config.Makefile, config.Targets)
}
