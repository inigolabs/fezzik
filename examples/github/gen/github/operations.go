// Code generated by github.com/inigolabs/fezzik, DO NOT EDIT.

package github

import (
	"context"

	"github.com/inigolabs/fezzik/client"
)

type GetInfoViewer struct {
	Login        string
	PullRequests GetInfoViewerPullRequests
}

type GetInfoViewerPullRequests struct {
	Nodes []*GetInfoViewerPullRequestsNodes
}

type GetInfoViewerPullRequestsNodes struct {
	Id             string
	Number         int
	Title          string
	BaseRefName    string
	HeadRefName    string
	Mergeable      MergeableState
	ReviewDecision *PullRequestReviewDecision
	Repository     GetInfoViewerPullRequestsNodesRepository
	Commits        GetInfoViewerPullRequestsNodesCommits
}

type GetInfoViewerPullRequestsNodesRepository struct {
	Id string
}

type GetInfoViewerPullRequestsNodesCommits struct {
	Nodes []*GetInfoViewerPullRequestsNodesCommitsNodes
}

type GetInfoViewerPullRequestsNodesCommitsNodes struct {
	Commit GetInfoViewerPullRequestsNodesCommitsNodesCommit
}

type GetInfoViewerPullRequestsNodesCommitsNodesCommit struct {
	Oid               string
	MessageHeadline   string
	MessageBody       string
	StatusCheckRollup *GetInfoViewerPullRequestsNodesCommitsNodesCommitStatusCheckRollup
}

type GetInfoViewerPullRequestsNodesCommitsNodesCommitStatusCheckRollup struct {
	State StatusState
}

type GetInfoRepository struct {
	Id string
}

// GetInfoResponse response type for GetInfo
type GetInfoResponse struct {
	Viewer     GetInfoViewer
	Repository *GetInfoRepository
}

// GetInfo from examples/github/operations/operations.graphql:1
func (c *gqlclient) GetInfo(ctx context.Context,
	repoOwner string,
	repoName string,
) (*GetInfoResponse, error) {

	var getInfoOperation string = `
	query GetInfo ($repo_owner: String!, $repo_name: String!) {
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

	gqlreq := &client.GQLRequest{
		OperationName: "GetInfo",
		Query:         getInfoOperation,
		Variables: map[string]interface{}{
			"repo_owner": repoOwner,
			"repo_name":  repoName,
		},
	}

	resp := &client.GQLResponse{
		Data: &GetInfoResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *GetInfoResponse
	if resp.Data != nil {
		data = resp.Data.(*GetInfoResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type CreatePullRequestCreatePullRequest struct {
	PullRequest *CreatePullRequestCreatePullRequestPullRequest
}

type CreatePullRequestCreatePullRequestPullRequest struct {
	Id     string
	Number int
}

// CreatePullRequestResponse response type for CreatePullRequest
type CreatePullRequestResponse struct {
	CreatePullRequest *CreatePullRequestCreatePullRequest
}

// CreatePullRequest from examples/github/operations/operations.graphql:39
func (c *gqlclient) CreatePullRequest(ctx context.Context,
	input CreatePullRequestInput,
) (*CreatePullRequestResponse, error) {

	var createPullRequestOperation string = `
	mutation CreatePullRequest ($input: CreatePullRequestInput!) {
	createPullRequest(input: $input) {
		pullRequest {
			id
			number
		}
	}
}
`

	gqlreq := &client.GQLRequest{
		OperationName: "CreatePullRequest",
		Query:         createPullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	resp := &client.GQLResponse{
		Data: &CreatePullRequestResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *CreatePullRequestResponse
	if resp.Data != nil {
		data = resp.Data.(*CreatePullRequestResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type UpdatePullRequestUpdatePullRequest struct {
	PullRequest *UpdatePullRequestUpdatePullRequestPullRequest
}

type UpdatePullRequestUpdatePullRequestPullRequest struct {
	Number int
}

// UpdatePullRequestResponse response type for UpdatePullRequest
type UpdatePullRequestResponse struct {
	UpdatePullRequest *UpdatePullRequestUpdatePullRequest
}

// UpdatePullRequest from examples/github/operations/operations.graphql:50
func (c *gqlclient) UpdatePullRequest(ctx context.Context,
	input UpdatePullRequestInput,
) (*UpdatePullRequestResponse, error) {

	var updatePullRequestOperation string = `
	mutation UpdatePullRequest ($input: UpdatePullRequestInput!) {
	updatePullRequest(input: $input) {
		pullRequest {
			number
		}
	}
}
`

	gqlreq := &client.GQLRequest{
		OperationName: "UpdatePullRequest",
		Query:         updatePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	resp := &client.GQLResponse{
		Data: &UpdatePullRequestResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *UpdatePullRequestResponse
	if resp.Data != nil {
		data = resp.Data.(*UpdatePullRequestResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type CommentPullRequestAddComment struct {
	ClientMutationId *string
}

// CommentPullRequestResponse response type for CommentPullRequest
type CommentPullRequestResponse struct {
	AddComment *CommentPullRequestAddComment
}

// CommentPullRequest from examples/github/operations/operations.graphql:60
func (c *gqlclient) CommentPullRequest(ctx context.Context,
	input AddCommentInput,
) (*CommentPullRequestResponse, error) {

	var commentPullRequestOperation string = `
	mutation CommentPullRequest ($input: AddCommentInput!) {
	addComment(input: $input) {
		clientMutationId
	}
}
`

	gqlreq := &client.GQLRequest{
		OperationName: "CommentPullRequest",
		Query:         commentPullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	resp := &client.GQLResponse{
		Data: &CommentPullRequestResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *CommentPullRequestResponse
	if resp.Data != nil {
		data = resp.Data.(*CommentPullRequestResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type MergePullRequestMergePullRequest struct {
	PullRequest *MergePullRequestMergePullRequestPullRequest
}

type MergePullRequestMergePullRequestPullRequest struct {
	Number int
}

// MergePullRequestResponse response type for MergePullRequest
type MergePullRequestResponse struct {
	MergePullRequest *MergePullRequestMergePullRequest
}

// MergePullRequest from examples/github/operations/operations.graphql:68
func (c *gqlclient) MergePullRequest(ctx context.Context,
	input MergePullRequestInput,
) (*MergePullRequestResponse, error) {

	var mergePullRequestOperation string = `
	mutation MergePullRequest ($input: MergePullRequestInput!) {
	mergePullRequest(input: $input) {
		pullRequest {
			number
		}
	}
}
`

	gqlreq := &client.GQLRequest{
		OperationName: "MergePullRequest",
		Query:         mergePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	resp := &client.GQLResponse{
		Data: &MergePullRequestResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *MergePullRequestResponse
	if resp.Data != nil {
		data = resp.Data.(*MergePullRequestResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type ClosePullRequestClosePullRequest struct {
	PullRequest *ClosePullRequestClosePullRequestPullRequest
}

type ClosePullRequestClosePullRequestPullRequest struct {
	Number int
}

// ClosePullRequestResponse response type for ClosePullRequest
type ClosePullRequestResponse struct {
	ClosePullRequest *ClosePullRequestClosePullRequest
}

// ClosePullRequest from examples/github/operations/operations.graphql:78
func (c *gqlclient) ClosePullRequest(ctx context.Context,
	input ClosePullRequestInput,
) (*ClosePullRequestResponse, error) {

	var closePullRequestOperation string = `
	mutation ClosePullRequest ($input: ClosePullRequestInput!) {
	closePullRequest(input: $input) {
		pullRequest {
			number
		}
	}
}
`

	gqlreq := &client.GQLRequest{
		OperationName: "ClosePullRequest",
		Query:         closePullRequestOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	resp := &client.GQLResponse{
		Data: &ClosePullRequestResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *ClosePullRequestResponse
	if resp.Data != nil {
		data = resp.Data.(*ClosePullRequestResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}
