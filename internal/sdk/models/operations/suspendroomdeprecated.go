// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type SuspendRoomDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *SuspendRoomDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type SuspendRoomDeprecatedRequest struct {
	AppID  *string `pathParam:"style=simple,explode=false,name=appId"`
	RoomID string  `pathParam:"style=simple,explode=false,name=roomId"`
}

func (o *SuspendRoomDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *SuspendRoomDeprecatedRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}
