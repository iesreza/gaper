package main

import (
	"os"

	"github.com/maxcnunes/gaper"
)

// build info
var (
	// keep the version hardcoded because on installing it through "go get/install"
	// it doesn't apply the build tags to override it. So, it is make easier for
	// people using in that case to find out which version they are using
	version = "1.0.3-dev"
)

func main() {

	cfg := gaper.Config{
		PollInterval: gaper.DefaultPoolInterval,
	}
	chOSSiginal := make(chan os.Signal, 2)
	if err := gaper.Run(&cfg, chOSSiginal); err != nil {
		panic(err)
	}

}
