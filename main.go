package main

import (
	"fmt"
	"os"

	"bitbucket.org/avd/go-ipc/fifo"
)

func main() {
	sendInterrupt()
	fmt.Println("Interrupt Sent...Exiting...")
}

func sendInterrupt() {
	//interruptOpCode := []byte{0xEF}
	interruptOpCode := []byte{0xAB}

	interruptPipe, err := fifo.New("intel8080_interrupt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Unable to connect to named pipe [intel8080_interrupt]: " + err.Error())
	}

	defer interruptPipe.Close()

	interruptPipe.Write(interruptOpCode)
}
