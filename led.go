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

func (led *Led) GetLedId() (int32, error) {
	variant, err := led.DbusObject.GetProperty(LedInterface + ".LedId")
	if err != nil {
		return 0, err
	}
	values := variant.Value().([]interface{})
	return values[0].(int32), nil
}

func (led *Led) GetBrightness() (uint8, error) {
	var brightness uint8
	err := led.DbusObject.Call(LedInterface+".getBrightness", 0).Store(&brightness)
	if err != nil {
		return 0, err
	}
	return brightness, nil
}

func (led *Led) SetBrightness(brightness uint8) (bool, error) {
	var ret bool
	err := led.DbusObject.Call(LedInterface+".setBrightness", 0, brightness).Store(&ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}
