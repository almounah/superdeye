package manalocator

import (
	"fmt"
	"unsafe"

	"SuperdEye/internal/utils/superdwindows"
)

func LookupSSN(syscallName string, hModule superdwindows.HANDLE) (ssn int, err error) {
	pBase := unsafe.Pointer(hModule)
	pImgExportDir := superdwindows.GetImageExportDirectory(hModule)
	numFunction := pImgExportDir.NumberOfFunctions

	//	AddressOfFuntionArray := unsafe.Slice((*superdwindows.DWORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfFunctions))), pImgExportDir.NumberOfFunctions)
	AddressOfNamesArray := unsafe.Slice((*superdwindows.DWORD)(unsafe.Pointer(uintptr(pBase)+uintptr(pImgExportDir.AddressOfNames))), pImgExportDir.NumberOfFunctions)

	for i := superdwindows.DWORD(0); i < numFunction; i++ {
		functionNameRVA := AddressOfNamesArray[i]
        s := superdwindows.NameRvaToString(uintptr(pBase), functionNameRVA)
        fmt.Print(s)
	}
	return 0, nil
}
