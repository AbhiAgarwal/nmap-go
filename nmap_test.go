package nmap

import (
	"fmt"
	"testing"
)

func TestPingScan(t *testing.T) {
	allPings := PingScan("192.168.1.*")
	fmt.Println(len(allPings))
}
