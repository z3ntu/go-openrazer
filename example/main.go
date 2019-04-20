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

		serial, err := device.GetSerial()
		if err != nil {
			printErrorExit("Failed to get serial: ", err)
		}
		fmt.Println("Serial: ", serial)

		if contains(features, "dpi") {
			dpi, err := device.GetDPI()
			if err != nil {
				printErrorExit("Failed to get DPI: ", err)
			}
			fmt.Println("DPI: ", dpi)
		}

		leds := device.GetLeds()
		for _, led := range leds {
			if contains(fx, "brightness") {
				brightness, err := led.GetBrightness()
				if err != nil {
					printErrorExit("Failed to get brightness: ", err)
				}
				fmt.Println("Brightness: ", brightness)
			}
		}
	}
}
