package superdwindows

const (
	IMAGE_DOS_SIGNATURE           = 0x5a4d
	IMAGE_NT_SIGNATURE            = 0x00004550
	IMAGE_FILE_MACHINE_AMD64      = 0x8664
	IMAGE_NT_OPTIONAL_HDR64_MAGIC = 0x20b
	IMAGE_NT_OPTIONAL_HDR32_MAGIC = 0x10b
)

type (
	DWORD     uint32
	WORD      uint16
	PWORD     *uint16
	HANDLE    uintptr
	LPCSTR    *uint8
	PBYTE     *uint8
	LONG      int32
	BYTE      uint8
	ULONGLONG uint64
	ULONG     uint32
	CHAR      byte
	PVOID     uintptr
)

type IMAGE_EXPORT_DIRECTORY struct {
	Characteristics       DWORD
	TimeDateStamp         DWORD
	MajorVersion          WORD
	MinorVersion          WORD
	Name                  DWORD
	Base                  DWORD
	NumberOfFunctions     DWORD
	NumberOfNames         DWORD
	AddressOfFunctions    DWORD
	AddressOfNames        DWORD
	AddressOfNameOrdinals DWORD
}
type PIMAGE_EXPORT_DIRECTORY *IMAGE_EXPORT_DIRECTORY

type IMAGE_DOS_HEADER struct {
	E_magic    WORD
	E_cblp     WORD
	E_cp       WORD
	E_crlc     WORD
	E_cparhdr  WORD
	E_minalloc WORD
	E_maxalloc WORD
	E_ss       WORD
	E_sp       WORD
	E_csum     WORD
	E_ip       WORD
	E_cs       WORD
	E_lfarlc   WORD
	E_ovno     WORD
	E_res      [4]WORD
	E_oemid    WORD
	E_oeminfo  WORD
	E_res2     [10]WORD
	E_lfanew   LONG
}
type PIMAGE_DOS_HEADER *IMAGE_DOS_HEADER

type IMAGE_FILE_HEADER struct {
	Machine              WORD
	NumberOfSections     WORD
	TimeDateStamp        DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
	SizeOfOptionalHeader WORD
	Characteristics      WORD
}
type PIMAGE_FILE_HEADER *IMAGE_FILE_HEADER

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress DWORD
	Size           DWORD
}
type PIMAGE_DATA_DIRECTORY *IMAGE_DATA_DIRECTORY

type IMAGE_OPTIONAL_HEADER32 struct {
	Magic                       WORD
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	BaseOfData                  DWORD
	ImageBase                   DWORD
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          DWORD
	SizeOfStackCommit           DWORD
	SizeOfHeapReserve           DWORD
	SizeOfHeapCommit            DWORD
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               IMAGE_DATA_DIRECTORY
}

type IMAGE_NT_HEADERS32 struct {
	Signature      DWORD
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER32
}
type PIMAGE_NT_HEADERS32 *IMAGE_NT_HEADERS32

type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       WORD
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	ImageBase                   ULONGLONG
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          ULONGLONG
	SizeOfStackCommit           ULONGLONG
	SizeOfHeapReserve           ULONGLONG
	SizeOfHeapCommit            ULONGLONG
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               IMAGE_DATA_DIRECTORY
}
type IMAGE_NT_HEADERS64 struct {
	Signature      DWORD
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER64
}
type PIMAGE_NT_HEADERS64 *IMAGE_NT_HEADERS64

type UNICODE_STRING struct {
	Length        WORD
	MaximumLength WORD
	Buffer        PWORD
}

type RTL_USER_PROCESS_PARAMETERS struct {
	Reserved1     [16]BYTE
	Reserved2     [10]PVOID
	ImagePathName UNICODE_STRING
	CommandLine   UNICODE_STRING
}
type PRTL_USER_PROCESS_PARAMETERS *RTL_USER_PROCESS_PARAMETERS

type LIST_ENTRY struct {
	Flink *LIST_ENTRY
	Blink *LIST_ENTRY
}

type PEB_LDR_DATA struct {
	reserved1               [8]BYTE
	reserved2               [3]PVOID
	InMemoryOrderModuleList LIST_ENTRY
}
type PPEB_LDR_DATA *PEB_LDR_DATA

type PEB32 struct {
	reserved1              [2]BYTE
	BeingDebugged          BYTE
	BitField               BYTE
	reserved3              PVOID
	ImageBaseAddress       PVOID
	Ldr                    PPEB_LDR_DATA
	ProcessParameters      PRTL_USER_PROCESS_PARAMETERS
	reserved4              [3]PVOID
	AtlThunkSListPtr       PVOID
	reserved5              PVOID
	reserved6              ULONG
	reserved7              PVOID
	reserved8              ULONG
	AtlThunkSListPtr32     ULONG
	reserved9              [45]PVOID
	reserved10             [96]BYTE
	PostProcessInitRoutine PVOID
	reserved11             [128]BYTE
	reserved12             [1]PVOID
	SessionId              ULONG
}
type PPEB32 *PEB32

type PEB64 struct {
	Reserved1              [2]BYTE
	BeingDebugged          BYTE
	Reserved2              [21]BYTE
	LoaderData             PPEB_LDR_DATA
	ProcessParameters      PRTL_USER_PROCESS_PARAMETERS
	Reserved3              [520]BYTE
	PostProcessInitRoutine PVOID
	Reserved4              [136]BYTE
	SessionId              ULONG
}
type PPEB64 *PEB64

type LDR_DATA_TABLE_ENTRY struct {
	Reserved1                  [2]PVOID
	InMemoryOrderLinks         LIST_ENTRY
	InInitializationOrderLinks LIST_ENTRY
	DllBase                    PVOID
	EntryPoint                 PVOID
	Reserved3                  PVOID
	FullDllName                UNICODE_STRING
	Reserved4                  [8]BYTE
	Reserved5                  [3]PVOID
	Reserved6                  PVOID
	TimeDateStamp              ULONG
}
type PLDR_DATA_TABLE_ENTRY *LDR_DATA_TABLE_ENTRY
