package session

import (
	"errors"
	"net"
)

type DeviceMetadata struct {
	userAgent  *string
	ipAddress  *net.IP
	deviceName *string
}

func NewDeviceMetadata(userAgent, ipAddress, deviceName *string) (*DeviceMetadata, error) {
	var validIPAddressPtr *net.IP = nil
	if ipAddress != nil {
		validIPAddress := net.ParseIP(*ipAddress)
		if validIPAddress == nil {
			return nil, errors.New("invalid IP Address")
		}
		validIPAddressPtr = &validIPAddress
	}

	return &DeviceMetadata{
		userAgent:  userAgent,
		ipAddress:  validIPAddressPtr,
		deviceName: deviceName,
	}, nil
}

func (metadata DeviceMetadata) UserAgent() *string {
	return metadata.userAgent
}

func (metadata DeviceMetadata) IPAddress() *net.IP {
	return metadata.ipAddress
}

func (metadata DeviceMetadata) DeviceName() *string {
	return metadata.deviceName
}
