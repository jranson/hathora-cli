// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// TransportType - Transport type specifies the underlying communication protocol to the exposed port.
type TransportType string

const (
	TransportTypeTCP TransportType = "tcp"
	TransportTypeUDP TransportType = "udp"
	TransportTypeTLS TransportType = "tls"
)

func (e TransportType) ToPointer() *TransportType {
	return &e
}
