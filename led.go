package openrazer

import "github.com/godbus/dbus/v5"

type Led struct {
	DbusObject dbus.BusObject
}

const LedInterface = "io.github.openrazer1.Led"

func NewLed(conn *dbus.Conn, path dbus.ObjectPath) (*Led, error) {
	led := Led{}
	led.DbusObject = conn.Object(BusName, path)

	return &led, nil
}

func (led *Led) GetBrightness() (int, error) {
	var brightness int
	err := led.DbusObject.Call(LedInterface+".getBrightness", 0).Store(&brightness)
	if err != nil {
		return 0, err
	}
	return brightness, nil
}
