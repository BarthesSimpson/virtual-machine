@17
D=A
@SP
A=M
M=D
@SP
M=M+1
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE0
D;JNE
@SP
A=M-1
M=-1
@SUB0
0;JMP
(FALSE0)
@SP
A=M-1
M=0
(SUB0)
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE1
D;JNE
@SP
A=M-1
M=-1
@SUB1
0;JMP
(FALSE1)
@SP
A=M-1
M=0
(SUB1)
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE2
D;JNE
@SP
A=M-1
M=-1
@SUB2
0;JMP
(FALSE2)
@SP
A=M-1
M=0
(SUB2)
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE3
D;JGE
@SP
A=M-1
M=-1
@SUB3
0;JMP
(FALSE3)
@SP
A=M-1
M=0
(SUB3)
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE4
D;JGE
@SP
A=M-1
M=-1
@SUB4
0;JMP
(FALSE4)
@SP
A=M-1
M=0
(SUB4)
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE5
D;JGE
@SP
A=M-1
M=-1
@SUB5
0;JMP
(FALSE5)
@SP
A=M-1
M=0
(SUB5)
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE6
D;JLE
@SP
A=M-1
M=-1
@SUB6
0;JMP
(FALSE6)
@SP
A=M-1
M=0
(SUB6)
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE7
D;JLE
@SP
A=M-1
M=-1
@SUB7
0;JMP
(FALSE7)
@SP
A=M-1
M=0
(SUB7)
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
@FALSE8
D;JLE
@SP
A=M-1
M=-1
@SUB8
0;JMP
(FALSE8)
@SP
A=M-1
M=0
(SUB8)
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
@53
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=M+D
@112
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=M-D
D=0
@SP
A=M-1
M=D-M
@SP
AM=M-1
D=M
A=A-1
M=M&D
@82
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=M|D
@SP
A=M-1
M=!M
