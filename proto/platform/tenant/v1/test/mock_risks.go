/*
Copyright 2023 Chainguard, Inc.
SPDX-License-Identifier: Apache-2.0
*/

package test

import (
	"context"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/testing/protocmp"

	tenant "chainguard.dev/sdk/proto/platform/tenant/v1"
)

var _ tenant.RisksClient = (*MockRisksClient)(nil)

type MockRisksClient struct {
	OnList []RisksOnList
}

type RisksOnList struct {
	Given *tenant.RiskFilter
	List  *tenant.RiskList
	Error error
}

func (m MockRisksClient) List(_ context.Context, given *tenant.RiskFilter, _ ...grpc.CallOption) (*tenant.RiskList, error) {
	for _, o := range m.OnList {
		if cmp.Equal(o.Given, given, protocmp.Transform()) {
			return o.List, o.Error
		}
	}
	return nil, fmt.Errorf("mock not found for %v", given)
}
