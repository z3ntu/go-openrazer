package main

import (
	"fmt"
	"github.com/z3ntu/go-openrazer"
	"os"
)

func printErrorExit(message string, err error) {
	_, _ = fmt.Fprintln(os.Stderr, message, err)
	os.Exit(1)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	manager, err := openrazer.NewManager()
	if err != nil {
		printErrorExit("Failed to create manager: ", err)
	}

	version, err := manager.GetVersion()
	if err != nil {
		printErrorExit("Failed to get version: ", err)
	}
	fmt.Println("Version: ", version)

	devices := manager.GetDevices()
	for _, device := range devices {
		features, err := device.GetSupportedFeatures()
		if err != nil {
			printErrorExit("Failed to get features: ", err)
		}
		fx, err := device.GetSupportedFx()
		if err != nil {
			printErrorExit("Failed to get fx: ", err)
		}

		name, err := device.GetName()
		if err != nil {
			printErrorExit("Failed to get name: ", err)
		}
		fmt.Println("Name: ", name)

		typestr, err := device.GetType()
		if err != nil {
			printErrorExit("Failed to get type: ", err)
		}
		fmt.Println("  Type: ", typestr)

		serial, err := device.GetSerial()
		if err != nil {
			printErrorExit("Failed to get serial: ", err)
		}
		fmt.Println("  Serial: ", serial)

		fwversion, err := device.GetFirmwareVersion()
		if err != nil {
			printErrorExit("Failed to get firmware version: ", err)
		}
		fmt.Println("  Firmware Version: ", fwversion)

		if contains(features, "dpi") {
			dpi, err := device.GetDPI()
			if err != nil {
				printErrorExit("Failed to get DPI: ", err)
			}
			fmt.Println("  DPI: ", dpi)
			_, err = device.SetDPI(openrazer.DPI{X: 500, Y: 500})
			if err != nil {
				printErrorExit("Failed to set DPI: ", err)
			}
			maxdpi, err := device.GetMaxDPI()
			if err != nil {
				printErrorExit("Failed to get max DPI: ", err)
			}
			fmt.Println("  Max DPI: ", maxdpi)
		}

		if contains(features, "keyboard_layout") {
			layout, err := device.GetKeyboardLayout()
			if err != nil {
				printErrorExit("Failed to get keyboard layout: ", err)
			}
			fmt.Println("  Keyboard Layout: ", layout)
		}

		if contains(features, "poll_rate") {
			pollrate, err := device.GetPollRate()
			if err != nil {
				printErrorExit("Failed to get poll rate: ", err)
			}
			fmt.Println("  Poll rate: ", pollrate)
			_, err = device.SetPollRate(1000)
			if err != nil {
				printErrorExit("Failed to set poll rate: ", err)
			}
		}

		matrixdims, err := device.GetMatrixDimensions()
		if err != nil {
			printErrorExit("Failed to get matrix dimensions: ", err)
		}
		fmt.Println("  Matrix dimensions: ", matrixdims)

		leds := device.GetLeds()
		for _, led := range leds {
			ledId, err := led.GetLedId()
			if err != nil {
				printErrorExit("Failed to get led id: ", err)
			}
			fmt.Println("  Led: ", ledId)

			if contains(fx, "brightness") {
				brightness, err := led.GetBrightness()
				if err != nil {
					printErrorExit("Failed to get brightness: ", err)
				}
				fmt.Println("    Brightness: ", brightness)
				_, err = led.SetBrightness(50)
				if err != nil {
					printErrorExit("Failed to set brightness: ", err)
				}
			}
		}
	}
}
