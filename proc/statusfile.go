package proc

import (
	"strconv"
	"strings"
)

type StatusFileFactory struct {
	lines []string
}

func newStatusBuilder(content string) *StatusFileFactory {
	return &StatusFileFactory{
		lines: strings.Split(content, "\n"),
	}
}

type StatusFile struct {
	// name of the command that started the process
	name string

	// file creation permissions; see man umask(2)
	umask int
	// State  Current state of the process.  One of "R (running)",
	// "S (sleeping)", "D (disk sleep)", "T (stopped)", "t
	// (tracing stop)", "Z (zombie)", or "X (dead)".
	state string

	// this is the real process id  returned from os.Getpid()
	threadGroupId int
	numaGroupId   int

	// thread id (returned from gettid())
	pid int

	// parent process id
	ppid int

	// pid of process tracing this process, 0 if none
	tracerPid int

	// number of threads in the process containing this thread
	threadCount int
}

func (s *StatusFileFactory) build() *StatusFile {
	sf := new(StatusFile)

	for _, line := range s.lines {
		if strings.HasPrefix(line, "Name:") {
			sf.name = line[6:]
		}
		if strings.HasPrefix(line, "Tgid:") {
			tgid, err := strconv.Atoi(line[6:])
			if err != nil {
				panic(err)
			}
			sf.threadGroupId = tgid
		}
		if strings.HasPrefix(line, "PPid:") {
			ppid, err := strconv.Atoi(line[6:])
			if err != nil {
				panic(err)
			}
			sf.ppid = ppid
		}
	}

	return sf
}

func NewStatusFile(content string) *StatusFile {
	factory := newStatusBuilder(content)
	return factory.build()
}
