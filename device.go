package openrazer

import (
	"github.com/godbus/dbus"
)

type Device struct {
	DbusConnection *dbus.Conn
	DbusObject     dbus.BusObject
	Leds           []*Led
}

const DeviceInterface = "io.github.openrazer1.Device"

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (*Device, error) {
	dbusObj := conn.Object(BusName, path)

	variant, err := dbusObj.GetProperty(DeviceInterface + ".Leds")
	if err != nil {
		return nil, err
	}
	var leds []*Led
	for _, path := range variant.Value().([]dbus.ObjectPath) {
		led, err := NewLed(conn, path)
		if err != nil {
			return nil, err
		}
		leds = append(leds, led)
	}

	dev := Device{
		DbusConnection: conn,
		DbusObject:     dbusObj,
		Leds:           leds,
	}

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

func (dev *Device) GetDPI() (DPI, error) {
	var dpi DPI
	err := dev.DbusObject.Call(DeviceInterface+".getDPI", 0).Store(&dpi)
	if err != nil {
		return DPI{}, err
	}
	return dpi, nil
}

func (dev *Device) GetSupportedFeatures() ([]string, error) {
	variant, err := dev.DbusObject.GetProperty(DeviceInterface + ".SupportedFeatures")
	if err != nil {
		return nil, err
	}
	return variant.Value().([]string), nil
}

func (dev *Device) GetSupportedFx() ([]string, error) {
	variant, err := dev.DbusObject.GetProperty(DeviceInterface + ".SupportedFx")
	if err != nil {
		return nil, err
	}
	return variant.Value().([]string), nil
}

func (dev *Device) GetLeds() []*Led {
	return dev.Leds
}
