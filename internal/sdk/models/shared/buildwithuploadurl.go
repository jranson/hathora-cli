// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/hathora/ci/internal/sdk/internal/utils"
	"time"
)

type BuildWithUploadURLRegionalContainerTags struct {
	ContainerTag string `json:"containerTag"`
	Region       Region `json:"region"`
}

func (o *BuildWithUploadURLRegionalContainerTags) GetContainerTag() string {
	if o == nil {
		return ""
	}
	return o.ContainerTag
}

func (o *BuildWithUploadURLRegionalContainerTags) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

type UploadBodyParams struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

func (o *UploadBodyParams) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *UploadBodyParams) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

// BuildWithUploadURL - A build represents a game server artifact and its associated metadata.
type BuildWithUploadURL struct {
	BuildTag *string `json:"buildTag,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RegionalContainerTags []BuildWithUploadURLRegionalContainerTags `json:"regionalContainerTags"`
	// The size (in bytes) of the Docker image built by Hathora.
	ImageSize int64       `json:"imageSize"`
	Status    BuildStatus `json:"status"`
	// When the build was deleted.
	DeletedAt *time.Time `json:"deletedAt"`
	// When [`RunBuild()`](https://hathora.dev/api#tag/BuildV2/operation/RunBuild) finished executing.
	FinishedAt *time.Time `json:"finishedAt"`
	// When [`RunBuild()`](https://hathora.dev/api#tag/BuildV2/operation/RunBuild) is called.
	StartedAt *time.Time `json:"startedAt"`
	// When [`CreateBuild()`](https://hathora.dev/api#tag/BuildV2/operation/CreateBuild) is called.
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	// System generated id for a build. Increments by 1.
	BuildID int `json:"buildId"`
	// System generated unique identifier for an application.
	AppID            string             `json:"appId"`
	UploadBodyParams []UploadBodyParams `json:"uploadBodyParams"`
	UploadURL        string             `json:"uploadUrl"`
}

func (b BuildWithUploadURL) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BuildWithUploadURL) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BuildWithUploadURL) GetBuildTag() *string {
	if o == nil {
		return nil
	}
	return o.BuildTag
}

func (o *BuildWithUploadURL) GetRegionalContainerTags() []BuildWithUploadURLRegionalContainerTags {
	if o == nil {
		return []BuildWithUploadURLRegionalContainerTags{}
	}
	return o.RegionalContainerTags
}

func (o *BuildWithUploadURL) GetImageSize() int64 {
	if o == nil {
		return 0
	}
	return o.ImageSize
}

func (o *BuildWithUploadURL) GetStatus() BuildStatus {
	if o == nil {
		return BuildStatus("")
	}
	return o.Status
}

func (o *BuildWithUploadURL) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *BuildWithUploadURL) GetFinishedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.FinishedAt
}

func (o *BuildWithUploadURL) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *BuildWithUploadURL) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *BuildWithUploadURL) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *BuildWithUploadURL) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *BuildWithUploadURL) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *BuildWithUploadURL) GetUploadBodyParams() []UploadBodyParams {
	if o == nil {
		return []UploadBodyParams{}
	}
	return o.UploadBodyParams
}

func (o *BuildWithUploadURL) GetUploadURL() string {
	if o == nil {
		return ""
	}
	return o.UploadURL
}
