TEXT add(SB), $-8-24
    MOVD    a(FP), R0
    MOVD    b+8(FP), R1
    ADD     R1, R0
    MOVD    R0, 16(FP)
    RET
