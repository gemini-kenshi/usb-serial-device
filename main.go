package main

import (
	"fmt"
)

func main() {
	port, err := GetPortBySerialID("usb-FTDI_FT231X_USB_UART")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(port)
}
