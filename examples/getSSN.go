package main

import (
	"SuperdEye/internal/manalocator"
	"SuperdEye/internal/utils/superdwindows"
	"fmt"

	"golang.org/x/sys/windows"
)


func main()  {

    ntdllHandle := windows.NewLazyDLL("ntdll.dll").Handle()
    ssn, _ := manalocator.LookupSSN("NtCreateThread", superdwindows.HANDLE(ntdllHandle))
    fmt.Println(ssn)
    
}
