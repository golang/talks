TEXT add(SB), $0-24
    MOVD    a(FP), R1
    MOVD    b+8(FP), R2
    ADD     R2, R1, R1
    MOVD    R1, 16(FP)
    RET
