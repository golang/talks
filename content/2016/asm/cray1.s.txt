ident slice
         V6        0               ; initialize S
         A4        S0              ; initialize *x
         A5        S1              ; initialize *y
         A3        S2              ; initialize i
loop     S0        A3
         JSZ       exit            ; if S0 == 0 goto exit
         VL        A3              ; set vector length
         V11       ,A4,1           ; load slice of x[i], stride 1
         V12       ,A5,1           ; load slice of y[i], stride 1
         V13       V11 *F V12      ; slice of x[i] * y[i]
         V6        V6 +F V13       ; partial sum
         A14       VL              ; get vector length of this iteration
         A4        A4 + A14        ; *x = *x + VL
         A5        A5 + A14        ; *y = *y + VL
         A3        A3 - A14        ; i = i - VL
         J        loop
 exit
