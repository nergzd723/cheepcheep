package main

import(
	"syscall"
	"unsafe"
	"fmt"
)

func AssembleAddition(kk uint16, vx uint16) func() uint8{
	memory, err := syscall.Mmap(100500, 0, 4096, syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC, syscall.MAP_PRIVATE | syscall.MAP_ANONYMOUS)
	if err != nil{
		panic(err)
	}
	i := 0
	memory[0] = 0xc3 // RET opcode
	i+=1
	fmt.Println(memory)
	f := unsafe.Pointer(&memory)
	r := *(*func() uint8) (f)
	r()
	fmt.Println("alive")
	return r
}