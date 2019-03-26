package main

import (
	"fmt"
	"github.com/z3ntu/go-openrazer"
	"os"
)

func main() {
	manager, err := openrazer.NewManager()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create manager: ", err)
		os.Exit(1)
	}

	version, err := manager.GetVersion()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get version: ", err)
		os.Exit(1)
	}
	fmt.Println("Version: ", version)

	devices, err := manager.GetDevices()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get devices: ", err)
		os.Exit(1)
	}
	fmt.Println("Devices: ", devices)
}
