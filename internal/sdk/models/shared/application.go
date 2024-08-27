// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/hathora/ci/internal/sdk/internal/utils"
	"time"
)

// Application - An application object is the top level namespace for the game server.
type Application struct {
	// The email address or token id for the user that deleted the application.
	DeletedBy *string `json:"deletedBy"`
	// When the application was deleted.
	DeletedAt *time.Time `json:"deletedAt"`
	// When the application was created.
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	// System generated unique identifier for an organization. Not guaranteed to have a specific format.
	OrgID string `json:"orgId"`
	// Configure [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service) for your application. Use Hathora's built-in auth providers or use your own [custom authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service#custom-auth-provider).
	AuthConfiguration AuthConfiguration `json:"authConfiguration"`
	// Secret that is used for identity and access management.
	AppSecret string `json:"appSecret"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
	// Readable name for an application. Must be unique within an organization.
	AppName string `json:"appName"`
}

func (a Application) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Application) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Application) GetDeletedBy() *string {
	if o == nil {
		return nil
	}
	return o.DeletedBy
}

func (o *Application) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Application) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Application) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *Application) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *Application) GetAuthConfiguration() AuthConfiguration {
	if o == nil {
		return AuthConfiguration{}
	}
	return o.AuthConfiguration
}

func (o *Application) GetAppSecret() string {
	if o == nil {
		return ""
	}
	return o.AppSecret
}

func (o *Application) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *Application) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}
