package main

import(
	"fmt"
	"github.com/nelhage/gojit"
	"strconv"
	"encoding/binary"
)

func AssembleAddition(kk uint16, vx uint16, chp *Cheep8){
	//fmt.Printf("DEBUG: assembling JIT blog for %d+%d\n", kk, chp.registers[vx])
	block, err := gojit.Alloc(1024) // alloc 64B executable page
	if err != nil{
		panic("JIT allocation failed")
	}
	kkpointer, err := strconv.ParseInt(fmt.Sprintf("%p", &kk), 0, 64)
	vxpointer, err := strconv.ParseInt(fmt.Sprintf("%p", &chp.registers[x]), 0, 64)
	kkpointeru := uint64(kkpointer)
	vxpointeru := uint64(vxpointer)
	fmt.Println(uint64(kkpointer), vxpointer)

	Code := []uint64{
		0x488b0425, kkpointeru, 0x488b1425, vxpointeru, 0x4801d0, 0x48890425, vxpointeru, 0xc3,
	}

	fmt.Println(Code)
	temp := make([]byte, 0)
	for i := range Code{
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, Code[i])
		temp = append(temp, b...)
		fmt.Println(temp)
	}
	block = temp
	fmt.Println(block)
	gojit.Release(block)
}