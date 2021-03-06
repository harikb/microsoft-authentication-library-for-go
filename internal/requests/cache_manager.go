// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package requests

import (
	"context"

	"github.com/AzureAD/microsoft-authentication-library-for-go/internal/msalbase"
)

// CacheManager is the interface for the handling of caching operations
// TODO(jdoak): Remove this.
type CacheManager interface {
	TryReadCache(ctx context.Context, authParameters msalbase.AuthParametersInternal, webRequestManager WebRequestManager) (msalbase.StorageTokenResponse, error)
	CacheTokenResponse(authParameters msalbase.AuthParametersInternal, tokenResponse msalbase.TokenResponse) (msalbase.Account, error)
	// DeleteCachedRefreshToken(authParameters msalbase.AuthParametersInternal) error
	GetAllAccounts() ([]msalbase.Account, error)
	Serialize() (string, error)
	Deserialize(data []byte) error
}
