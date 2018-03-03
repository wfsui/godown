package main

import (
	"log"
	"os"

	"github.com/mattn/godown"
)

func main() {
	if err := godown.Convert(os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
}