package serverStat

import (
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

func (s *ServerStat) GetCurrentServerStat() *Usage {
	disk := make(chan DiskInformation)
	go DiskUsage(disk)
	usage := CpuAndMemoryUsage()
	usage.Disk = <-disk
	usage.ResponseTime = time.Now().UnixMilli()
	return usage
}
