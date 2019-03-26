package openrazer

import (
	"github.com/godbus/dbus"
)

type Manager struct {
	DbusConnection *dbus.Conn
	DbusObject     dbus.BusObject
	Devices        []*Device
}

const BusName = "io.github.openrazer1"
const ManagerPath = "/io/github/openrazer1"
const ManagerInterface = "io.github.openrazer1.Manager"

func NewManager() (*Manager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	dbusObj := conn.Object(BusName, ManagerPath)

	// Initialize the devices
	variant, err := dbusObj.GetProperty(ManagerInterface + ".Devices")
	if err != nil {
		return nil, err
	}
	var devices []*Device
	for _, path := range variant.Value().([]dbus.ObjectPath) {
		device, err := NewDevice(conn, path)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	mgr := Manager{
		DbusConnection: conn,
		DbusObject:     dbusObj,
		Devices:        devices,
	}

	return &mgr, nil
}

func (mgr *Manager) GetVersion() (string, error) {
	variant, err := mgr.DbusObject.GetProperty(ManagerInterface + ".Version")
	if err != nil {
		return "", err
	}
	return variant.String(), nil
}

func (mgr *Manager) GetDevices() []*Device {
	return mgr.Devices
}
