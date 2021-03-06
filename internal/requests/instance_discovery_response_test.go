// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package requests

import (
	"net/http"
	"reflect"
	"testing"
)

const (
	responseData = `{"tenant_discovery_response": "hello", "metadata":
				 [{"preferred_network": "hello", "preferred_cache": "hello", "tenant_discovery_endpoint": "hello"}]}`
)

func TestCreateInstanceDiscoveryResponse(t *testing.T) {
	expInstDisc := &InstanceDiscoveryResponse{
		TenantDiscoveryEndpoint: "hello",
		Metadata: []InstanceDiscoveryMetadata{
			{
				PreferredCache:          "hello",
				PreferredNetwork:        "hello",
				TenantDiscoveryEndpoint: "hello",
			},
		},
	}
	actualInstDisc, err := CreateInstanceDiscoveryResponse(createFakeResp(http.StatusOK, responseData))
	if err != nil {
		t.Errorf("Error should be nil, but it is %v", err)
	}
	if !reflect.DeepEqual(actualInstDisc.TenantDiscoveryEndpoint, expInstDisc.TenantDiscoveryEndpoint) &&
		!reflect.DeepEqual(actualInstDisc.Metadata, expInstDisc.Metadata) {
		t.Errorf("Actual instance discovery response %+v differs from expected instance discovery response %+v",
			actualInstDisc, expInstDisc)
	}
}
