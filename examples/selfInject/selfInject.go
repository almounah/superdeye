package main

import (
	"fmt"
	"unsafe"

	"github.com/almounah/superdeye"
	"golang.org/x/sys/windows"
)

func main()  {
    
    var enter string;

	size := 100
    hSelf := uintptr(0xffffffffffffffff)
    var baseAddr uintptr;

    _, err := superdeye.SuperdSyscall("NtAllocateVirtualMemory",
        hSelf, 
        uintptr(unsafe.Pointer(&baseAddr)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&size)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE,
	)
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Scanln(&enter)
}
