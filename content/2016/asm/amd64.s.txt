TEXT add(SB), $0-24
    MOVQ    b+16(FP), AX
    MOVQ    a+8(FP), CX
    ADDQ    CX, AX
    MOVQ    AX, 24(FP)
    RET
