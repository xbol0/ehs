package ehs

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type ScanParams struct {
	Ips         []string
	BaseDomains []string
	KeysReader  *os.File
	Timeout     time.Duration
	Threads     int
}

func getTitle(body string) string {
	match := regexp.
		MustCompile(`<title>([\s\S]*?)</title>`).
		FindStringSubmatch(body)

	if match == nil {
		return ""
	}

	return strings.TrimSpace(match[1])
}

func workerScan(input <-chan *RequestConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	for cfg := range input {
		res, err := Request(cfg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		fmt.Fprintf(os.Stdout, "%d\t%s\t%s\t%s\t%s\n",
			res.Status, cfg.Url, cfg.Host, res.Type, res.Title)
	}
}

func startWorkers(n int) (chan<- *RequestConfig, *sync.WaitGroup) {
	wg, ch := new(sync.WaitGroup), make(chan *RequestConfig, n)
	wg.Add(n)

	for i := 0; i < n; i++ {
		go workerScan(ch, wg)
	}

	return ch, wg
}

func Scan(sp *ScanParams) {
	ipList := make([]net.IP, len(sp.Ips))

	for _, ip := range sp.Ips {
		i := net.ParseIP(ip)
		if i != nil {
			ipList = append(ipList, i)
		}
	}

	domainList := make([]string, 0, len(topRootDomains))

	for _, s := range sp.BaseDomains {
		if strings.Index(s, ".") == -1 {
			for _, root := range topRootDomains {
				domainList = append(domainList, s+"."+root)
			}
		} else {
			domainList = append(domainList, s)
		}
	}

	ch, wg := startWorkers(sp.Threads)
	rd := bufio.NewReader(sp.KeysReader)

	for _, ip := range sp.Ips {
		for _, base := range domainList {
			ch <- &RequestConfig{
				Url:     "http://" + ip + "/",
				Host:    base,
				Timeout: sp.Timeout,
			}
			ch <- &RequestConfig{
				Url:     "https://" + ip + "/",
				Host:    base,
				Timeout: sp.Timeout,
			}
		}
	}

	for {
		sub, _, err := rd.ReadLine()

		if err != nil {
			break
		}

		// Skip comment and empty line.
		if len(bytes.TrimSpace(sub)) == 0 || sub[0] == byte('#') {
			continue
		}

		for _, ip := range sp.Ips {
			for _, base := range domainList {
				ch <- &RequestConfig{
					Url:     "http://" + ip + "/",
					Host:    string(sub) + "." + base,
					Timeout: sp.Timeout,
				}
				ch <- &RequestConfig{
					Url:     "https://" + ip + "/",
					Host:    string(sub) + "." + base,
					Timeout: sp.Timeout,
				}
			}
		}
	}

	close(ch)
	wg.Wait()
}
