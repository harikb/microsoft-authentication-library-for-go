// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package requests

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/AzureAD/microsoft-authentication-library-for-go/internal/msalbase"
)

// TODO(jdoak): Replace these tests with subtests to eliminate these globals
// or table driven.

var (
	testUPAuthorityInfo msalbase.AuthorityInfo
	uprTestAuthParams   msalbase.AuthParametersInternal

	upWRM               = new(MockWebRequestManager)
	usernamePassRequest *UsernamePasswordRequest

	managedUserRealm = msalbase.UserRealm{
		AccountType: "Managed",
	}
	errorUserRealm = msalbase.UserRealm{
		AccountType: "",
	}
	federatedUserRealm = msalbase.UserRealm{
		AccountType:           "Federated",
		FederationMetadataURL: "fedMetaURL",
	}
)

func init() {
	var err error
	testUPAuthorityInfo, err = msalbase.CreateAuthorityInfoFromAuthorityURI("https://login.microsoftonline.com/v2.0/", true)
	if err != nil {
		panic(err)
	}

	uprTestAuthParams = msalbase.CreateAuthParametersInternal("clientID", createTestAuthorityInfo())
	uprTestAuthParams.Endpoints = msalbase.CreateAuthorityEndpoints(
		"https://login.microsoftonline.com/v2.0/authorize",
		"https://login.microsoftonline.com/v2.0/token",
		"https://login.microsoftonline.com/v2.0",
		"login.microsoftonline.com",
	)

	usernamePassRequest = &UsernamePasswordRequest{
		authParameters: uprTestAuthParams,
	}
}

// TODO(msal expert): This test SEEMS borked. .Execute()
// calls wsTrustResp.GetSAMLAssertion(). Because mexDoc.UsernamePasswordEndpoint is
// set to wsEndpoint with wstrust.Trust2005, Response.GetSAMLAssertion() will error
// with non-supported. This is what should have been happening as far as I can tell.
// So this test should have never worked.  Trust13 is supported, but switching to it
// just causes it to give an EOL error, probably because it has to do some parsing.
// I don't know what that needs to be and would defer to experts to let me know.
/*
func TestUsernamePassExecuteWithFederated(t *testing.T) {
	upWRM = new(MockWebRequestManager)
	usernamePassRequest.webRequestManager = upWRM
	upWRM.On("GetTenantDiscoveryResponse",
		"https://login.microsoftonline.com/v2.0/v2.0/.well-known/openid-configuration").Return(tdr, nil)
	upWRM.On("GetUserRealm", usernamePassRequest.authParameters).Return(federatedUserRealm, nil)
	wsEndpoint := wstrust.Endpoint{EndpointVersion: wstrust.Trust2005, URL: "upEndpoint"}
	mexDoc := wstrust.MexDocument{
		UsernamePasswordEndpoint: wsEndpoint,
	}
	upWRM.On("GetMex", "fedMetaURL").Return(mexDoc, nil)
	wsTrustResp := wstrust.Response{}
	upWRM.On("GetWsTrustResponse", uprTestAuthParams, "", wsEndpoint).Return(wsTrustResp, nil)
	_, err := usernamePassRequest.Execute()
	if err != nil {
		t.Errorf("Error should be nil, but is %v", err)
	}
}
*/

func TestUsernamePassExecuteWithManaged(t *testing.T) {
	upWRM = new(MockWebRequestManager)
	usernamePassRequest.webRequestManager = upWRM
	upWRM.On("GetTenantDiscoveryResponse",
		"https://login.microsoftonline.com/v2.0/v2.0/.well-known/openid-configuration").Return(createTDR(), nil)
	upWRM.On("GetUserRealm", usernamePassRequest.authParameters).Return(managedUserRealm, nil)
	actualTokenResp := msalbase.TokenResponse{}
	upWRM.On("GetAccessTokenFromUsernamePassword", usernamePassRequest.authParameters).Return(actualTokenResp, nil)
	_, err := usernamePassRequest.Execute(context.Background())
	if err != nil {
		t.Errorf("Error is supposed to be nil, instead it is %v", err)
	}
}

func TestUsernamePassExecuteWithAcctError(t *testing.T) {
	newUpWRM := new(MockWebRequestManager)
	usernamePassRequest.webRequestManager = newUpWRM
	newUpWRM.On("GetTenantDiscoveryResponse",
		"https://login.microsoftonline.com/v2.0/v2.0/.well-known/openid-configuration").Return(createTDR(), nil)
	newUpWRM.On("GetUserRealm", usernamePassRequest.authParameters).Return(errorUserRealm, nil)
	_, acctError := usernamePassRequest.Execute(context.Background())
	expectedErrorMessage := "unknown account type"
	if acctError == nil {
		t.Errorf("Error is nil, should be %v", errors.New(expectedErrorMessage))
	}
	if !reflect.DeepEqual(acctError.Error(), expectedErrorMessage) {
		t.Errorf("Actual error message %v differs from expected error message %v", acctError.Error(), expectedErrorMessage)
	}
}
