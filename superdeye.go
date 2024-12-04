package superdeye

import (
	"github.com/almounah/superdeye/internal/manalocator"
	"github.com/almounah/superdeye/internal/superdsyscall"
	"github.com/almounah/superdeye/internal/utils/helper"
)

func SuperdSyscall(syscallName string, argh ...uintptr) (NTSTATUS uint32, err error) {
	ntdllAddress := helper.GetNTDLLAddress()
	syscallTool, _ := manalocator.LookupSSNAndTrampoline(syscallName, ntdllAddress)
	NTSTATUS = superdsyscall.ExecIndirectSyscall(uint16(syscallTool.Ssn), uintptr(syscallTool.SyscallInstructionAddress), argh...)
	return NTSTATUS, nil
}
