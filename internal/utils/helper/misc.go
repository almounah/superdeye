package helper

import (
	"errors"
	"unsafe"

	"github.com/almounah/superdeye/internal/utils/superdwindows"
)

func GetPEB() uintptr

func GetImageExportDirectory(hModule superdwindows.HANDLE) (pImgExpDir superdwindows.PIMAGE_EXPORT_DIRECTORY, err error) {
	pBase := unsafe.Pointer(hModule)
	pImgDosHeader := superdwindows.PIMAGE_DOS_HEADER(pBase)
	if pImgDosHeader.E_magic != superdwindows.IMAGE_DOS_SIGNATURE {
		return nil, errors.New("Messed Up Getting the DosHeader")
	}

	pImgNtHdrs := superdwindows.PIMAGE_NT_HEADERS32(unsafe.Pointer(uintptr(pBase) + uintptr(pImgDosHeader.E_lfanew)))
	if pImgNtHdrs.Signature != superdwindows.IMAGE_NT_SIGNATURE {
		return nil, errors.New("Messed Up Getting NtHeader")
	}

	if pImgNtHdrs.FileHeader.Machine == superdwindows.IMAGE_FILE_MACHINE_AMD64 {
		pImgNtHdrs64 := superdwindows.PIMAGE_NT_HEADERS64(unsafe.Pointer(pImgNtHdrs))
		ImgOptHdr := pImgNtHdrs64.OptionalHeader
		if ImgOptHdr.Magic != superdwindows.IMAGE_NT_OPTIONAL_HDR64_MAGIC {
			return nil, errors.New("Messed Up getting Image Optional Header for x64 arch")
		}
		pImgExportDir := superdwindows.PIMAGE_EXPORT_DIRECTORY(unsafe.Pointer(uintptr(pBase) + uintptr(ImgOptHdr.DataDirectory.VirtualAddress)))

		return pImgExportDir, nil
	}
	ImgOptHdr := pImgNtHdrs.OptionalHeader
	if ImgOptHdr.Magic != superdwindows.IMAGE_NT_OPTIONAL_HDR32_MAGIC {
		return nil, errors.New("Messed Up getting Image Optional Header for x32 arch")
	}
	pImgExportDir := superdwindows.PIMAGE_EXPORT_DIRECTORY(unsafe.Pointer(uintptr(pBase) + uintptr(ImgOptHdr.DataDirectory.VirtualAddress)))
	return pImgExportDir,nil
}

func NameRvaToString(pBase uintptr, rva superdwindows.DWORD) string {
	addr := uintptr(pBase + uintptr(rva))

	var res []byte
	for i := uintptr(0); ; i++ {
		char := *(*byte)(unsafe.Pointer(addr + i))
		if char == 0 {
			return string(res)
		}
		res = append(res, char)
	}
}

func GetNTDLLAddress() superdwindows.HANDLE {
	ppeb := superdwindows.PPEB64(unsafe.Pointer(uintptr(GetPEB())))

	pDte := superdwindows.PLDR_DATA_TABLE_ENTRY(unsafe.Pointer(unsafe.Add(unsafe.Pointer(ppeb.LoaderData.InMemoryOrderModuleList.Flink.Flink), -0x10)))
	return superdwindows.HANDLE(pDte.DllBase)
}
