package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/jmarren/shazam/proc"
)

const usage string = `
	--- shazam ---`

func main() {

	switch os.Args[1] {
	case "cmdlines":
		cmdlines := proc.ListFromProcsW(proc.W_CMDL)
		for _, cmdline := range cmdlines {
			fmt.Println(cmdline)
		}
	case "pids":
		cmdlines := proc.ListFromProcsW(proc.W_PID)
		for _, cmdline := range cmdlines {
			fmt.Println(cmdline)
		}
	case "pid-cmdlines":
		cmdlines := proc.ListFromProcsW(proc.W_PID | proc.W_CMDL)
		for _, cmdline := range cmdlines {
			fmt.Println(cmdline)
		}
	case "pid-mustcmdlines":
		cmdlines := proc.ListFromProcsW(proc.W_PID | proc.W_CMDL | proc.MUST_CMDL)
		for _, cmdline := range cmdlines {
			fmt.Println(cmdline)
		}
	case "kill-cmdline-includes":
		proc.KillAll_CmdlineIncludes(os.Args[2])
	case "print-this-tgid":
		fmt.Println(proc.New(os.Getpid()).Tgid())
	case "print-this-ppid":
		fmt.Println(proc.New(os.Getpid()).Ppid())
	case "print-lineage":
		lineage := proc.ThisProcLineage
		slices.Reverse(lineage)
		for _, proc := range lineage {
			fmt.Printf("[ %d %s ] -> ", proc.Tgid(), proc.Name())
		}
		fmt.Println()

	}

}
