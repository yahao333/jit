package main

import (
	"log"

	"github.com/yahao333/jit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
