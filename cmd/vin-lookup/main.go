package main

import (
	"fmt"
	"os"

	"github.com/c-wide/vin-lookup/internal/excel"
	"github.com/c-wide/vin-lookup/internal/lookup"
)

func main() {
	fmt.Println("Vin Lookup Initialized")

	argsLen := len(os.Args[1:])
	if argsLen == 0 || argsLen > 1 {
		fmt.Println("Incorrect number of arguments provided")
		os.Exit(1)
	}

	data, err := excel.ProcessFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := lookup.RequestVinInfo(data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := excel.WriteFile(data, os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Vin Lookup Completed")
}
