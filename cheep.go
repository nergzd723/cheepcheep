package main


import(
	"os"
	"fmt"
	"math/rand"
	"time"
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
func ldfont(chp *Cheep8) error{
	font := []byte{0xF0, 0x90, 0x90, 0x90, 0xF0,
		0x20, 0x60, 0x20, 0x20, 0x70,
		0xF0, 0x10, 0xF0, 0x80, 0xF0,
		0xF0, 0x10, 0xF0, 0x10, 0xF0,
		0x90, 0x90, 0xF0, 0x10, 0x10,
		0xF0, 0x80, 0xF0, 0x10, 0xF0,
		0xF0, 0x80, 0xF0, 0x90, 0xF0,
		0xF0, 0x10, 0x20, 0x40, 0x40,
		0xF0, 0x90, 0xF0, 0x90, 0xF0,
		0xF0, 0x90, 0xF0, 0x10, 0xF0,
		0xF0, 0x90, 0xF0, 0x90, 0x90,
		0xE0, 0x90, 0xE0, 0x90, 0xE0,
		0xF0, 0x80, 0x80, 0x80, 0xF0,
		0xE0, 0x90, 0x90, 0x90, 0xE0,
		0xF0, 0x80, 0xF0, 0x80, 0xF0,
		0xF0, 0x80, 0xF0, 0x80, 0x80}
	for count, bt := range font {
		chp.mem[count] = uint8(bt)
	}
	return nil
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
	ldfont(chp)
	fmt.Println("debug: got", n, "bytes of data")
	for i := 0; i < 4096-512; i++ { // some space for a compiler
		chp.mem[i+int(chp.pc)] = temp[i]	
	}
}

func updateTimers(chp *Cheep8){
	if chp.dT > 0{
		chp.dT--
	}
	if chp.sT > 0{
		chp.sT--
	}
}

func Process(chp *Cheep8){
	chp.opc = uint16(chp.mem[chp.pc]) // here's the bug, opcode assumed to be uint8
	chp.opc <<= 8
	chp.opc |= uint16(chp.mem[chp.pc+1])
	Interpret(chp)
	chp.pc +=  2
	updateTimers(chp)
	
}

func Interpret(chp *Cheep8){
	opcode := chp.opc
	x = (opcode & 0x0F00) >> 8
	nnn = opcode & 0x0FFF
	n = opcode & 0x000F
	y = (opcode & 0x00F0) >> 4
	kk = opcode & 0x00FF
	fmt.Printf("opcode:0x%X\n", opcode)
	if opcode == 0x00EE{
		chp.pc = chp.stack[chp.sp]
		chp.sp -= 1
		return
	}
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
		chp.sp++
		chp.stack[chp.sp] = chp.pc
		chp.pc = nnn - 2
	case 0xF000:
		switch opcode & 0x00FF {
		case 0x0033: // store BCD representation of Vx at [i], [i+1], [i+2]
			opreg := chp.registers[x]
			chp.mem[chp.i] = opreg / 100
			chp.mem[chp.i+1] = (opreg / 10) % 10
			chp.mem[chp.i+2] = (opreg % 100) / 10
		case 0x0007:
			chp.registers[x] = chp.dT
		case 0x0015:
			chp.dT = chp.registers[x]
		case 0x0065:
			for V := 0; V < 16; V++ {
				chp.registers[V] = chp.mem[uint16(V)+chp.i]
			}
		case 0x0029:
			chp.i = uint16(chp.registers[x]) * 5  // we have sprites starting at 0x0, 5B each
		default:
			hexopc := fmt.Sprintf("%x", opcode)
			fmt.Println("cheepcheep: bad opcode", hexopc)
			panic("cheepcheep: bad opcode")
		}
	case 0x7000:
		if (jit){
			AssembleAddition(x, kk, chp)
		}else{
			chp.registers[x] += uint8(kk)
		}
	case 0x3000:
		if chp.registers[x] == uint8(kk){
			chp.pc += 2
		}
	case 0x1000:
		chp.pc = nnn - 2
	case 0xC000:
		rand.Seed(time.Now().UnixNano())
		chp.registers[x] = uint8(kk & uint16(rand.Int31n(256)))
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