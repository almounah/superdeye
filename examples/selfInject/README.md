# Simple Shellcode Injector

This will chain Syscalls to:

- Allocate Memory in the current process
- Write Shellcode in it (without encryption)
- Run a new thread with the shellcode address

It uses SuperdEye.
