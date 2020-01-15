package openrazer

import (
	"github.com/godbus/dbus/v5"
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

func (dev *Device) GetFirmwareVersion() (string, error) {
	var version string
	err := dev.DbusObject.Call(DeviceInterface+".getFirmwareVersion", 0).Store(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}

func (dev *Device) GetKeyboardLayout() (string, error) {
	var layout string
	err := dev.DbusObject.Call(DeviceInterface+".getKeyboardLayout", 0).Store(&layout)
	if err != nil {
		return "", err
	}
	return layout, nil
}

func (dev *Device) GetDPI() (DPI, error) {
	var dpi DPI
	err := dev.DbusObject.Call(DeviceInterface+".getDPI", 0).Store(&dpi)
	if err != nil {
		return DPI{}, err
	}
	return dpi, nil
}

func (dev *Device) SetDPI(dpi DPI) (bool, error) {
	var ret bool
	err := dev.DbusObject.Call(DeviceInterface+".setDPI", 0, dpi).Store(&ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func (dev *Device) GetMaxDPI() (uint16, error) {
	var maxDPI uint16
	err := dev.DbusObject.Call(DeviceInterface+".getMaxDPI", 0).Store(&maxDPI)
	if err != nil {
		return 0, err
	}
	return maxDPI, nil
}

func (dev *Device) GetPollRate() (uint16, error) {
	var pollRate uint16
	err := dev.DbusObject.Call(DeviceInterface+".getPollRate", 0).Store(&pollRate)
	if err != nil {
		return 0, err
	}
	return pollRate, nil
}

func (dev *Device) SetPollRate(pollRate uint16) (bool, error) {
	var ret bool
	err := dev.DbusObject.Call(DeviceInterface+".setPollRate", 0, pollRate).Store(&ret)
	if err != nil {
		return false, err
	}
	return ret, nil
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

func (dev *Device) GetName() (string, error) {
	variant, err := dev.DbusObject.GetProperty(DeviceInterface + ".Name")
	if err != nil {
		return "", err
	}
	return variant.String(), nil
}

func (dev *Device) GetType() (string, error) {
	variant, err := dev.DbusObject.GetProperty(DeviceInterface + ".Type")
	if err != nil {
		return "", err
	}
	return variant.String(), nil
}

func (dev *Device) GetMatrixDimensions() (MatrixDimensions, error) {
	variant, err := dev.DbusObject.GetProperty(DeviceInterface + ".MatrixDimensions")
	if err != nil {
		return MatrixDimensions{}, err
	}
	values := variant.Value().([]interface{})
	return MatrixDimensions{values[0].(uint8), values[1].(uint8)}, nil
}

func (dev *Device) GetLeds() []*Led {
	return dev.Leds
}
