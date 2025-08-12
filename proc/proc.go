package proc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var ThisProcLineage []*Proc

func init() {
	ThisProcLineage = New(os.Getpid()).Lineage()
}

type Proc struct {
	pid        int
	statusFile *StatusFile
}

func New(pid int) *Proc {
	return &Proc{
		pid: pid,
	}
}

func (p *Proc) pidStr() string {
	return strconv.Itoa(p.pid)
}

func (p *Proc) Cmdline() (string, error) {
	cmdline, err := os.ReadFile(p.cmdlinePath())
	return string(cmdline), err
}

func (p *Proc) Process() *os.Process {
	process, err := os.FindProcess(p.pid)
	if err != nil {
		fmt.Printf("error finding process with pid: %s\n", err)
		panic(err)
	}
	return process
}

func (p *Proc) Kill_CmdLineIncludes(s string) {
	cmdline, err := p.Cmdline()
	if err != nil {
		return
	}
	if strings.Contains(cmdline, s) {
		if !IsAncestorOfThisProc(p.pid) {
			fmt.Println("killing: " + p.pidStr())
			p.Process().Kill()
		}
	}
}

const W_PID = 0b01
const W_CMDL = 0b10
const MUST_CMDL = 0b100

func (p *Proc) Info(flags int) string {
	var infoStr string
	var wg sync.WaitGroup

	var items [3]string

	if flags&W_PID > 0 {
		items[0] = p.pidStr()
	}

	if flags&W_CMDL > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cmdline, err := p.Cmdline()
			if err == nil && len(cmdline) > 0 {
				items[1] = cmdline
			}
		}()
	}

	wg.Wait()

	if flags&MUST_CMDL > 0 && len(items[1]) < 1 {
		return ""
	}

	for _, item := range items {
		if len(item) > 0 {
			infoStr += item + "\t"
		}
	}

	return infoStr
}

func (p *Proc) statusFileContents() (string, error) {
	statusFile, err := os.ReadFile(p.statusPath())
	return string(statusFile), err
}

func (p *Proc) StatusFile() *StatusFile {
	if p.statusFile != nil {
		return p.statusFile
	}
	sfContents, err := p.statusFileContents()
	if err != nil {
		panic(err)
	}

	sf := NewStatusFile(sfContents)
	p.statusFile = sf
	return sf
}

func (p *Proc) Name() string {
	return p.StatusFile().name
}

func (p *Proc) Tgid() int {
	return p.StatusFile().threadGroupId
}

func (p *Proc) Ppid() int {
	return p.StatusFile().ppid
}

func (p *Proc) Parent() *Proc {
	return New(p.Ppid())
}

func (p *Proc) Lineage() []*Proc {
	var lineage []*Proc
	curr := p

	for curr.pid != 1 {
		lineage = append(lineage, curr)
		curr = curr.Parent()
	}

	lineage = append(lineage, curr)
	return lineage
}

func IsAncestorOfThisProc(pid int) bool {
	for _, ancestor := range ThisProcLineage {
		if ancestor.Tgid() == pid {
			return true
		}
	}
	return false
}
