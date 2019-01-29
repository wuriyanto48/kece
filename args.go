package kece

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const (
	// Version ,  the version of Kece
	Version = "0.0.0"

	//Banner , show me :)
	Banner = `
 _  __ _____  ______  _____
| |/ /| |__| |   ___|| |__| |
| |\ \| |___ |  |    | |___
|____________|_____________**%**
	`
)

// Arguments struct will hold flag and arguments from stdin
type Arguments struct {
	Network     string
	Port        string
	ShowVersion bool
	Help        func()
}

// ParseArgs function, this function will parse flag and arguments from stdin to Arguments struct
func ParseArgs() (*Arguments, error) {
	var (
		network     string
		port        string
		showVersion bool
	)

	flag.StringVar(&network, "net", "tcp", "network type eg: -net tcp")
	flag.StringVar(&port, "port", "9000", "port to listen eg: -port 9000")

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, Banner)
		fmt.Fprintln(os.Stderr, "   **-----------------------------------------------**   ")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "    Kece (an Experimental Distributed Key Value Store)   ")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "	-net | --net", "network type eg: -net tcp")
		fmt.Fprintln(os.Stderr, "	-port | --port", "port to listen eg: -port 9000")
		fmt.Fprintln(os.Stderr, "	-v | --version", "show kece version")
		fmt.Fprintln(os.Stderr, "	-h | --help", "show help and usage")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "   **-----------------------------------------------**   ")
		fmt.Fprintln(os.Stderr, "   Running: ")
		fmt.Fprintln(os.Stderr, "   kece -port 8000 -net tcp")
		fmt.Fprintln(os.Stderr, "")

	}

	flag.Parse()

	if len(network) <= 0 {
		return &Arguments{Help: flag.Usage}, errors.New("	(-net) arg required")
	}

	if len(port) <= 0 {
		return &Arguments{Help: flag.Usage}, errors.New("	(-port) arg required")
	}

	return &Arguments{
		Network:     network,
		Port:        port,
		ShowVersion: showVersion,
		Help:        flag.Usage,
	}, nil
}
