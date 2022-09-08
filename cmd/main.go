package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/xbol0/ehs"
)

var (
	wordFile = flag.String("f", "", "Word list file, read from stdin if not set.")
	timeout  = flag.Int("t", 5, "Timeout limit.")
	threads  = flag.Int("n", 5, "Threads count to used send request.")
	showHelp = flag.Bool("h", false, "Show this help message.")
	domain   = flag.String("d", "", "REQUIRED. Domain or keyword.")
	skipFail = flag.Bool("S", false, "Filter non-2XX response.")
)

func main() {
	flag.Parse()

	if *showHelp {
		help()
		return
	}

	args := flag.Args()

	if len(args) == 0 {
		help()
		os.Exit(1)
		return
	}

	if len(*domain) == 0 {
		help()
		os.Exit(1)
		return
	}

	subReader := os.Stdin

	if len(*wordFile) != 0 {
		f, err := os.Open(*wordFile)

		if err != nil {
			panic(err)
		}

		subReader = f
	}

	ehs.Scan(&ehs.ScanParams{
		Ips:         args,
		BaseDomains: []string{*domain},
		KeysReader:  subReader,
		Timeout:     time.Second * time.Duration(*timeout),
		Threads:     *threads,
		SkipFail:    *skipFail,
	})
}

func help() {
	me, err := os.Executable()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Usage:\n  %s [-d domain] [-f file] [-t timeout] [-n threads] [-h] [-S] IP...\n",
		path.Base(me))
	fmt.Printf("\nParameters:\n")

	flag.PrintDefaults()
	fmt.Printf("\nMore information:\n  https://github.com/xbol0/ehs")
}
