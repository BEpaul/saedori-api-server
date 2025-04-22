package main

import (
	"flag"

	cmd "github.com/bestkkii/saedori-api-server/cmd/app"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
