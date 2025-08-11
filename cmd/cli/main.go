package main

import (
	"fmt"

	"github.com/jmarren/shazam/proc"
)

const usage string = `
	--- shazam ---`

func main() {
	cmdlines := proc.ListCmdlines()
	for _, cmdline := range cmdlines {
		fmt.Println(cmdline)
	}

	// fmt.Println(usage)

}
