TEXT add(SB), $-4-12
    MOVW    a(FP), R0
    MOVW    b+4(FP), R1
    ADD     R1, R0
    MOVW    R0, 8(FP)
    RET
