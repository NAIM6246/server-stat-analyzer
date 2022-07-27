package serverStat

import "sync"

type ServerStat struct{}

func NewServerStat() *ServerStat {
	return &ServerStat{}
}

var wait sync.WaitGroup

func (s *ServerStat) ServerStat() {
	CpuAndMemoryUsage()
	DiskUsage()
}
