package main

import (
	"fmt"
)

func main() {
	ports, err := GetCompleteSerialID("usb-FTDI_FT231X_USB_UART")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ports)
}
