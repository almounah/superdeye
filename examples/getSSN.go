package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"

	"SuperdEye/internal/manalocator"
	"SuperdEye/internal/superdsyscall"
	"SuperdEye/internal/utils/superdwindows"
)

func main() {
    var enter string;
	ntdllHandle := windows.NewLazyDLL("ntdll.dll").Handle()
	syscallTool, _ := manalocator.LookupSSNAndTrampoline("NtAllocateVirtualMemory", superdwindows.HANDLE(ntdllHandle))
	fmt.Println(syscallTool.Ssn, syscallTool.SyscallInstructionAddress)

	size := 100
    hSelf := uintptr(0xffffffffffffffff)
    var baseAddr uintptr;
	superdsyscall.ExecIndirectSyscall(uint16(syscallTool.Ssn), 
        uintptr(syscallTool.SyscallInstructionAddress), 
        hSelf, 
        uintptr(unsafe.Pointer(&baseAddr)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&size)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE,
	)

    fmt.Scanln(&enter)
}
