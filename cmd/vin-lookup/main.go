package main

import (
	"fmt"
	"os"

	"github.com/c-wide/go-nhtsa"
	"github.com/c-wide/vin-lookup/internal/excel"
)

func main() {
	fmt.Println("Vin Lookup Initialized")

	argsLen := len(os.Args[1:])
	if argsLen == 0 || argsLen > 1 {
		fmt.Println("Incorrect number of arguments provided")
		os.Exit(1)
	}

	vReqs, err := excel.ProcessFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	vData, err := nhtsa.DecodeVinBatch(vReqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := excel.WriteFile(vReqs, vData, os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Vin Lookup Completed")
}
