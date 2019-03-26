package openrazer

import (
	"github.com/godbus/dbus"
)

type Manager struct {
	DbusConnection *dbus.Conn
	DbusObject     dbus.BusObject
}

func NewManager() (*Manager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	mgr := Manager{
		DbusConnection: conn,
	}
	mgr.DbusObject = conn.Object("io.github.openrazer1", "/io/github/openrazer1")

	return &mgr, nil

}

func (mgr *Manager) GetVersion() (string, error) {
	variant, err := mgr.DbusObject.GetProperty("io.github.openrazer1.Manager.Version")
	if err != nil {
		return "", err
	}
	return variant.String(), nil
}

func (mgr *Manager) GetDevices() ([]dbus.ObjectPath, error) {
	variant, err := mgr.DbusObject.GetProperty("io.github.openrazer1.Manager.Devices")
	if err != nil {
		return nil, err
	}
	return variant.Value().([]dbus.ObjectPath), nil
}
