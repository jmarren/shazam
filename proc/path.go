package proc

func (p *Proc) path() string {
	return "/proc/" + p.pidStr()
}

func (p *Proc) cmdlinePath() string {
	return p.path() + "/cmdline"
}

func (p *Proc) exePath() string {
	return p.path() + "/exe"
}

func (p *Proc) statusPath() string {
	return p.path() + "/status"
}
