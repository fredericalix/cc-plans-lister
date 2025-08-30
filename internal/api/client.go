package api

import (
	"context"
	"os"
	"sort"

	"go.clever-cloud.dev/client"

	"cc-plans-lister/pkg/clevercloud"
)

// Client wraps the Clever Cloud API client
type Client struct {
	cc *client.Client
}

// NewClient creates a new API client with the provided token
func NewClient(token string) *Client {
	// Set up environment variables for OAuth authentication
	os.Setenv("CLEVER_TOKEN", token)

	// Create client with auto OAuth configuration
	cc := client.New(client.WithAutoOauthConfig())

	return &Client{cc: cc}
}

// GetAddonProviders fetches all addon providers from the Clever Cloud API
func (c *Client) GetAddonProviders(ctx context.Context) ([]clevercloud.AddonProvider, error) {
	addonRes := client.Get[[]clevercloud.AddonProvider](ctx, c.cc, "/v2/products/addonproviders")

	if addonRes.HasError() {
		return nil, addonRes.Error()
	}

	providers := *addonRes.Payload()

	// Sort providers by ID for consistent output
	sort.Slice(providers, func(i, j int) bool {
		return providers[i].ID < providers[j].ID
	})

	return providers, nil
}

// GetProductInstances fetches all application instances from the Clever Cloud API
func (c *Client) GetProductInstances(ctx context.Context) ([]clevercloud.ProductInstance, error) {
	instanceRes := client.Get[[]clevercloud.ProductInstance](ctx, c.cc, "/v2/products/instances")

	if instanceRes.HasError() {
		return nil, instanceRes.Error()
	}

	instances := *instanceRes.Payload()

	// Sort instances by type for consistent output
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].Type < instances[j].Type
	})

	return instances, nil
}
