package serverStat

import (
	"sync"
	"time"
)

type ServerStat struct{}

func NewServerStat() *ServerStat {
	return &ServerStat{}
}

type Usage struct {
	Cpu          float64         `json:"cpu"`
	Memory       float64         `json:"memory"`
	Disk         DiskInformation `json:"disk"`
	ResponseTime int64           `json:"responseTime"`
}

var wait sync.WaitGroup

func (s *ServerStat) ServerStat() *Usage {
	disk := make(chan DiskInformation)
	wait.Add(1)
	go DiskUsage(disk)
	wait.Add(1)
	usage := CpuAndMemoryUsage()
	wait.Wait()
	usage.Disk = <-disk
	usage.ResponseTime = time.Now().Unix()
	return usage
}
