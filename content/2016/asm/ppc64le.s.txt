TEXT add(SB), $0-24
    MOVD    a(FP), R2
    MOVD    b+8(FP), R3
    ADD     R3, R2
    MOVD    R2, 16(FP)
    RET
