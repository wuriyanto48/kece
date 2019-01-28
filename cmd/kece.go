package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Bhinneka/kece"
)

func main() {

	args, err := kece.ParseArgs()

	if err != nil {
		fmt.Printf("\033[31m%s\033[0m%s", err.Error(), "\n")
		args.Help()
		os.Exit(0)
	}

	if args.ShowVersion {
		fmt.Printf("\033[35m%s version %s (runtime: %s)\033[0m%s", os.Args[0], kece.Version, runtime.Version(), "\n")
		os.Exit(0)
	}

	server := kece.NewServer(args.Network, args.Port, nil)

	if err := server.Start(); err != nil {
		fmt.Printf("\033[31m%s\033[0m%s", err.Error(), "\n")
		os.Exit(1)
	}
}
