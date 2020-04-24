mov 0xFFF, %rax ; mov ADDRESS_OF_KK, rax
mov 0xFFF, %rdx ; mov ADDRESS_OF_VX, rdx
add %rax, %rdx ; add rax, rdx
movq %rdx, 0xFFF ; mov rdx, ADDRESS_OF_VX
ret ; RET
; WHY DOESN't IT WORK?