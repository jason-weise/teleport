/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package labels

import (
	"context"
	"testing"
	"time"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"
)

type mockIMDSClient struct {
	tagsDisabled bool
	tags         map[string]string
}

func (m *mockIMDSClient) IsAvailable(ctx context.Context) bool {
	return true
}

func (m *mockIMDSClient) GetType() types.InstanceMetadataType {
	return "mock"
}

func (m *mockIMDSClient) GetTags(ctx context.Context) (map[string]string, error) {
	if m.tagsDisabled {
		return nil, trace.NotFound("Tags not available")
	}
	return m.tags, nil
}

func (m *mockIMDSClient) GetHostname(ctx context.Context) (string, error) {
	value, ok := m.tags[types.CloudHostnameTag]
	if !ok {
		return "", trace.NotFound("Tag TeleportHostname not found")
	}
	return value, nil
}

func TestCloudLabelsSync(t *testing.T) {
	ctx := context.Background()
	tags := map[string]string{"a": "1", "b": "2"}
	expectedTags := map[string]string{"aws/a": "1", "aws/b": "2"}
	imdsClient := &mockIMDSClient{
		tags: tags,
	}
	ec2Labels, err := NewCloudImporter(ctx, &CloudConfig{
		Client:    imdsClient,
		namespace: "aws",
	})
	require.NoError(t, err)
	require.NoError(t, ec2Labels.Sync(ctx))
	require.Equal(t, expectedTags, ec2Labels.Get())
}

func TestCloudLabelsAsync(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	imdsClient := &mockIMDSClient{}
	clock := clockwork.NewFakeClock()
	ec2Labels, err := NewCloudImporter(ctx, &CloudConfig{
		Client:    imdsClient,
		namespace: "aws",
		Clock:     clock,
	})
	require.NoError(t, err)

	compareLabels := func(m map[string]string) func() bool {
		return func() bool {
			labels := ec2Labels.Get()
			if len(labels) != len(m) {
				return false
			}
			for k, v := range labels {
				if m[k] != v {
					return false
				}
			}
			return true
		}
	}

	// Check that initial tags are read.
	initialTags := map[string]string{"a": "1", "b": "2"}
	imdsClient.tags = initialTags
	ec2Labels.Start(ctx)
	require.Eventually(t, compareLabels(map[string]string{"aws/a": "1", "aws/b": "2"}), time.Second, 100*time.Microsecond)

	// Check that tags are updated over time.
	updatedTags := map[string]string{"a": "3", "c": "4"}
	imdsClient.tags = updatedTags
	clock.Advance(labelUpdatePeriod)
	require.Eventually(t, compareLabels(map[string]string{"aws/a": "3", "aws/c": "4"}), time.Second, 100*time.Millisecond)

	// Check that service stops updating when closed.
	cancel()
	imdsClient.tags = map[string]string{"x": "8", "y": "9", "z": "10"}
	clock.Advance(labelUpdatePeriod)
	require.Eventually(t, compareLabels(map[string]string{"aws/a": "3", "aws/c": "4"}), time.Second, 100*time.Millisecond)
}

func TestCloudLabelsValidKey(t *testing.T) {
	ctx := context.Background()
	tags := map[string]string{"good-label": "1", "bad-l@bel": "2"}
	expectedTags := map[string]string{"aws/good-label": "1"}
	imdsClient := &mockIMDSClient{
		tags: tags,
	}
	ec2Labels, err := NewCloudImporter(ctx, &CloudConfig{
		Client:    imdsClient,
		namespace: "aws",
	})
	require.NoError(t, err)
	require.NoError(t, ec2Labels.Sync(ctx))
	require.Equal(t, expectedTags, ec2Labels.Get())
}

func TestCloudLabelsDisabled(t *testing.T) {
	ctx := context.Background()
	imdsClient := &mockIMDSClient{
		tagsDisabled: true,
	}
	ec2Labels, err := NewCloudImporter(ctx, &CloudConfig{
		Client:    imdsClient,
		namespace: "aws",
	})
	require.NoError(t, err)
	require.NoError(t, ec2Labels.Sync(ctx))
	require.Equal(t, map[string]string{}, ec2Labels.Get())
}
