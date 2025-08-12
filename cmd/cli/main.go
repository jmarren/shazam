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
		lines := [][]string{
			{},
			{},
			{},
		}
		for _, proc := range lineage {
			// cmd, _ := proc.Cmdline()
			lines[0] = append(lines[0], fmt.Sprintf("%d", proc.Tgid()))
			lines[1] = append(lines[1], proc.Name())
			// lines[2] = append(lines[2], cmd)
			// fmt.Printf("[ %s %d %s ] -> ", cmd, proc.Tgid(), proc.Name())
		}

		fmt.Println("_______________________________________________________________________________________________________________________")

		for _, line := range lines {
			for _, item := range line {
				itemLen := len(item)
				diff := 20 - itemLen
				margin1 := repeat(" ", diff/2)
				margin2 := repeat(" ", diff/2)
				if diff%2 == 1 {
					margin2 += " "
				}

				fmt.Printf("| %s%s%s ", margin1, item, margin2)
			}
			fmt.Println()
		}

		fmt.Printf("%v\n", lines)

		fmt.Println()

	}

}

func repeat(s string, i int) string {
	str := ""
	for range i {
		str += s
	}
	return str
}
