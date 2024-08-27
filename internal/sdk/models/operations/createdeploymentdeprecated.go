// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type CreateDeploymentDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateDeploymentDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateDeploymentDeprecatedRequest struct {
	AppID            *string                 `pathParam:"style=simple,explode=false,name=appId"`
	BuildID          int                     `pathParam:"style=simple,explode=false,name=buildId"`
	DeploymentConfig shared.DeploymentConfig `request:"mediaType=application/json"`
}

func (o *CreateDeploymentDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateDeploymentDeprecatedRequest) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *CreateDeploymentDeprecatedRequest) GetDeploymentConfig() shared.DeploymentConfig {
	if o == nil {
		return shared.DeploymentConfig{}
	}
	return o.DeploymentConfig
}

type CreateDeploymentDeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Deployment  *shared.Deployment
}

func (o *CreateDeploymentDeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDeploymentDeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDeploymentDeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDeploymentDeprecatedResponse) GetDeployment() *shared.Deployment {
	if o == nil {
		return nil
	}
	return o.Deployment
}
