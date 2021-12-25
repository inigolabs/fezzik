package github

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik/client"
)

type Client interface {
	// GetInfo from examples/github/operations/operations.graphql:1
	GetInfo(ctx context.Context,
		repoOwner string,
		repoName string,
	) (*GetInfoResponse, error)

	// CreatePullRequest from examples/github/operations/operations.graphql:39
	CreatePullRequest(ctx context.Context,
		input CreatePullRequestInput,
	) (*CreatePullRequestResponse, error)

	// UpdatePullRequest from examples/github/operations/operations.graphql:50
	UpdatePullRequest(ctx context.Context,
		input UpdatePullRequestInput,
	) (*UpdatePullRequestResponse, error)

	// CommentPullRequest from examples/github/operations/operations.graphql:60
	CommentPullRequest(ctx context.Context,
		input AddCommentInput,
	) (*CommentPullRequestResponse, error)

	// MergePullRequest from examples/github/operations/operations.graphql:68
	MergePullRequest(ctx context.Context,
		input MergePullRequestInput,
	) (*MergePullRequestResponse, error)

	// ClosePullRequest from examples/github/operations/operations.graphql:78
	ClosePullRequest(ctx context.Context,
		input ClosePullRequestInput,
	) (*ClosePullRequestResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient),
	}
}

type gqlclient struct {
	gql *client.GQLClient
}
