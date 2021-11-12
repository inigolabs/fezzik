package github

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/mitchellh/mapstructure"
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

type GetInfoInputArgs struct {
	RepoOwner string `json:"repo_owner"`
	RepoName  string `json:"repo_name"`
}

type GetInfoResponse struct {
	Viewer struct {
		Login        string
		PullRequests struct {
			Nodes []*struct {
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
					Nodes []*struct {
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

func (c *client) GetInfo(ctx context.Context, input *GetInfoInputArgs) (
	*GetInfoResponse, error) {

	q := graphql.NewRequest(GetInfoOperation)
	q.Var("repo_owner", input.RepoOwner)
	q.Var("repo_name", input.RepoName)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)

	if err != nil {
		return nil, err
	}

	output := GetInfoResponse{}
	err = mapstructure.Decode(resp, &output)

	if err != nil {
		return nil, err
	}
	return &output, err
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

type CreatePullRequestInputArgs struct {
	Input CreatePullRequestInput `json:"input"`
}

type CreatePullRequestResponse struct {
	CreatePullRequest *struct {
		PullRequest *struct {
			Id     string
			Number int
		}
	}
}

func (c *client) CreatePullRequest(ctx context.Context, input *CreatePullRequestInputArgs) (
	*CreatePullRequestResponse, error) {

	q := graphql.NewRequest(CreatePullRequestOperation)
	q.Var("input", input.Input)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)

	if err != nil {
		return nil, err
	}

	output := CreatePullRequestResponse{}
	err = mapstructure.Decode(resp, &output)

	if err != nil {
		return nil, err
	}
	return &output, err
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

type UpdatePullRequestInputArgs struct {
	Input UpdatePullRequestInput `json:"input"`
}

type UpdatePullRequestResponse struct {
	UpdatePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

func (c *client) UpdatePullRequest(ctx context.Context, input *UpdatePullRequestInputArgs) (
	*UpdatePullRequestResponse, error) {

	q := graphql.NewRequest(UpdatePullRequestOperation)
	q.Var("input", input.Input)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)

	if err != nil {
		return nil, err
	}

	output := UpdatePullRequestResponse{}
	err = mapstructure.Decode(resp, &output)

	if err != nil {
		return nil, err
	}
	return &output, err
}

var CommentPullRequestOperation string = `
mutation CommentPullRequest($input : AddCommentInput!) {
   addComment(input: $input) {
      clientMutationId
   }
}
`

type CommentPullRequestInputArgs struct {
	Input AddCommentInput `json:"input"`
}

type CommentPullRequestResponse struct {
	AddComment *struct {
		ClientMutationId *string
	}
}

func (c *client) CommentPullRequest(ctx context.Context, input *CommentPullRequestInputArgs) (
	*CommentPullRequestResponse, error) {

	q := graphql.NewRequest(CommentPullRequestOperation)
	q.Var("input", input.Input)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)

	if err != nil {
		return nil, err
	}

	output := CommentPullRequestResponse{}
	err = mapstructure.Decode(resp, &output)

	if err != nil {
		return nil, err
	}
	return &output, err
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

type MergePullRequestInputArgs struct {
	Input MergePullRequestInput `json:"input"`
}

type MergePullRequestResponse struct {
	MergePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

func (c *client) MergePullRequest(ctx context.Context, input *MergePullRequestInputArgs) (
	*MergePullRequestResponse, error) {

	q := graphql.NewRequest(MergePullRequestOperation)
	q.Var("input", input.Input)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)

	if err != nil {
		return nil, err
	}

	output := MergePullRequestResponse{}
	err = mapstructure.Decode(resp, &output)

	if err != nil {
		return nil, err
	}
	return &output, err
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

type ClosePullRequestInputArgs struct {
	Input ClosePullRequestInput `json:"input"`
}

type ClosePullRequestResponse struct {
	ClosePullRequest *struct {
		PullRequest *struct {
			Number int
		}
	}
}

func (c *client) ClosePullRequest(ctx context.Context, input *ClosePullRequestInputArgs) (
	*ClosePullRequestResponse, error) {

	q := graphql.NewRequest(ClosePullRequestOperation)
	q.Var("input", input.Input)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)

	if err != nil {
		return nil, err
	}

	output := ClosePullRequestResponse{}
	err = mapstructure.Decode(resp, &output)

	if err != nil {
		return nil, err
	}
	return &output, err
}
