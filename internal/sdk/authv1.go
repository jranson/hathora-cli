// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/hathora/ci/internal/sdk/internal/hooks"
	"github.com/hathora/ci/internal/sdk/internal/utils"
	"github.com/hathora/ci/internal/sdk/models/operations"
	"github.com/hathora/ci/internal/sdk/models/sdkerrors"
	"github.com/hathora/ci/internal/sdk/models/shared"
	"io"
	"net/http"
)

// AuthV1 - Operations that allow you to generate a Hathora-signed [JSON web token (JWT)](https://jwt.io/) for [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service).
type AuthV1 struct {
	sdkConfiguration sdkConfiguration
}

func newAuthV1(sdkConfig sdkConfiguration) *AuthV1 {
	return &AuthV1{
		sdkConfiguration: sdkConfig,
	}
}

// LoginAnonymous - Returns a unique player token for an anonymous user.
func (s *AuthV1) LoginAnonymous(ctx context.Context, appID *string, opts ...operations.Option) (*operations.LoginAnonymousResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "LoginAnonymous",
		SecuritySource: nil,
	}

	request := operations.LoginAnonymousRequest{
		AppID: appID,
	}

	globals := operations.LoginAnonymousGlobals{
		AppID: s.sdkConfiguration.Globals.AppID,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/auth/v1/{appId}/login/anonymous", request, globals)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request, globals)

	if err := utils.PopulateQueryParams(ctx, req, request, globals); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"404", "429", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.LoginAnonymousResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.PlayerTokenObject
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PlayerTokenObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.APIError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// LoginNickname - Returns a unique player token with a specified nickname for a user.
func (s *AuthV1) LoginNickname(ctx context.Context, nicknameObject shared.NicknameObject, appID *string, opts ...operations.Option) (*operations.LoginNicknameResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "LoginNickname",
		SecuritySource: nil,
	}

	request := operations.LoginNicknameRequest{
		AppID:          appID,
		NicknameObject: nicknameObject,
	}

	globals := operations.LoginNicknameGlobals{
		AppID: s.sdkConfiguration.Globals.AppID,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/auth/v1/{appId}/login/nickname", request, globals)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "NicknameObject", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request, globals)

	if err := utils.PopulateQueryParams(ctx, req, request, globals); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"404", "429", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.LoginNicknameResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.PlayerTokenObject
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PlayerTokenObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.APIError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// LoginGoogle - Returns a unique player token using a Google-signed OIDC `idToken`.
func (s *AuthV1) LoginGoogle(ctx context.Context, googleIDTokenObject shared.GoogleIDTokenObject, appID *string, opts ...operations.Option) (*operations.LoginGoogleResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "LoginGoogle",
		SecuritySource: nil,
	}

	request := operations.LoginGoogleRequest{
		AppID:               appID,
		GoogleIDTokenObject: googleIDTokenObject,
	}

	globals := operations.LoginGoogleGlobals{
		AppID: s.sdkConfiguration.Globals.AppID,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/auth/v1/{appId}/login/google", request, globals)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "GoogleIDTokenObject", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request, globals)

	if err := utils.PopulateQueryParams(ctx, req, request, globals); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				return nil, backoff.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"401", "404", "429", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.LoginGoogleResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.PlayerTokenObject
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.PlayerTokenObject = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.APIError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
