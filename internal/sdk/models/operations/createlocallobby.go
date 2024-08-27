// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type CreateLocalLobbyGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateLocalLobbyGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateLocalLobbySecurity struct {
	PlayerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *CreateLocalLobbySecurity) GetPlayerAuth() string {
	if o == nil {
		return ""
	}
	return o.PlayerAuth
}

type CreateLocalLobbyRequestBody struct {
	// User input to initialize the game state. Object must be smaller than 64KB.
	InitialConfig any           `json:"initialConfig"`
	Region        shared.Region `json:"region"`
}

func (o *CreateLocalLobbyRequestBody) GetInitialConfig() any {
	if o == nil {
		return nil
	}
	return o.InitialConfig
}

func (o *CreateLocalLobbyRequestBody) GetRegion() shared.Region {
	if o == nil {
		return shared.Region("")
	}
	return o.Region
}

type CreateLocalLobbyRequest struct {
	AppID       *string                     `pathParam:"style=simple,explode=false,name=appId"`
	RoomID      *string                     `queryParam:"style=form,explode=true,name=roomId"`
	RequestBody CreateLocalLobbyRequestBody `request:"mediaType=application/json"`
}

func (o *CreateLocalLobbyRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateLocalLobbyRequest) GetRoomID() *string {
	if o == nil {
		return nil
	}
	return o.RoomID
}

func (o *CreateLocalLobbyRequest) GetRequestBody() CreateLocalLobbyRequestBody {
	if o == nil {
		return CreateLocalLobbyRequestBody{}
	}
	return o.RequestBody
}

type CreateLocalLobbyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Lobby       *shared.Lobby
}

func (o *CreateLocalLobbyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLocalLobbyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLocalLobbyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLocalLobbyResponse) GetLobby() *shared.Lobby {
	if o == nil {
		return nil
	}
	return o.Lobby
}
