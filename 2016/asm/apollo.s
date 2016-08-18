# TO ENTER A JOB REQUEST REQUIRING NO VAC AREA:

          COUNT     02/EXEC
                
NOVAC     INHINT
          AD        FAKEPRET     # LOC(MPAC +6) - LOC(QPRET)
          TS        NEWPRIO      # PRIORITY OF NEW JOB + NOVAC C(FIXLOC)

          EXTEND
          INDEX     Q            # Q WILL BE UNDISTURBED THROUGHOUT.
          DCA       0            # 2CADR OF JOB ENTERED.
          DXCH      NEWLOC
          CAF       EXECBANK
          XCH       FBANK
          TS        EXECTEM1
          TCF       NOVAC2       # ENTER EXECUTIVE BANK.
