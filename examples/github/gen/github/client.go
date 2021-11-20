package github

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik"
)

type Client interface {
	GetInfo(ctx context.Context, input *GetInfoInputArgs) (*GetInfoResponse, fezzik.GQLErrors, error)
	CreatePullRequest(ctx context.Context, input *CreatePullRequestInputArgs) (*CreatePullRequestResponse, fezzik.GQLErrors, error)
	UpdatePullRequest(ctx context.Context, input *UpdatePullRequestInputArgs) (*UpdatePullRequestResponse, fezzik.GQLErrors, error)
	CommentPullRequest(ctx context.Context, input *CommentPullRequestInputArgs) (*CommentPullRequestResponse, fezzik.GQLErrors, error)
	MergePullRequest(ctx context.Context, input *MergePullRequestInputArgs) (*MergePullRequestResponse, fezzik.GQLErrors, error)
	ClosePullRequest(ctx context.Context, input *ClosePullRequestInputArgs) (*ClosePullRequestResponse, fezzik.GQLErrors, error)
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
