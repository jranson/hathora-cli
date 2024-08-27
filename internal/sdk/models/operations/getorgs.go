// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetOrgsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	OrgsPage *shared.OrgsPage
}

func (o *GetOrgsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOrgsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOrgsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetOrgsResponse) GetOrgsPage() *shared.OrgsPage {
	if o == nil {
		return nil
	}
	return o.OrgsPage
}
