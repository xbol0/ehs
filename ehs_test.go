package ehs

import (
	"os"
	"testing"
	"time"
)

func TestScan(t *testing.T) {
	f, err := os.Open("test_sub.txt")

	if err != nil {
		t.Fatal(err)
	}

	sp := &ScanParams{
		Ips:         []string{"45.33.32.156"},
		BaseDomains: []string{"nmap.org"},
		KeysReader:  f,
		Timeout:     time.Second * 3,
		Threads:     4,
	}

	Scan(sp)
}
