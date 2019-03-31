package  main

// alint reports issues in ansible inventory hosts files

import (
	"flag"
	"fmt"
	"github.com/davecb/inventoryTree/pkg/alint"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "add verbose messages")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("You must supply an inventory directory containing hosts files\n")
		fmt.Printf("usage: alint [-d] dir ...\n")

	}
	for i:=0; i < flag.NArg(); i++ {
		alint.LintHostFiles(flag.Arg(i), verbose)
	}
}
