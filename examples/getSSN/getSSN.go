package main

import (
	"fmt"

	"github.com/almounah/superdeye/internal/manalocator"
	"github.com/almounah/superdeye/internal/utils/helper"
)

func main() {
	ntdllHandle := helper.GetNTDLLAddress()

	for true {
		fmt.Println("Enter Syscall Name: ")
		var syscallName string
		fmt.Scanln(&syscallName)
		syscallTool, err := manalocator.LookupSSNAndTrampoline(syscallName, ntdllHandle)
		if err != nil {
			fmt.Println("Messed up ...")
			fmt.Println(err)
		} else {
			fmt.Println("SSN for ", syscallName, " is ", syscallTool.Ssn)
		}
	}
}
