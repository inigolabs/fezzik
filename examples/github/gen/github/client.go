package github

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik"
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
	gqlclient := fezzik.NewGQLClient(url, httpclient)
	return &client{
		gql: gqlclient,
	}
}

type client struct {
	gql *fezzik.GQLClient
}
