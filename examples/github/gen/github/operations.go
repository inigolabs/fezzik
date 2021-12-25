package github

import (
	"context"

	"github.com/inigolabs/fezzik/client"
)

var GetInfoOperation string = `
query GetInfo($repo_owner : String!, $repo_name : String!) {
   viewer {
      login
      pullRequests(first: 100, states: [OPEN]) {
         nodes {
            id
            number
            title
            baseRefName
            headRefName
            mergeable
            reviewDecision
            repository {
               id
            }
            commits(first: 100) {
               nodes {
                  commit {
                     oid
                     messageHeadline
                     messageBody
                     statusCheckRollup {
                        state
                     }
                  }
               }
            }
         }
      }
   }
   repository(owner: $repo_owner, name: $repo_name) {
      id
   }
}
`

type GetInfoResponse struct {
	Viewer struct {
		Login        string
		PullRequests struct {
			Nodes *[]*struct {
				Id             string
				Number         int
				Title          string
				BaseRefName    string
				HeadRefName    string
				Mergeable      MergeableState
				ReviewDecision *PullRequestReviewDecision
				Repository     struct {
					Id string
				}
				Commits struct {
					Nodes *[]*struct {
						Commit struct {
							Oid               string
							MessageHeadline   string
							MessageBody       string
							StatusCheckRollup *struct {
								State StatusState
							}
						}
					}
				}
			}
		}
	}
	Repository *struct {
		Id string
	}
}

// GetInfo from examples/github/operations/operations.graphql:1
func (c *gqlclient) GetInfo(ctx context.Context,
	repoOwner string,
	repoName string,
) (*GetInfoResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "GetInfo",
		Query:         GetInfoOperation,
		Variables: map[string]interface{}{
			"repo_owner": repoOwner,
			"repo_name":  repoName,
		},
	}

	var gqldata GetInfoResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var CreatePullRequestOperation string = `
mutation CreatePullRequest($input : CreatePullRequestInput!) {
   createPullRequest(input: $input) {
      pullRequest {
         id
         number
      }
   }
}
`

type CreatePullRequestResponse struct {
	CreatePullRequest *struct {
		PullRequest *struct {
			Id     string
			Number int
		}
	}
}

// CreatePullRequest from examples/github/operations/operations.graphql:39
func (c *gqlclient) CreatePullRequest(ctx context.Context,
	input CreatePullRequestInput,
) (*CreatePullRequestResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "CreatePullRequest",
		Query:         CreatePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	var gqldata CreatePullRequestResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var UpdatePullRequestOperation string = `
mutation UpdatePullRequest($input : UpdatePullRequestInput!) {
   updatePullRequest(input: $input) {
      pullRequest {
         number
      }
   }
}
`

type UpdatePullRequestResponse struct {
	UpdatePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

// UpdatePullRequest from examples/github/operations/operations.graphql:50
func (c *gqlclient) UpdatePullRequest(ctx context.Context,
	input UpdatePullRequestInput,
) (*UpdatePullRequestResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "UpdatePullRequest",
		Query:         UpdatePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	var gqldata UpdatePullRequestResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var CommentPullRequestOperation string = `
mutation CommentPullRequest($input : AddCommentInput!) {
   addComment(input: $input) {
      clientMutationId
   }
}
`

type CommentPullRequestResponse struct {
	AddComment *struct {
		ClientMutationId *string
	}
}

// CommentPullRequest from examples/github/operations/operations.graphql:60
func (c *gqlclient) CommentPullRequest(ctx context.Context,
	input AddCommentInput,
) (*CommentPullRequestResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "CommentPullRequest",
		Query:         CommentPullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	var gqldata CommentPullRequestResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var MergePullRequestOperation string = `
mutation MergePullRequest($input : MergePullRequestInput!) {
   mergePullRequest(input: $input) {
      pullRequest {
         number
      }
   }
}
`

type MergePullRequestResponse struct {
	MergePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

// MergePullRequest from examples/github/operations/operations.graphql:68
func (c *gqlclient) MergePullRequest(ctx context.Context,
	input MergePullRequestInput,
) (*MergePullRequestResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "MergePullRequest",
		Query:         MergePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	var gqldata MergePullRequestResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

var ClosePullRequestOperation string = `
mutation ClosePullRequest($input : ClosePullRequestInput!) {
   closePullRequest(input: $input) {
      pullRequest {
         number
      }
   }
}
`

type ClosePullRequestResponse struct {
	ClosePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

// ClosePullRequest from examples/github/operations/operations.graphql:78
func (c *gqlclient) ClosePullRequest(ctx context.Context,
	input ClosePullRequestInput,
) (*ClosePullRequestResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "ClosePullRequest",
		Query:         ClosePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	var gqldata ClosePullRequestResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}
