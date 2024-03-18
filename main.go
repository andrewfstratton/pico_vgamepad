package main

import (
	"fmt"
	"machine"
	"strconv"
	"strings"
	"time"

	"machine/usb/hid/joystick"
)

const MAX_BTN_NUM = 16 // N.B. Min must start at 1
var CustomGamepadHIDReport = []byte{
	0x05, 0x01, // Usage Page (Generic Desktop Ctrls)
	0x09, 0x05, // Usage (Game Pad)
	0xA1, 0x01, // Collection (Application)
	0x85, 1, //   Report ID
	0x05, 0x09, //   Usage Page (Button)
	0x19, 1, //   Usage Minimum (Button 1)
	0x29, MAX_BTN_NUM, //   Usage Maximum (Button 16)
	0x15, 0, //   Logical Minimum
	0x25, 1, //   Logical Maximum
	0x75, 1, //   Report Size
	0x95, MAX_BTN_NUM, //   Report Count
	0x81, 0x02, //   Input (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0x05, 0x01, //   Usage Page (Generic Desktop Ctrls)
	0x16, 0x01, 0x80, // Logical Minimum (-32767)
	0x26, 0xFF, 0x7F, // Logical Maximum (32767)
	0x09, 0x30, //   Usage (X) Left Stick
	0x09, 0x31, //   Usage (Y) Left Stick
	0x09, 0x32, //   Usage (Z) Right Stick
	0x09, 0x33, //   Usage (Rx) Right Stick
	0x09, 0x34, //   Usage (Ry) Left trigger
	0x09, 0x35, //   Usage (Rz) Right trigger
	0x75, 16, //   Report Size
	0x95, 6, //   Report Count
	0x81, 0x02, //   Input (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0xC0, // End Collection
}

func CustomDefinitions() joystick.Definitions {
	return joystick.Definitions{
		ReportID:     1,
		ButtonCnt:    MAX_BTN_NUM,
		HatSwitchCnt: 0,
		AxisDefs: []joystick.Constraint{
			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // -Left +Right
			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // **-Up +Down**
			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // -Left +Right
			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767}, // **-Up +Down**
			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767},
			{MinIn: -32767, MaxIn: 32767, MinOut: -32767, MaxOut: 32767},
		},
	}
}

func init() {
	// Do this through flash ldflags options
	// descriptor.Configure(0x2e8a, 0x0004)
	joystick.UseSettings(CustomDefinitions(), nil, nil, CustomGamepadHIDReport) // XboxHIDDescriptor)
}

func main() {
	gamepad := joystick.Port()

	uart := machine.Serial
	uart.Configure(machine.UARTConfig{TX: machine.UART_TX_PIN, RX: machine.UART_RX_PIN})

	message := ""

	time.Sleep(4 * time.Second)

	fmt.Println("Started")
	for {
		if uart.Buffered() > 0 {
			b, err := uart.ReadByte()
			if err != nil {
				print("err:" + err.Error())
			} else {
				if b != byte(13) {
					message += string(b)
				} else {
					prefix := message[0:1]
					message = strings.TrimPrefix(message, prefix)
					switch prefix {
					case "B", "b": // button press/release
						btn_64, err := strconv.Atoi(message)
						if err == nil {
							press := false     // release by default
							if prefix == "B" { // Button press
								press = true
							}
							gamepad.SetButton(int(btn_64), press)
							gamepad.SendState()
						}
					case "X", "Y", "Z", "x", "y", "z": // Axis
						val_64, err := strconv.Atoi(message)
						if err == nil {
							axis := -1 // In case it doesn't match
							if -32767 <= val_64 && val_64 <= 32767 {
								switch prefix {
								case "X": // Left Stick < to >
									axis = 0
								case "Y": // Left Stick ^ to v
									axis = 1
								case "Z": // Right Stick < to >
									axis = 2
								case "x": // Right Stick ^ to v
									axis = 3
								case "y": // Left Trigger
									axis = 4
								case "z": // Right Trigger
									axis = 5
								}
								if axis != -1 {
									gamepad.SetAxis(axis, int(val_64))
									gamepad.SendState()
								}
							}
						}
					}
					message = ""
				}
			}
		}
	}
}
