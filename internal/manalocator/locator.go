package manalocator

import (
	"fmt"
	"unsafe"

	"SuperdEye/internal/utils/superdwindows"
)

func LookupSSN(syscallName string, hModule superdwindows.HANDLE) (ssn uint8, err error) {
	pBase := unsafe.Pointer(hModule)
	pImgExportDir := superdwindows.GetImageExportDirectory(hModule)
	numFunction := pImgExportDir.NumberOfFunctions

	AddressOfFuntionArray := unsafe.Slice((*superdwindows.DWORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfFunctions))), pImgExportDir.NumberOfFunctions)
	AddressOfNamesArray := unsafe.Slice((*superdwindows.DWORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfNames))), pImgExportDir.NumberOfFunctions)
	AddressOfNameOrdinalArray := unsafe.Slice((*superdwindows.WORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfNameOrdinals))), pImgExportDir.NumberOfFunctions)

	for num := superdwindows.DWORD(0); num < numFunction; num++ {
		functionNameRVA := AddressOfNamesArray[num]
		s := superdwindows.NameRvaToString(uintptr(pBase), functionNameRVA)
		if s == syscallName {
			fmt.Println("Found " + s)
			functionAddress := uintptr(pBase) + uintptr(AddressOfFuntionArray[AddressOfNameOrdinalArray[num]])

			if checkIfCleanSSN(functionAddress) {
				low := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 4))
				high := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(functionAddress), 5))

				return uint8((high << 8) | low), nil
			}

			// Search the neighbors if SSN is hooked
			for neighborIndex := 1; neighborIndex < 200; neighborIndex++ {
				upNeighborFunctionAddress := uintptr(unsafe.Add(unsafe.Pointer(functionAddress), neighborIndex*32))
				if checkIfCleanSSN(upNeighborFunctionAddress) {
					low := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(upNeighborFunctionAddress), 4))
					high := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(upNeighborFunctionAddress), 5))
					return uint8((high<<8)|low) - uint8(neighborIndex), nil
				}
				downNeighborFunctionAddress := uintptr(unsafe.Add(unsafe.Pointer(functionAddress), -neighborIndex*32))
				if checkIfCleanSSN(downNeighborFunctionAddress) {
					low := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(downNeighborFunctionAddress), 4))
					high := *(*superdwindows.BYTE)(unsafe.Add(unsafe.Pointer(downNeighborFunctionAddress), 5))
					return uint8((high<<8)|low) + uint8(neighborIndex), nil
				}

			}

		}
	}
	return 0, nil
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
