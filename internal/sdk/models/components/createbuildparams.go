// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateBuildParams struct {
	// Tag to associate an external version with a build. It is accessible via [`GetBuildInfo()`](https://hathora.dev/api#tag/BuildV2/operation/GetBuildInfo).
	BuildTag *string `json:"buildTag,omitempty"`
}

func (o *CreateBuildParams) GetBuildTag() *string {
	if o == nil {
		return nil
	}
	return o.BuildTag
}
