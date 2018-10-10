/ a3 -- pdp-11 assembler pass 1

assem:
        jsr     pc,readop
        jsr     pc,checkeos
        br      ealoop
        tst     ifflg
        beq     3f
        cmp     r4,$200
        blos    assem
        cmpb    (r4),$21   /if
        bne     2f
        inc     ifflg
2:
        cmpb    (r4),$22   /endif
        bne     assem
        dec     ifflg
        br      assem

