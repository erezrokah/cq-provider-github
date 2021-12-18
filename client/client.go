package client

import (
	"context"
	"os"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v40/github"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

	GithubClient *github.Client
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	ctx := context.Background()
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		token = providerConfig.GitHubToken
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)

	_ = providerConfig
	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	client := Client{
		logger: logger,
		// CHANGEME: pass the initialized third pard client
		GithubClient: githubClient,
	}

	// Return the initialized client and it will be passed to your resources
	return &client, nil
}
