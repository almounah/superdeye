package superdeye

import (
	"superdeye/internal/manalocator"
	"superdeye/internal/superdsyscall"
	"superdeye/internal/utils/helper"
)

func SuperdSyscall(syscallName string, argh ...uintptr) (NTSTATUS uint32, err error) {
	ntdllAddress := helper.GetNTDLLAddress()
	syscallTool, _ := manalocator.LookupSSNAndTrampoline(syscallName, ntdllAddress)
	NTSTATUS = superdsyscall.ExecIndirectSyscall(uint16(syscallTool.Ssn), uintptr(syscallTool.SyscallInstructionAddress), argh...)
	return NTSTATUS, nil
}
