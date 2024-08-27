// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/hathora/ci/internal/sdk/internal/globals"
	"github.com/hathora/ci/internal/sdk/internal/hooks"
	"github.com/hathora/ci/internal/sdk/internal/utils"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"github.com/hathora/ci/internal/sdk/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.hathora.dev",
	"https:///",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	Globals           globals.Globals
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// SDK - Hathora Cloud API: Welcome to the Hathora Cloud API documentation! Learn how to use the Hathora Cloud APIs to build and scale your game servers globally.
type SDK struct {
	//
	TokensV1 *TokensV1
	RoomsV1  *RoomsV1
	RoomsV2  *RoomsV2
	// Deprecated. Use [ProcessesV2](https://hathora.dev/api#tag/ProcessesV2).
	ProcessesV1 *ProcessesV1
	// Operations to get data on active and stopped [processes](https://hathora.dev/docs/concepts/hathora-entities#process).
	ProcessesV2     *ProcessesV2
	OrganizationsV1 *OrganizationsV1
	// Operations to get metrics by [process](https://hathora.dev/docs/concepts/hathora-entities#process). We store 72 hours of metrics data.
	MetricsV1 *MetricsV1
	//
	ManagementV1 *ManagementV1
	LogsV1       *LogsV1
	LobbiesV1    *LobbiesV1
	LobbiesV2    *LobbiesV2
	LobbiesV3    *LobbiesV3
	// Deprecated. Does not include latest Regions (missing Dallas region). Use [DiscoveryV2](https://hathora.dev/api#tag/DiscoveryV2).
	DiscoveryV1 *DiscoveryV1
	// Service that allows clients to directly ping all Hathora regions to get latency information
	DiscoveryV2   *DiscoveryV2
	DeploymentsV1 *DeploymentsV1
	DeploymentsV2 *DeploymentsV2
	BuildsV1      *BuildsV1
	BuildsV2      *BuildsV2
	//
	BillingV1 *BillingV1
	// Operations that allow you to generate a Hathora-signed [JSON web token (JWT)](https://jwt.io/) for [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service).
	AuthV1 *AuthV1
	AppsV1 *AppsV1

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

// WithAppID allows setting the AppID parameter for all supported operations
func WithAppID(appID string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Globals.AppID = &appID
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "0.4.0",
			GenVersion:        "2.380.2",
			UserAgent:         "speakeasy-sdk/go 0.4.0 2.380.2 0.0.1 github.com/hathora/ci/internal/sdk",
			Globals:           globals.Globals{},
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.TokensV1 = newTokensV1(sdk.sdkConfiguration)

	sdk.RoomsV1 = newRoomsV1(sdk.sdkConfiguration)

	sdk.RoomsV2 = newRoomsV2(sdk.sdkConfiguration)

	sdk.ProcessesV1 = newProcessesV1(sdk.sdkConfiguration)

	sdk.ProcessesV2 = newProcessesV2(sdk.sdkConfiguration)

	sdk.OrganizationsV1 = newOrganizationsV1(sdk.sdkConfiguration)

	sdk.MetricsV1 = newMetricsV1(sdk.sdkConfiguration)

	sdk.ManagementV1 = newManagementV1(sdk.sdkConfiguration)

	sdk.LogsV1 = newLogsV1(sdk.sdkConfiguration)

	sdk.LobbiesV1 = newLobbiesV1(sdk.sdkConfiguration)

	sdk.LobbiesV2 = newLobbiesV2(sdk.sdkConfiguration)

	sdk.LobbiesV3 = newLobbiesV3(sdk.sdkConfiguration)

	sdk.DiscoveryV1 = newDiscoveryV1(sdk.sdkConfiguration)

	sdk.DiscoveryV2 = newDiscoveryV2(sdk.sdkConfiguration)

	sdk.DeploymentsV1 = newDeploymentsV1(sdk.sdkConfiguration)

	sdk.DeploymentsV2 = newDeploymentsV2(sdk.sdkConfiguration)

	sdk.BuildsV1 = newBuildsV1(sdk.sdkConfiguration)

	sdk.BuildsV2 = newBuildsV2(sdk.sdkConfiguration)

	sdk.BillingV1 = newBillingV1(sdk.sdkConfiguration)

	sdk.AuthV1 = newAuthV1(sdk.sdkConfiguration)

	sdk.AppsV1 = newAppsV1(sdk.sdkConfiguration)

	return sdk
}
