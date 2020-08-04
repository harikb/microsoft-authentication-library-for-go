// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package msalgo

import (
	"net/url"

	"github.com/AzureAD/microsoft-authentication-library-for-go/src/internal/msalbase"
	"github.com/AzureAD/microsoft-authentication-library-for-go/src/internal/requests"
)

//AuthorizationCodeURLParameters has the parameters to create the URL to generate an authorization code
type AuthorizationCodeURLParameters struct {
	ClientID            string
	RedirectURI         string
	ResponseType        string
	ResponseMode        string
	State               string
	Prompt              string
	LoginHint           string
	DomainHint          string
	CodeChallenge       string
	CodeChallengeMethod string
	Scopes              []string
}

//CreateAuthorizationCodeURLParameters creates an AuthorizationCodeURLParameters instance
func CreateAuthorizationCodeURLParameters(clientID string, redirectURI string, scopes []string, codeChallenge string) *AuthorizationCodeURLParameters {
	p := &AuthorizationCodeURLParameters{
		ClientID:      clientID,
		ResponseType:  msalbase.DefaultAuthCodeResponseType,
		RedirectURI:   redirectURI,
		Scopes:        scopes,
		CodeChallenge: codeChallenge,
	}
	return p
}

//CreateURL creates the URL required to generate an authorization code from the parameters
func (p *AuthorizationCodeURLParameters) CreateURL(wrm requests.WebRequestManager, authParams *msalbase.AuthParametersInternal) (string, error) {
	resolutionManager := requests.CreateAuthorityEndpointResolutionManager(wrm)
	endpoints, err := resolutionManager.ResolveEndpoints(authParams.AuthorityInfo, "")
	if err != nil {
		return "", err
	}
	baseURL, err := url.Parse(endpoints.AuthorizationEndpoint)
	if err != nil {
		return "", err
	}
	urlParams := url.Values{}
	urlParams.Add("client_id", p.ClientID)
	urlParams.Add("response_type", p.ResponseType)
	urlParams.Add("redirect_uri", p.RedirectURI)
	urlParams.Add("scope", p.getSeparatedScopes())
	urlParams.Add("code_challenge", p.CodeChallenge)
	if p.State != "" {
		urlParams.Add("state", p.State)
	}
	if p.ResponseMode != "" {
		urlParams.Add("response_mode", p.ResponseMode)
	}
	if p.Prompt != "" {
		urlParams.Add("prompt", p.Prompt)
	}
	if p.LoginHint != "" {
		urlParams.Add("login_hint", p.LoginHint)
	}
	if p.DomainHint != "" {
		urlParams.Add("domain_hint", p.DomainHint)
	}
	if p.CodeChallengeMethod != "" {
		urlParams.Add("code_challenge_method", p.CodeChallengeMethod)
	}
	baseURL.RawQuery = urlParams.Encode()
	return baseURL.String(), nil
}

func (p *AuthorizationCodeURLParameters) getSeparatedScopes() string {
	return msalbase.ConcatenateScopes(p.Scopes)
}
