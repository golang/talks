TEXT add(SB), $0-12
    MOVL    a+4(FP), BX
    ADDL    b+8(FP), BX
    MOVL    BX, 12(FP)
    RET
