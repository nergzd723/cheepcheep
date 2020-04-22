package main

import(
	"os"
	"fmt"
)
var mem [4096]byte
var registers [16]byte
var vram [32*64*1]byte // Black and white 32*64 display
func checkForErr(err error){
	if err != nil{
		panic("cheepcheep: error")
	}
}
func loadROM(rom string){
	dat, err := os.Open(rom)
	checkForErr(err)
	temp := make([]byte, 4096)
	n, err := dat.Read(temp)
	checkForErr(err)
	fmt.Println("debug: got ", n, " bytes of data")
	for i := 0; i < 4096; i++ {
		mem[i] = temp[i]	
	}
	fmt.Println(mem)
}
func main(){
	loadROM("pong.ch8")
}