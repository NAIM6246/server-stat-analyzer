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

func CpuAndMemoryUsage() *Usage {
	cmd := exec.Command("ps", "aux")
	cmd.Stderr = os.Stdout
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:  ", err)
	}

	totalCpuUsage := 0.0
	totalMemUsage := 0.0
	for {
		line, err := cmdOutput.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "USER" {
			continue
		}
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}

		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		mem, err := strconv.ParseFloat(ft[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		totalCpuUsage += cpu
		totalMemUsage += mem

	}
	return &Usage{
		Cpu:    totalCpuUsage,
		Memory: totalMemUsage,
	}
}
