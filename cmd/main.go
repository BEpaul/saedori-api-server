package main

import (
	"flag"

	cmd "github.com/bestkkii/saedori-api-server/cmd/app"
)


func main() {
	flag.Parse()

	cmd.NewCmd()
}
