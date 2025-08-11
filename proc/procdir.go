package proc

import (
	"os"
	"strconv"
	"sync"
)

// type ProcDir struct {
// }

func listProcEntries() []os.DirEntry {
	entries, err := os.ReadDir("/proc")

	if err != nil {
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

func ListCmdlines() []string {
	var cmdlines []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, proc := range listProcs() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cmdline, err := proc.Cmdline()
			if err == nil {
				mu.Lock()
				cmdlines = append(cmdlines, cmdline)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	return cmdlines

}
