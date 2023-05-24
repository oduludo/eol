package main

import (
	"log"
	"oduludo.io/eol/eol"
	"os"
)

func main() {
	if err := eol.NewRootCmd(os.Stdout).Execute(); err != nil {
		log.Fatal(err)
	}
}
