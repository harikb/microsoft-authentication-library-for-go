// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package msalbase

// StorageTokenResponse mimics a token response that was pulled from the cache.
type StorageTokenResponse struct {
	RefreshToken Credential
	AccessToken  accessTokenProvider
	IDToken      Credential
	account      Account
}

// CreateStorageTokenResponse creates a token response from cache.
func CreateStorageTokenResponse(accessToken accessTokenProvider, refreshToken Credential, idToken Credential, account Account) StorageTokenResponse {
	return StorageTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IDToken:      idToken,
		account:      account,
	}
}
