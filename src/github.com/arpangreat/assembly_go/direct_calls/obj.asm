main.Add STEXT nosplit size=3 args=0x8 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (direct_calls.go:3)	TEXT	main.Add(SB), NOSPLIT|ABIInternal, $0-8
	0x0000 00000 (direct_calls.go:3)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (direct_calls.go:3)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (direct_calls.go:3)	FUNCDATA	$5, main.Add.arginfo1(SB)
	0x0000 00000 (direct_calls.go:3)	FUNCDATA	$6, main.Add.argliveinfo(SB)
	0x0000 00000 (direct_calls.go:3)	PCDATA	$3, $1
	0x0000 00000 (direct_calls.go:4)	ADDL	BX, AX
	0x0002 00002 (direct_calls.go:4)	RET
	0x0000 01 d8 c3                                         ...
main.(*Adder).AddPtr STEXT nosplit size=4 args=0x10 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (direct_calls.go:9)	TEXT	main.(*Adder).AddPtr(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (direct_calls.go:9)	FUNCDATA	$0, gclocals·Plqv2ff52JtlYaDd2Rwxbg==(SB)
	0x0000 00000 (direct_calls.go:9)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (direct_calls.go:9)	FUNCDATA	$5, main.(*Adder).AddPtr.arginfo1(SB)
	0x0000 00000 (direct_calls.go:9)	FUNCDATA	$6, main.(*Adder).AddPtr.argliveinfo(SB)
	0x0000 00000 (direct_calls.go:9)	PCDATA	$3, $1
	0x0000 00000 (direct_calls.go:10)	LEAL	(BX)(CX*1), AX
	0x0003 00003 (direct_calls.go:10)	RET
	0x0000 8d 04 0b c3                                      ....
main.(*Adder).AddVal STEXT nosplit size=4 args=0x10 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (direct_calls.go:13)	TEXT	main.(*Adder).AddVal(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (direct_calls.go:13)	FUNCDATA	$0, gclocals·Plqv2ff52JtlYaDd2Rwxbg==(SB)
	0x0000 00000 (direct_calls.go:13)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (direct_calls.go:13)	FUNCDATA	$5, main.(*Adder).AddVal.arginfo1(SB)
	0x0000 00000 (direct_calls.go:13)	FUNCDATA	$6, main.(*Adder).AddVal.argliveinfo(SB)
	0x0000 00000 (direct_calls.go:13)	PCDATA	$3, $1
	0x0000 00000 (direct_calls.go:14)	LEAL	(BX)(CX*1), AX
	0x0003 00003 (direct_calls.go:14)	RET
	0x0000 8d 04 0b c3                                      ....
main.main STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (direct_calls.go:17)	TEXT	main.main(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (direct_calls.go:17)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (direct_calls.go:17)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (direct_calls.go:26)	RET
	0x0000 c3                                               .
go.cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.main.Add$abstract SDWARFABSFCN dupok size=29
	0x0000 05 6d 61 69 6e 2e 41 64 64 00 01 01 13 61 00 00  .main.Add....a..
	0x0010 00 00 00 00 13 62 00 00 00 00 00 00 00           .....b.......
	rel 16+4 t=31 go.info.int32+0
	rel 24+4 t=31 go.info.int32+0
go.info.main.(*Adder).AddPtr$abstract SDWARFABSFCN dupok size=53
	0x0000 05 6d 61 69 6e 2e 28 2a 41 64 64 65 72 29 2e 41  .main.(*Adder).A
	0x0010 64 64 50 74 72 00 01 01 13 61 64 64 65 72 00 00  ddPtr....adder..
	0x0020 00 00 00 00 13 61 00 00 00 00 00 00 13 62 00 00  .....a.......b..
	0x0030 00 00 00 00 00                                   .....
	rel 32+4 t=31 go.info.*main.Adder+0
	rel 40+4 t=31 go.info.int32+0
	rel 48+4 t=31 go.info.int32+0
go.info.main.(*Adder).AddVal$abstract SDWARFABSFCN dupok size=53
	0x0000 05 6d 61 69 6e 2e 28 2a 41 64 64 65 72 29 2e 41  .main.(*Adder).A
	0x0010 64 64 56 61 6c 00 01 01 13 61 64 64 65 72 00 00  ddVal....adder..
	0x0020 00 00 00 00 13 61 00 00 00 00 00 00 13 62 00 00  .....a.......b..
	0x0030 00 00 00 00 00                                   .....
	rel 32+4 t=31 go.info.*main.Adder+0
	rel 40+4 t=31 go.info.int32+0
	rel 48+4 t=31 go.info.int32+0
main..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
runtime.memequal32·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal32+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*main.Adder. SRODATA dupok size=13
	0x0000 01 0b 2a 6d 61 69 6e 2e 41 64 64 65 72           ..*main.Adder
type..namedata.*func(*main.Adder, int32, int32) int32- SRODATA dupok size=40
	0x0000 00 26 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 41 64  .&*func(*main.Ad
	0x0010 64 65 72 2c 20 69 6e 74 33 32 2c 20 69 6e 74 33  der, int32, int3
	0x0020 32 29 20 69 6e 74 33 32                          2) int32
type.*func(*main.Adder, int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 58 ac 5b be 08 08 08 36 00 00 00 00 00 00 00 00  X.[....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Adder, int32, int32) int32-+0
	rel 48+8 t=1 type.func(*main.Adder, int32, int32) int32+0
type.func(*main.Adder, int32, int32) int32 SRODATA dupok size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 0d c1 03 0f 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 03 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Adder, int32, int32) int32-+0
	rel 44+4 t=-32763 type.*func(*main.Adder, int32, int32) int32+0
	rel 56+8 t=1 type.*main.Adder+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
	rel 80+8 t=1 type.int32+0
type..importpath.main. SRODATA dupok size=6
	0x0000 00 04 6d 61 69 6e                                ..main
type..namedata.AddPtr. SRODATA dupok size=8
	0x0000 01 06 41 64 64 50 74 72                          ..AddPtr
type..namedata.*func(int32, int32) int32- SRODATA dupok size=27
	0x0000 00 19 2a 66 75 6e 63 28 69 6e 74 33 32 2c 20 69  ..*func(int32, i
	0x0010 6e 74 33 32 29 20 69 6e 74 33 32                 nt32) int32
type.*func(int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ac 2e 3a 90 08 08 08 36 00 00 00 00 00 00 00 00  ..:....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int32, int32) int32-+0
	rel 48+8 t=1 type.func(int32, int32) int32+0
type.func(int32, int32) int32 SRODATA dupok size=80
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 7d 9a 59 34 02 08 08 33 00 00 00 00 00 00 00 00  }.Y4...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int32, int32) int32-+0
	rel 44+4 t=-32763 type.*func(int32, int32) int32+0
	rel 56+8 t=1 type.int32+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
type..namedata.AddVal. SRODATA dupok size=8
	0x0000 01 06 41 64 64 56 61 6c                          ..AddVal
type.*main.Adder SRODATA size=104
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ee d6 43 b1 09 08 08 36 00 00 00 00 00 00 00 00  ..C....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 02 00 02 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Adder.+0
	rel 48+8 t=1 type.main.Adder+0
	rel 56+4 t=5 type..importpath.main.+0
	rel 72+4 t=5 type..namedata.AddPtr.+0
	rel 76+4 t=26 type.func(int32, int32) int32+0
	rel 80+4 t=26 main.(*Adder).AddPtr+0
	rel 84+4 t=26 main.(*Adder).AddPtr+0
	rel 88+4 t=5 type..namedata.AddVal.+0
	rel 92+4 t=26 type.func(int32, int32) int32+0
	rel 96+4 t=26 main.(*Adder).AddVal+0
	rel 100+4 t=26 main.(*Adder).AddVal+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.id- SRODATA dupok size=4
	0x0000 00 02 69 64                                      ..id
type.main.Adder SRODATA size=120
	0x0000 04 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 ae d0 cd 54 0f 04 04 19 00 00 00 00 00 00 00 00  ...T............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal32·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.Adder.+0
	rel 44+4 t=5 type.*main.Adder+0
	rel 48+8 t=1 type..importpath.main.+0
	rel 56+8 t=1 type.main.Adder+96
	rel 80+4 t=5 type..importpath.main.+0
	rel 96+8 t=1 type..namedata.id-+0
	rel 104+8 t=1 type.int32+0
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
main.Add.arginfo1 SRODATA static dupok size=5
	0x0000 00 04 04 04 ff                                   .....
main.Add.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·Plqv2ff52JtlYaDd2Rwxbg== SRODATA dupok size=9
	0x0000 01 00 00 00 01 00 00 00 00                       .........
main.(*Adder).AddPtr.arginfo1 SRODATA static dupok size=7
	0x0000 00 08 08 04 0c 04 ff                             .......
main.(*Adder).AddPtr.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
main.(*Adder).AddVal.arginfo1 SRODATA static dupok size=7
	0x0000 00 08 08 04 0c 04 ff                             .......
main.(*Adder).AddVal.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
