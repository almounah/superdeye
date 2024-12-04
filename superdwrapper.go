package superdeye

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func NtAllocateVirtualMemory(handle windows.Handle, address uintptr, size uintptr, alloctype uint32, protect uint32) (value uintptr, NTSTATUS uint32, err error) {
	var baseAddr uintptr
	NTSTATUS, err = SuperdSyscall("NtAllocateVirtualMemory",
		uintptr(handle),
		uintptr(unsafe.Pointer(&baseAddr)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&size)),
		uintptr(alloctype),
		uintptr(protect),
	)

	return baseAddr, NTSTATUS, err
}

func NtWriteVirtualMemory(handle windows.Handle, baseAddress uintptr, buffer *byte, size uintptr, numberOfBytesWritten *uintptr) (NTSTATUS uint32, err error) {
	NTSTATUS, err = SuperdSyscall("NtWriteVirtualMemory", 
        uintptr(handle),
        baseAddress, 
        uintptr(unsafe.Pointer(buffer)), 
        uintptr(size), 
        uintptr(unsafe.Pointer(numberOfBytesWritten)),
    )

    return NTSTATUS, err
}
