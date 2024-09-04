// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetProcessInfoV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetProcessInfoV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetProcessInfoV2DeprecatedRequest struct {
	AppID     *string `pathParam:"style=simple,explode=false,name=appId"`
	ProcessID string  `pathParam:"style=simple,explode=false,name=processId"`
}

func (o *GetProcessInfoV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetProcessInfoV2DeprecatedRequest) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}

type GetProcessInfoV2DeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	ProcessV2 *shared.ProcessV2
}

func (o *GetProcessInfoV2DeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProcessInfoV2DeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProcessInfoV2DeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetProcessInfoV2DeprecatedResponse) GetProcessV2() *shared.ProcessV2 {
	if o == nil {
		return nil
	}
	return o.ProcessV2
}
