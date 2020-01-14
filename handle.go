package makeflags

import (
	"fmt"
	"os"
)

// Handle will look at the config and either output version
// info (if -v or --version are given) or output an error if
// no make targets are given.
func (config *Config) Handle() {
	if config.VersionInfoAndExit {
		fmt.Println(Version)
		os.Exit(0)
	}
	if len(config.Targets) == 0 {
		fmt.Println("make: *** No targets specified and no makefile found.  Stop.")
		os.Exit(2)
	}
}
