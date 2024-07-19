// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ActiveConnectionInfoStatus string

const (
	ActiveConnectionInfoStatusActive ActiveConnectionInfoStatus = "active"
)

func (e ActiveConnectionInfoStatus) ToPointer() *ActiveConnectionInfoStatus {
	return &e
}

type ActiveConnectionInfo struct {
	Status ActiveConnectionInfoStatus `json:"status"`
	// Transport type specifies the underlying communication protocol to the exposed port.
	TransportType TransportType `json:"transportType"`
	Port          float64       `json:"port"`
	Host          string        `json:"host"`
	// Unique identifier to a game session or match. Use the default system generated ID or overwrite it with your own.
	// Note: error will be returned if `roomId` is not globally unique.
	RoomID string `json:"roomId"`
}

func (o *ActiveConnectionInfo) GetStatus() ActiveConnectionInfoStatus {
	if o == nil {
		return ActiveConnectionInfoStatus("")
	}
	return o.Status
}

func (o *ActiveConnectionInfo) GetTransportType() TransportType {
	if o == nil {
		return TransportType("")
	}
	return o.TransportType
}

func (o *ActiveConnectionInfo) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *ActiveConnectionInfo) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *ActiveConnectionInfo) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}
