// +build !noasm
#include "textflag.h"

// func GetPEB() uintptr
TEXT ·GetPEB(SB),NOSPLIT|NOFRAME,$0-8
    MOVQ 0x60(GS), CX   // Read the value at offset 0x60 from the GS segment into CX
    MOVQ CX, ret+0(FP)  // Store the value in the return slot (CX -> return)
    RET


