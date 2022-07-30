package serverStat

import (
	"time"

	"github.com/naim6246/server-stat-analyzer/configs"
)

var stats []*Usage

func (s *ServerStat) StoreStats() {
	config := configs.GetAppConfig()
	for {
		if len(stats) < config.MaxAllowedToStore {
			stats = append(stats, s.GetCurrentServerStat())
		} else {
			stats = append(stats[1:], s.GetCurrentServerStat())
		}
		time.Sleep(time.Second * time.Duration(config.IntervalInSec))
	}
}

func (s *ServerStat) GetLoggedServerStat() []*Usage {
	return stats
}
