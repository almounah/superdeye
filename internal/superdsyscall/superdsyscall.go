package superdsyscall


func ExecIndirectSyscall(ssn uint16, trampoline uintptr, argh ...uintptr) uint32

