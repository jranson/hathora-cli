// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetDeploymentsV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetDeploymentsV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetDeploymentsV2DeprecatedRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetDeploymentsV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}
