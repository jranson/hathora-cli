// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type LoginGoogleGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *LoginGoogleGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type LoginGoogleRequest struct {
	AppID               *string                    `pathParam:"style=simple,explode=false,name=appId"`
	GoogleIDTokenObject shared.GoogleIDTokenObject `request:"mediaType=application/json"`
}

func (o *LoginGoogleRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *LoginGoogleRequest) GetGoogleIDTokenObject() shared.GoogleIDTokenObject {
	if o == nil {
		return shared.GoogleIDTokenObject{}
	}
	return o.GoogleIDTokenObject
}

type LoginGoogleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	PlayerTokenObject *shared.PlayerTokenObject
}

func (o *LoginGoogleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LoginGoogleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LoginGoogleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LoginGoogleResponse) GetPlayerTokenObject() *shared.PlayerTokenObject {
	if o == nil {
		return nil
	}
	return o.PlayerTokenObject
}
