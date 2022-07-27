package serverStat

import (
	"sync"
)

type ServerStat struct{}

func NewServerStat() *ServerStat {
	return &ServerStat{}
}

type Usage struct {
	Cpu float64
	Memory float64
	Disk string
}

var wait sync.WaitGroup

func (s *ServerStat) ServerStat() *Usage {
	disk := make(chan string)
	wait.Add(1)
	go DiskUsage(disk)
	wait.Add(1)
	usage :=  CpuAndMemoryUsage()
	wait.Wait()
	usage.Disk = <-disk
	return usage
}
