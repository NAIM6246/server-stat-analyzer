package serverStat

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type DiskInformation struct {
	TotalSize       float64
	TotalUsed       float64
	TotalAvailable  float64
	UsagePercentage float64
}

func DiskUsage(str chan DiskInformation) {
	cmd := exec.Command("df", "-h", "/")
	cmd.Stderr = os.Stdout
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:  ", err)
	}
	totalSize := 0.0
	totalUsed := 0.0
	totalAvailable := 0.0
	totalUsagePercentage := 0.0
	// diskOuput := ""
	for {
		line, err := cmdOutput.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]float64, 0)

		if !strings.Contains(tokens[0], "/dev/sda") {
			continue
		}

		for _, t := range tokens {
			if t != "" && t != "\t" && !strings.Contains(t, "/dev/") && (t[0] >= '0' && t[0] <= '9') {
				str := ""
				for _, s := range t {
					if s >= '0' && s <= '9' || s == '.' {
						str += string(s)
					}
				}
				data, err := strconv.ParseFloat(str, 64)
				if err != nil {
					log.Fatal(err)
				}
				ft = append(ft, data)
			}
		}
		size := ft[0]
		used := ft[1]
		available := ft[2]
		usagePercentage := ft[3]

		totalSize += size
		totalUsed += used
		totalAvailable += available
		totalUsagePercentage += usagePercentage
	}

	wait.Done()
	str <- DiskInformation{
		TotalSize:       totalSize,
		TotalUsed:       totalUsed,
		TotalAvailable:  totalAvailable,
		UsagePercentage: totalUsagePercentage,
	}
}
