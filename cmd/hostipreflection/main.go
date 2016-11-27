package main

import (
	"flag"
	"fmt"
	_ "github.com/andyxning/host_ip_reflection/handler"
	"github.com/golang/glog"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	version           string
	commitHash        string
	maxProcs          uint
	logFlushFrequency time.Duration

	ip   string
	port uint
)

const (
	defaultLogFlushFrequency = 5 * time.Second
	defaultMacProcs          = 4

	defaultIP   = "0.0.0.0"
	defaultPort = 3087
)

func processCmdFlags() {
	var versionFlag bool

	flag.BoolVar(&versionFlag, "version", false, "version info")
	flag.UintVar(&maxProcs, "max_procs", defaultMacProcs,
		"max cpu number host ip reflection will use. Can not exceeds the number of max cpu cores.")
	flag.DurationVar(&logFlushFrequency, "log_flush_frequency", defaultLogFlushFrequency,
		"Maximum number of seconds between glog flushes")

	flag.StringVar(&ip, "ip", defaultIP, "host ip reflection listen ip")
	flag.UintVar(&port, "port", defaultPort, "host ip reflection listen port")

	flag.Parse()

	if versionFlag {
		fmt.Printf("version: %s\ncommit: %s\n", version, commitHash)
		os.Exit(0)
	}
}

func flushLogPeriodically() {
	glog.Infof("set glog flush interval to %v", logFlushFrequency)

	go func() {
		flusher := time.NewTicker(logFlushFrequency)
		for {
			select {
			case <-flusher.C:
				glog.Flush()
			}
		}
	}()
}

func mustSetMaxProcs() {
	cpuCapacity := runtime.NumCPU()

	if maxProcs > uint(cpuCapacity) {
		glog.Errorf(
			"max procs specified exceeds max available cpu number. max: %d, specified: %d",
			cpuCapacity, maxProcs,
		)
		panic("max procs exceeds available cpu cores")
	}

	runtime.GOMAXPROCS(int(maxProcs))
	glog.Infof("set max procs to %d", maxProcs)
}

func main() {
	processCmdFlags()

	mustSetMaxProcs()

	flushLogPeriodically()

	glog.Error(http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), nil))
}
