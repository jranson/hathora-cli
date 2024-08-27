// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetBuildsGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetBuildsGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetBuildsRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetBuildsRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetBuildsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	Builds []shared.Build
}

func (o *GetBuildsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBuildsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBuildsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetBuildsResponse) GetBuilds() []shared.Build {
	if o == nil {
		return nil
	}
	return o.Builds
}
