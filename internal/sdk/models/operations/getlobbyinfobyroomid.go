// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetLobbyInfoByRoomIDGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLobbyInfoByRoomIDGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLobbyInfoByRoomIDRequest struct {
	AppID  *string `pathParam:"style=simple,explode=false,name=appId"`
	RoomID string  `pathParam:"style=simple,explode=false,name=roomId"`
}

func (o *GetLobbyInfoByRoomIDRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetLobbyInfoByRoomIDRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}
