package main

import (
	"log"

	"github.com/cmepw/221b/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
