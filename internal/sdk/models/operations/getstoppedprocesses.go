// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetStoppedProcessesGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetStoppedProcessesGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetStoppedProcessesRequest struct {
	AppID  *string        `pathParam:"style=simple,explode=false,name=appId"`
	Region *shared.Region `queryParam:"style=form,explode=true,name=region"`
}

func (o *GetStoppedProcessesRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetStoppedProcessesRequest) GetRegion() *shared.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

type GetStoppedProcessesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	Processes []shared.Process
}

func (o *GetStoppedProcessesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetStoppedProcessesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetStoppedProcessesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetStoppedProcessesResponse) GetProcesses() []shared.Process {
	if o == nil {
		return nil
	}
	return o.Processes
}
