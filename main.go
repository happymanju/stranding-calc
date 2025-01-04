package main

import (
	"os"

	"github.com/happymanju/strandingcalc/strandingcalc"
)

func main() {
	os.Exit(strandingcalc.Run(os.Args[1:]))
}
