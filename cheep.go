package main


import(
	"os"
	"fmt"
)

var nnn uint16
var n uint16
var x uint16
var y uint16
var kk uint16
var jit bool
type Cheep8 struct {
    mem [4096]byte
    registers [16]byte
	vram [32*64*1]byte // Black and white 32*64 display
	stack [16]uint16
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
func loadROM(rom string, chp *Cheep8){
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
}

func Process(chp *Cheep8){
	chp.opc = uint16(chp.mem[chp.pc]) // here's the bug, opcode assumed to be uint8
	chp.opc <<= 8
	chp.opc |= uint16(chp.mem[chp.pc+1])
	Interpret(chp)
	chp.pc +=  2
	
}

func Interpret(chp *Cheep8){
	opcode := chp.opc
	x = (opcode & 0x0F00) >> 8
	nnn = opcode & 0x0FFF
	n = opcode & 0x000F
	y = (opcode & 0x00F0) >> 4
	kk = opcode & 0x00FF
	fmt.Printf("opcode:0x%X\n", opcode)
	switch opcode & 0xF000{ // get only first one
	case 0x6000:
		chp.registers[x] = uint8(kk)
		break
	case 0xA000: // set i = nnn
		chp.i = nnn
		break
	case 0xD000:
		// skip drawing
		break
	case 0x2000: // call subroutine at nnn
		chp.stack[chp.sp] = chp.pc
		chp.pc = nnn - 2
		chp.sp++
	case 0xF000:
		switch opcode & 0x00FF {
		case 0x0033: // store BCD representation of Vx at [i], [i+1], [i+2]
			opreg := chp.registers[x]
			chp.mem[chp.i] = opreg / 100
			chp.mem[chp.i+1] = (opreg / 10) % 10
			chp.mem[chp.i+2] = (opreg % 100) / 10
		}
	case 0x7000:
		if (jit){
			AssembleAddition(x, kk, chp)
		}
	default:
		hexopc := fmt.Sprintf("%x", opcode)
		fmt.Println("cheepcheep: bad opcode", hexopc)
		panic("cheepcheep: bad opcode")
	}
}

func main(){
	var chp Cheep8
	jit = true
	loadROM("pong.ch8", &chp)
	for ;;{
		Process(&chp)
	}
}