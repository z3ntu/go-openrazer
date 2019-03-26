package openrazer

import "github.com/godbus/dbus"

type Device struct {
	DbusConnection *dbus.Conn
	DbusObject     dbus.BusObject
}

const DeviceInterface = "io.github.openrazer1.Device"

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (*Device, error) {
	dev := Device{
		DbusConnection: conn,
	}
	dev.DbusObject = conn.Object(BusName, path)

	return &dev, nil
}

func (dev *Device) GetSerial() (string, error) {
	var serial string
	err := dev.DbusObject.Call(DeviceInterface+".getSerial", 0).Store(&serial)
	if err != nil {
		return "", err
	}
	return serial, nil
}

func (dev *Device) GetLeds() ([]*Led, error) {
	variant, err := dev.DbusObject.GetProperty(DeviceInterface + ".Leds")
	if err != nil {
		return nil, err
	}
	var leds []*Led
	for _, path := range variant.Value().([]dbus.ObjectPath) {
		led, err := NewLed(dev.DbusConnection, path)
		if err != nil {
			return nil, err
		}
		leds = append(leds, led)
	}
	return leds, nil
}
