package main

import(
	"syscall"
	"unsafe"
	"fmt"
)

func AssembleAddition(kk uint16, vx uint16, chp *Cheep8){
	FunctionCode := []uint16{
		0x488b, 0x0425, *(*uint16)(unsafe.Pointer(&kk)), 0x0, 0x488b, 0x1425, *(*uint16)(unsafe.Pointer(&vx)), 0x0, 0x4801, 0xc2, 0x4889, 0x0425, *(*uint16)(unsafe.Pointer(&vx)), 0x0, 0x00c3,
	   }
	memory, err := syscall.Mmap(-1, 0, 64, syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC, syscall.MAP_PRIVATE | syscall.MAP_ANONYMOUS)
	if err != nil{
		panic(err)
	}
	j := 0
 	for i := range FunctionCode {
  	memory[j] = byte(FunctionCode[i] >> 8)
  	memory[j+1] = byte(FunctionCode[i])
  	j = j + 2
	}
	fmt.Println(memory)
	f := (uintptr)(unsafe.Pointer(&memory))
	r := *(*func() uint8) (unsafe.Pointer(&f))
	r()
}