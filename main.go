package main

import (
	"log"
	"os"

	"github.com/mlin-aviatrix/avx/cmd"
)

func main() {
	f, _ := os.Open("/dev/null")
	log.SetOutput(f)
	cmd.Execute()
}
