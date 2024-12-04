package manalocator

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/almounah/superdeye/internal/utils/helper"
	"github.com/almounah/superdeye/internal/utils/superdwindows"
)

/*
 * ssn in the ssn number
 * cleanSyscall is the Syscall address found in NTDLL
 **/
type SuperdSyscallTool struct {
	Ssn                       uint8
	SyscallInstructionAddress superdwindows.PVOID
}

func LookupSSNAndTrampoline(syscallName string, hModule superdwindows.HANDLE) (superdSyscallTool SuperdSyscallTool, err error) {
	pBase := unsafe.Pointer(hModule)
	pImgExportDir := helper.GetImageExportDirectory(hModule)
	numFunction := pImgExportDir.NumberOfFunctions

	AddressOfFuntionArray := unsafe.Slice((*superdwindows.DWORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfFunctions))), pImgExportDir.NumberOfFunctions)
	AddressOfNamesArray := unsafe.Slice((*superdwindows.DWORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfNames))), pImgExportDir.NumberOfFunctions)
	AddressOfNameOrdinalArray := unsafe.Slice((*superdwindows.WORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfNameOrdinals))), pImgExportDir.NumberOfFunctions)

	for num := superdwindows.DWORD(0); num < numFunction; num++ {
		functionNameRVA := AddressOfNamesArray[num]
		s := helper.NameRvaToString(uintptr(pBase), functionNameRVA)
		if s == syscallName {
			fmt.Println("Found " + s)
			functionAddress := uintptr(pBase) + uintptr(AddressOfFuntionArray[AddressOfNameOrdinalArray[num]])

			if checkIfCleanSSN(functionAddress) {
				low := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 4))
				high := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 5))
				ssn := uint8((high << 8) | low)
				syscallAddress, _ := findSyscallAddress(functionAddress)
				res := SuperdSyscallTool{ssn, syscallAddress}
				return res, nil
			}

			// Search the neighbors if SSN is hooked
			for neighborIndex := 1; neighborIndex < 200; neighborIndex++ {
				upNeighborFunctionAddress := uintptr(unsafe.Add(unsafe.Pointer(functionAddress), neighborIndex*32))
				if checkIfCleanSSN(upNeighborFunctionAddress) {
					low := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(upNeighborFunctionAddress), 4))
					high := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(upNeighborFunctionAddress), 5))
					ssn := uint8((high<<8)|low) - uint8(neighborIndex)
					syscallAddress, _ := findSyscallAddress(upNeighborFunctionAddress)
					res := SuperdSyscallTool{ssn, syscallAddress}
					return res, nil
				}
				downNeighborFunctionAddress := uintptr(unsafe.Add(unsafe.Pointer(functionAddress), -neighborIndex*32))
				if checkIfCleanSSN(downNeighborFunctionAddress) {
					low := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(downNeighborFunctionAddress), 4))
					high := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(downNeighborFunctionAddress), 5))
					ssn := uint8((high<<8)|low) + uint8(neighborIndex)
					syscallAddress, _ := findSyscallAddress(downNeighborFunctionAddress)
					res := SuperdSyscallTool{ssn, syscallAddress}
					return res, nil
				}

			}

		}
	}
	return SuperdSyscallTool{0, 0}, nil
}

func checkIfCleanSSN(functionAddress uintptr) bool {
	val0 := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 0))
	val1 := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 1))
	val2 := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 2))
	val3 := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 3))
	val6 := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 6))
	val7 := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 7))

	if val0 == 0x4c && val1 == 0x8b && val2 == 0xd1 && val3 == 0xb8 && val6 == 0x00 && val7 == 0x00 {
		return true
	}

	return false
}

func findSyscallAddress(functionAddress uintptr) (pAddress superdwindows.PVOID, err error) {
	for z := 1; z < 200; z++ {
		pAddress := unsafe.Add(unsafe.Pointer(functionAddress), z)
		pAddressNext := unsafe.Add(unsafe.Pointer(functionAddress), z+1)

		value := *(*superdwindows.BYTE)(pAddress)
		valueNext := *(*superdwindows.BYTE)(pAddressNext)

		if value == 0x0F && valueNext == 0x05 {
			return superdwindows.PVOID(pAddress), nil
		}
	}
	return 0, errors.New("Not Found Syscall instruction in Ntdll")
}
