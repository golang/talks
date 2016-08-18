TITLE   COUNT
 
A=1                             ;Define a name for an accumulator.

START:  MOVSI A,-100            ;initialize loop counter.
                                ;A contains -100,,0
LOOP:   HRRZM A,TABLE(A)        ;Use right half of A to index.
        AOBJN A,LOOP            ;Add 1 to both halves (-77,,1 -76,,2 etc.)
                                ;Jump if still negative.
        .VALUE                  ;Halt program.

TABLE:  BLOCK 100               ;Assemble space to fill up.

END START                       ;End the assembly.
