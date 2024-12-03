package main

import (
	"SuperdEye/internal/manalocator"
	"SuperdEye/internal/utils/superdwindows"

	"golang.org/x/sys/windows"
)


func main()  {

    ntdllHandle := windows.NewLazyDLL("ntdll.dll").Handle()
    manalocator.LookupSSN("NtCreateThread", superdwindows.HANDLE(ntdllHandle))
    
}
