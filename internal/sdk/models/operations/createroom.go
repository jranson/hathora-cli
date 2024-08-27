// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type CreateRoomGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateRoomGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateRoomRequest struct {
	AppID            *string                 `pathParam:"style=simple,explode=false,name=appId"`
	RoomID           *string                 `queryParam:"style=form,explode=true,name=roomId"`
	CreateRoomParams shared.CreateRoomParams `request:"mediaType=application/json"`
}

func (o *CreateRoomRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateRoomRequest) GetRoomID() *string {
	if o == nil {
		return nil
	}
	return o.RoomID
}

func (o *CreateRoomRequest) GetCreateRoomParams() shared.CreateRoomParams {
	if o == nil {
		return shared.CreateRoomParams{}
	}
	return o.CreateRoomParams
}

type CreateRoomResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse        *http.Response
	RoomConnectionData *shared.RoomConnectionData
}

func (o *CreateRoomResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRoomResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRoomResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRoomResponse) GetRoomConnectionData() *shared.RoomConnectionData {
	if o == nil {
		return nil
	}
	return o.RoomConnectionData
}
