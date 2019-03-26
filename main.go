package main

import (
	"fmt"
	"github.com/godbus/dbus"
	"os"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to system bus:", err)
		os.Exit(1)
	}

	obj := conn.Object("io.github.openrazer1", "/io/github/openrazer1")
	version, err := obj.GetProperty("io.github.openrazer1.Manager.Version");
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get version: ", err)
		os.Exit(1)
	}
	fmt.Println("Version: ", version)


	devices, err := obj.GetProperty("io.github.openrazer1.Manager.Devices");
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get devices: ", err)
		os.Exit(1)
	}
	fmt.Println("Devices: ", devices)
}
