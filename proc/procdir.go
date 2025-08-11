package proc

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func listProcEntries() []os.DirEntry {
	entries, err := os.ReadDir("/proc")

	if err != nil {
		fmt.Printf("error reading /proc: %s\n", err)
		panic(err)
	}
	return entries
}

func listPids() []int {
	var pids []int

	for _, entry := range listProcEntries() {
		pid, err := strconv.Atoi(entry.Name())
		if err == nil {
			pids = append(pids, pid)
		}
	}

	return pids
}

func listProcs() []*Proc {
	var procs []*Proc

	for _, pid := range listPids() {
		proc := New(pid)
		procs = append(procs, proc)
	}
	return procs
}

func ListFromProcs(listFunc func(proc *Proc, list []string) []string) []string {
	var list []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, proc := range listProcs() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			list = listFunc(proc, list)
			mu.Unlock()

		}()
	}

	wg.Wait()

	return list
}

func ListFromProcsW(flags int) []string {
	var list []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, proc := range listProcs() {
		wg.Add(1)

		go func() {
			info := proc.Info(flags)
			if len(info) > 0 {
				mu.Lock()
				list = append(list, info)
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return list
}

func KillAll_CmdlineIncludes(s string) {
	for _, proc := range listProcs() {
		proc.Kill_CmdLineIncludes(s)
	}
}

// var cmdlines []string
// var mu sync.Mutex
// var wg sync.WaitGroup
//
// for _, proc := range listProcs() {
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		cmdline, err := proc.Cmdline()
// 		if err == nil && len(cmdline) > 0 {
// 			mu.Lock()
// 			cmdlines = append(cmdlines, proc.pidStr()+"\t"+cmdline)
// 			mu.Unlock()
// 		}
// 	}()
// }
//
// wg.Wait()
//
// return cmdlines
