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

	// version flag present
	if args.ShowVersion {
		fmt.Printf("\033[35m%s version %s (runtime: %s)\033[0m%s", os.Args[0], kece.Version, runtime.Version(), "\n")
		os.Exit(0)
	}

	var commander kece.Commander
	switch args.DataStorageType {
	case kece.HashMap:
		// create the global database
		db := make(map[string]*kece.Schema)
		// call commander constructor
		commander = kece.NewCommander(db)
	case kece.BinarySearchTree:
		commander = kece.NewCommanderBST()
	default:
		fmt.Printf("\033[31minvalid data storage type\033[0m\n")
		os.Exit(1)
	}

	// call kece constructor
	server := kece.NewServer(args, commander)

	if err := server.Start(); err != nil {
		fmt.Printf("\033[31m%s\033[0m\n", err.Error())
		os.Exit(1)
	}
}
