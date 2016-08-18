TEXT add(SB), $-8-24
    MOVV    a(FP), R1
    MOVV    b+8(FP), R2
    ADDVU   R2, R1
    MOVV    R1, 16(FP)
    RET
