package nmap

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Ping struct {
	Name    string
	IP      string
	Latency string
}

func PingScan(hosts string) []Ping {
	var (
		allPings []Ping
		cmd      *exec.Cmd
		out      bytes.Buffer
	)
	cmd = exec.Command("/usr/local/bin/nmap", "-sP", hosts)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return allPings
	}
	splitSection := strings.Split(out.String(), "\n")
	var fullString []string
	fullString = make([]string, 0)
	currentString := ""
	secondLine := false
	for i := 0; i < len(splitSection); i++ {
		if !secondLine {
			currentString += (splitSection[i] + " ")
			secondLine = true
		} else {
			currentString += splitSection[i]
			fullString = append(fullString, currentString)
			currentString = ""
			secondLine = false
		}
	}

	fullString = append(fullString[1 : len(fullString)-1])
	allPings = make([]Ping, 0)
	for i := 0; i < len(fullString); i++ {
		var ping Ping
		fullString[i] = strings.Replace(fullString[i], "Nmap scan report for ", "", -1)
		fullString[i] = strings.Replace(fullString[i], "Host is up (", "", -1)
		ping.Name = strings.Split(fullString[i], " ")[0]
		ping.IP = strings.Split(fullString[i], " ")[1]
		ping.IP = strings.Replace(ping.IP, "(", "", -1)
		ping.IP = strings.Replace(ping.IP, ")", "", -1)
		ping.Latency = strings.Split(fullString[i], " ")[2]
		allPings = append(allPings, ping)
	}

	return allPings
}
