package main

import(
	"os"
	"fmt"
)
type Cheep8 struct {
    mem [4096]byte
    registers [16]byte
	vram [32*64*1]byte // Black and white 32*64 display
	stack [16]byte
	pc uint16
	sp uint16
	i uint16
	opc uint16
	dT byte
	sT byte
}
func checkForErr(err error){
	if err != nil{
		panic("cheepcheep: error")
	}
}
func loadROM(rom string, chp Cheep8) Cheep8{
	chp.pc = 0x200
	dat, err := os.Open(rom)
	checkForErr(err)
	temp := make([]byte, 4096)
	n, err := dat.Read(temp)
	checkForErr(err)
	fmt.Println("debug: got", n, "bytes of data")
	for i := 0; i < 4096-512; i++ { // some space for a compiler
		chp.mem[i+int(chp.pc)] = temp[i]	
	}
	return chp
}

func Process(chp Cheep8){
	chp.pc = chp.pc + 2
	chp.opc = uint16(chp.mem[chp.pc]) // here's the bug, opcode assumed to be uint8
	chp.opc <<= 8
	chp.opc |= uint16(chp.mem[chp.pc+1])
	fmt.Println(chp.opc)
	Interpret(chp, uint16(chp.opc))
	chp.pc = chp.pc + 2
}

func Interpret(chp Cheep8, opcode uint16){
	switch opcode{
	default:
		hexopc := fmt.Sprintf("%x", opcode)
		fmt.Println("cheepcheep: bad opcode ", hexopc)
		panic("cheepcheep: bad opcode")
	}
}

func main(){
	var chp Cheep8
	chp = loadROM("pong.ch8", chp)
	for ;;{
		Process(chp)
	}
}