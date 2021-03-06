// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package requests

import (
	"context"

	"github.com/AzureAD/microsoft-authentication-library-for-go/internal/msalbase"
)

// TokenRequester is an interface that handles all token acquisition requests
type TokenRequester interface {
	Execute(context.Context) (msalbase.TokenResponse, error)
}
