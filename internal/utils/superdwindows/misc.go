package superdwindows

import (
	"fmt"
	"unsafe"
)

func GetPEB() uintptr

func GetImageExportDirectory(hModule HANDLE) PIMAGE_EXPORT_DIRECTORY {
	pBase := unsafe.Pointer(hModule)
	pImgDosHeader := PIMAGE_DOS_HEADER(pBase)
	if pImgDosHeader.E_magic != IMAGE_DOS_SIGNATURE {
		fmt.Println("Messed Up Getting the DosHeader")
	}

	pImgNtHdrs := PIMAGE_NT_HEADERS32(unsafe.Pointer(uintptr(pBase) + uintptr(pImgDosHeader.E_lfanew)))
	if pImgNtHdrs.Signature != IMAGE_NT_SIGNATURE {
		fmt.Println("Messed Up getting NTHeader")
	}

	if pImgNtHdrs.FileHeader.Machine == IMAGE_FILE_MACHINE_AMD64 {
		pImgNtHdrs64 := PIMAGE_NT_HEADERS64(unsafe.Pointer(pImgNtHdrs))
		ImgOptHdr := pImgNtHdrs64.OptionalHeader
		if ImgOptHdr.Magic != IMAGE_NT_OPTIONAL_HDR64_MAGIC {
			fmt.Println("Messed Up getting Image Optional Header for x64 arch")
		}
		pImgExportDir := PIMAGE_EXPORT_DIRECTORY(unsafe.Pointer(uintptr(pBase) + uintptr(ImgOptHdr.DataDirectory.VirtualAddress)))

		return pImgExportDir
	}
	ImgOptHdr := pImgNtHdrs.OptionalHeader
	if ImgOptHdr.Magic != IMAGE_NT_OPTIONAL_HDR32_MAGIC {
		fmt.Println("Messed Up getting Image Optional Header for x64 arch")
	}
	pImgExportDir := PIMAGE_EXPORT_DIRECTORY(unsafe.Pointer(uintptr(pBase) + uintptr(ImgOptHdr.DataDirectory.VirtualAddress)))
	return pImgExportDir
}

func NameRvaToString(pBase uintptr, rva DWORD) (string) {
	addr := uintptr(pBase + uintptr(rva))

    var res []byte
	for i := uintptr(0); ; i++ {
		char := *(*byte)(unsafe.Pointer(addr + i))
        res = append(res, char)
        if char == 0 {
            return string(res)
        }
	}
}
