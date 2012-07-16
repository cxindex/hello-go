package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	buf, err := ps.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	strs := strings.Split(string(buf), "\n")
	for i := 0; i < len(strs)-1; i++ { //newline
		pid := strings.Fields(strs[i])
		lock := false
		total := 0
		for i := 0; i < len(strs)-1; i++ {
			ppid := strings.Fields(strs[i])
			if pid[0] == ppid[1] {
				if !lock {
					fmt.Printf("%s %s \\____ ", pid[0], pid[2])
					lock = true
				}
				fmt.Printf("\n\t[ %s %s ]", ppid[0], ppid[2])
				total++
			}
		}
		if lock {
			if total == 1 {
				fmt.Printf("\n\tTOTAL %v CHILD\n", total)
			} else {
				fmt.Printf("\n\tTOTAL %v CHILDREN\n", total)
			}
		}
	}
}
