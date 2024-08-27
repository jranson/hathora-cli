// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DestroyRoomDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DestroyRoomDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type DestroyRoomDeprecatedRequest struct {
	AppID  *string `pathParam:"style=simple,explode=false,name=appId"`
	RoomID string  `pathParam:"style=simple,explode=false,name=roomId"`
}

func (o *DestroyRoomDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *DestroyRoomDeprecatedRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

type DestroyRoomDeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DestroyRoomDeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DestroyRoomDeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DestroyRoomDeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
