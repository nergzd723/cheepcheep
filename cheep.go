package main

import(
	"os"
	"fmt"
)
type Cheep8 struct {
    mem [4096]byte
    registers [16]byte
    vram [32*64*1]byte // Black and white 32*64 display
}
func checkForErr(err error){
	if err != nil{
		panic("cheepcheep: error")
	}
}
func loadROM(rom string, chp Cheep8){
	dat, err := os.Open(rom)
	checkForErr(err)
	temp := make([]byte, 4096)
	n, err := dat.Read(temp)
	checkForErr(err)
	fmt.Println("debug: got ", n, " bytes of data")
	for i := 0; i < 4096; i++ {
		chp.mem[i] = temp[i]	
	}
	fmt.Println(chp.mem)
}
func main(){
	var chp Cheep8
	loadROM("pong.ch8", chp)
}