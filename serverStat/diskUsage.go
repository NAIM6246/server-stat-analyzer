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

func DiskUsage() {
	cmd := exec.Command("pydf")
	cmd.Stderr = os.Stdout
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:  ", err)
	}
	totalSize := 0.0
	totalUsed := 0.0
	totalAvailable := 0.0
	for {
		line, err := cmdOutput.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		fmt.Println(ft[1])
		size, err := strconv.ParseFloat(ft[1], 64)
		if err != nil {
			continue
		}
		used, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		available, err := strconv.ParseFloat(ft[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		totalSize += size
		totalUsed += used
		totalAvailable += available
	}

	fmt.Println("total disk size : ", totalSize)
	fmt.Println("total disk usage : ", totalUsed)
	fmt.Println("total available : ", totalAvailable)
	wait.Done()
}
