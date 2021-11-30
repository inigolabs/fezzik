package github

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik/client"
)

type Client interface {
	GetInfo(ctx context.Context, input *GetInfoInputArgs) (*GetInfoResponse, error)
	CreatePullRequest(ctx context.Context, input *CreatePullRequestInputArgs) (*CreatePullRequestResponse, error)
	UpdatePullRequest(ctx context.Context, input *UpdatePullRequestInputArgs) (*UpdatePullRequestResponse, error)
	CommentPullRequest(ctx context.Context, input *CommentPullRequestInputArgs) (*CommentPullRequestResponse, error)
	MergePullRequest(ctx context.Context, input *MergePullRequestInputArgs) (*MergePullRequestResponse, error)
	ClosePullRequest(ctx context.Context, input *ClosePullRequestInputArgs) (*ClosePullRequestResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient),
	}
}

type gqlclient struct {
	gql *client.GQLClient
}
