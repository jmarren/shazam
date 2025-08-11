package proc

import (
	"os"
	"strconv"
)

type Proc struct {
	pid int
}

func New(pid int) *Proc {
	return &Proc{
		pid: pid,
	}
}

func (p *Proc) pidStr() string {
	return strconv.Itoa(p.pid)

}

func (p *Proc) path() string {
	return "/proc/" + p.pidStr()
}

func (p *Proc) cmdlinePath() string {
	return p.path() + "/cmdline"
}

func (p *Proc) exePath() string {
	return p.path() + "/exe"
}

func (p *Proc) Cmdline() (string, error) {
	cmdline, err := os.ReadFile(p.cmdlinePath())

	return string(cmdline), err
}
