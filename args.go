package kece

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Arguments struct will hold flag and arguments from stdin
type Arguments struct {
	Auth            string
	Network         string
	Port            string
	DataStorageType string
	ShowVersion     bool
	Help            func()
}

// ParseArgs function, this function will parse flag and arguments from stdin to Arguments struct
func ParseArgs() (*Arguments, error) {
	var (
		auth            string
		network         string
		port            string
		dataStorageType string
		showVersion     bool
	)

	flag.StringVar(&auth, "auth", "", "set server auth eg: -net tcp")
	flag.StringVar(&network, "net", "tcp", "network type eg: -net tcp")
	flag.StringVar(&port, "port", "9000", "port to listen eg: -port 9000")
	flag.StringVar(&dataStorageType, "ds", HashMap, "data storage type (hashmap or binary tree)")

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")

	flag.Usage = func() {
		printGreenColor(Banner)
		fmt.Fprintln(os.Stderr)
		printGreenColor("   **-----------------------------------------------**   ")
		fmt.Println()
		printGreenColor("    Kece (an Experimental Distributed Key Value Store)   ")
		fmt.Fprintln(os.Stderr, "")
		printGreenColor("	-net  | --net network type eg: -net tcp")
		printGreenColor("	-port | --port port to listen eg: -port 9000")
		printGreenColor("	-auth | --auth if you want to client send auth before exchange data")
		printGreenColor("	-ds   | --ds  acronym from (data storage), ")
		printGreenColor("	                you can choose either type (hashmap or binary tree)")
		printGreenColor("	-v    | --version show kece version")
		printGreenColor("	-h    | --help show help and usage")
		fmt.Println()
		printGreenColor("   **-----------------------------------------------**   ")
		printGreenColor("   Running: ")
		printGreenColor("   kece -port 8000 -net tcp -ds (bt/hashmap)")
		fmt.Println()

	}

	flag.Parse()

	if len(network) <= 0 {
		return &Arguments{Help: flag.Usage}, errors.New("	(-net) arg required")
	}

	if len(port) <= 0 {
		return &Arguments{Help: flag.Usage}, errors.New("	(-port) arg required")
	}

	return &Arguments{
		Auth:            auth,
		Network:         network,
		Port:            port,
		DataStorageType: dataStorageType,
		ShowVersion:     showVersion,
		Help:            flag.Usage,
	}, nil
}

func printRedColor(s string) {
	fmt.Printf("\033[31m%s\033[0m%s", s, "\n")
}

func printGreenColor(s string) {
	fmt.Printf("\033[32m%s\033[0m%s", s, "\n")
}

func printYellowColor(s string) {
	fmt.Printf("\033[33m%s\033[0m%s", s, "\n")
}

func printCyanColor(s string) {
	fmt.Printf("\033[36m%s\033[0m%s", s, "\n")
}
