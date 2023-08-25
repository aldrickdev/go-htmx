// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// GetIssuesRepository includes the requested fields of the GraphQL type Repository.
// The GraphQL type's documentation follows.
//
// A repository contains the content for a project.
type GetIssuesRepository struct {
	// A list of issues that have been opened in the repository.
	Issues GetIssuesRepositoryIssuesIssueConnection `json:"issues"`
}

// GetIssues returns GetIssuesRepository.Issues, and is useful for accessing the field via an interface.
func (v *GetIssuesRepository) GetIssues() GetIssuesRepositoryIssuesIssueConnection { return v.Issues }

// GetIssuesRepositoryIssuesIssueConnection includes the requested fields of the GraphQL type IssueConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Issue.
type GetIssuesRepositoryIssuesIssueConnection struct {
	// A list of edges.
	Edges []GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdge `json:"edges"`
}

// GetEdges returns GetIssuesRepositoryIssuesIssueConnection.Edges, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnection) GetEdges() []GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdge {
	return v.Edges
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdge includes the requested fields of the GraphQL type IssueEdge.
// The GraphQL type's documentation follows.
//
// An edge in a connection.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdge struct {
	// The item at the end of the edge.
	Node GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue `json:"node"`
}

// GetNode returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdge.Node, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdge) GetNode() GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue {
	return v.Node
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue includes the requested fields of the GraphQL type Issue.
// The GraphQL type's documentation follows.
//
// An Issue is a place to discuss ideas, enhancements, tasks, and bugs for a project.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue struct {
	// Identifies the issue title.
	Title string `json:"title"`
	// Identifies the issue number.
	Number int `json:"number"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
	// The actor who authored the comment.
	Author GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor `json:"-"`
	// A list of labels associated with the object.
	Labels GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnection `json:"labels"`
}

// GetTitle returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue.Title, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) GetTitle() string {
	return v.Title
}

// GetNumber returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue.Number, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) GetNumber() int {
	return v.Number
}

// GetCreatedAt returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue.CreatedAt, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) GetCreatedAt() time.Time {
	return v.CreatedAt
}

// GetAuthor returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue.Author, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) GetAuthor() GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor {
	return v.Author
}

// GetLabels returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue.Labels, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) GetLabels() GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnection {
	return v.Labels
}

func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue
		Author json.RawMessage `json:"author"`
		graphql.NoUnmarshalJSON
	}
	firstPass.GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Author
		src := firstPass.Author
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue.Author: %w", err)
			}
		}
	}
	return nil
}

type __premarshalGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue struct {
	Title string `json:"title"`

	Number int `json:"number"`

	CreatedAt time.Time `json:"createdAt"`

	Author json.RawMessage `json:"author"`

	Labels GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnection `json:"labels"`
}

func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue) __premarshalJSON() (*__premarshalGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue, error) {
	var retval __premarshalGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue

	retval.Title = v.Title
	retval.Number = v.Number
	retval.CreatedAt = v.CreatedAt
	{

		dst := &retval.Author
		src := v.Author
		var err error
		*dst, err = __marshalGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssue.Author: %w", err)
		}
	}
	retval.Labels = v.Labels
	return &retval, nil
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor includes the requested fields of the GraphQL interface Actor.
//
// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor is implemented by the following types:
// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot
// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount
// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin
// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization
// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser
// The GraphQL type's documentation follows.
//
// Represents an object which can take actions on GitHub. Typically a User or Bot.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor interface {
	implementsGraphQLInterfaceGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetLogin returns the interface-field "login" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// The username of the actor.
	GetLogin() string
}

func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot) implementsGraphQLInterfaceGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor() {
}
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount) implementsGraphQLInterfaceGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor() {
}
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin) implementsGraphQLInterfaceGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor() {
}
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization) implementsGraphQLInterfaceGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor() {
}
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser) implementsGraphQLInterfaceGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor() {
}

func __unmarshalGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor(b []byte, v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Bot":
		*v = new(GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot)
		return json.Unmarshal(b, *v)
	case "EnterpriseUserAccount":
		*v = new(GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount)
		return json.Unmarshal(b, *v)
	case "Mannequin":
		*v = new(GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin)
		return json.Unmarshal(b, *v)
	case "Organization":
		*v = new(GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization)
		return json.Unmarshal(b, *v)
	case "User":
		*v = new(GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing Actor.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor: "%v"`, tn.TypeName)
	}
}

func __marshalGetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor(v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot:
		typename = "Bot"

		result := struct {
			TypeName string `json:"__typename"`
			*GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot
		}{typename, v}
		return json.Marshal(result)
	case *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount:
		typename = "EnterpriseUserAccount"

		result := struct {
			TypeName string `json:"__typename"`
			*GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount
		}{typename, v}
		return json.Marshal(result)
	case *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin:
		typename = "Mannequin"

		result := struct {
			TypeName string `json:"__typename"`
			*GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin
		}{typename, v}
		return json.Marshal(result)
	case *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization:
		typename = "Organization"

		result := struct {
			TypeName string `json:"__typename"`
			*GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization
		}{typename, v}
		return json.Marshal(result)
	case *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser:
		typename = "User"

		result := struct {
			TypeName string `json:"__typename"`
			*GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorActor: "%T"`, v)
	}
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot includes the requested fields of the GraphQL type Bot.
// The GraphQL type's documentation follows.
//
// A special type of user which takes actions on behalf of GitHub Apps.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot struct {
	Typename string `json:"__typename"`
	// The username of the actor.
	Login string `json:"login"`
}

// GetTypename returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot.Typename, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot) GetTypename() string {
	return v.Typename
}

// GetLogin returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot.Login, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorBot) GetLogin() string {
	return v.Login
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount includes the requested fields of the GraphQL type EnterpriseUserAccount.
// The GraphQL type's documentation follows.
//
// An account for a user who is an admin of an enterprise or a member of an enterprise through one or more organizations.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount struct {
	Typename string `json:"__typename"`
	// The username of the actor.
	Login string `json:"login"`
}

// GetTypename returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount.Typename, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount) GetTypename() string {
	return v.Typename
}

// GetLogin returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount.Login, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorEnterpriseUserAccount) GetLogin() string {
	return v.Login
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin includes the requested fields of the GraphQL type Mannequin.
// The GraphQL type's documentation follows.
//
// A placeholder user for attribution of imported data on GitHub.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin struct {
	Typename string `json:"__typename"`
	// The username of the actor.
	Login string `json:"login"`
}

// GetTypename returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin.Typename, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin) GetTypename() string {
	return v.Typename
}

// GetLogin returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin.Login, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorMannequin) GetLogin() string {
	return v.Login
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An account on GitHub, with one or more owners, that has repositories, members and teams.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization struct {
	Typename string `json:"__typename"`
	// The username of the actor.
	Login string `json:"login"`
}

// GetTypename returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization.Typename, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization) GetTypename() string {
	return v.Typename
}

// GetLogin returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization.Login, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorOrganization) GetLogin() string {
	return v.Login
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser struct {
	Typename string `json:"__typename"`
	// The username of the actor.
	Login string `json:"login"`
}

// GetTypename returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser.Typename, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser) GetTypename() string {
	return v.Typename
}

// GetLogin returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser.Login, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueAuthorUser) GetLogin() string {
	return v.Login
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnection includes the requested fields of the GraphQL type LabelConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Label.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnection struct {
	// A list of nodes.
	Nodes []GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnectionNodesLabel `json:"nodes"`
}

// GetNodes returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnection) GetNodes() []GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnectionNodesLabel {
	return v.Nodes
}

// GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnectionNodesLabel includes the requested fields of the GraphQL type Label.
// The GraphQL type's documentation follows.
//
// A label for categorizing Issues, Pull Requests, Milestones, or Discussions with a given Repository.
type GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnectionNodesLabel struct {
	// Identifies the label name.
	Name string `json:"name"`
}

// GetName returns GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnectionNodesLabel.Name, and is useful for accessing the field via an interface.
func (v *GetIssuesRepositoryIssuesIssueConnectionEdgesIssueEdgeNodeIssueLabelsLabelConnectionNodesLabel) GetName() string {
	return v.Name
}

// GetIssuesResponse is returned by GetIssues on success.
type GetIssuesResponse struct {
	// Lookup a given repository by the owner and repository name.
	Repository GetIssuesRepository `json:"repository"`
}

// GetRepository returns GetIssuesResponse.Repository, and is useful for accessing the field via an interface.
func (v *GetIssuesResponse) GetRepository() GetIssuesRepository { return v.Repository }

// __GetIssuesInput is used internally by genqlient
type __GetIssuesInput struct {
	Owner    string `json:"owner"`
	RepoName string `json:"repoName"`
}

// GetOwner returns __GetIssuesInput.Owner, and is useful for accessing the field via an interface.
func (v *__GetIssuesInput) GetOwner() string { return v.Owner }

// GetRepoName returns __GetIssuesInput.RepoName, and is useful for accessing the field via an interface.
func (v *__GetIssuesInput) GetRepoName() string { return v.RepoName }

// The query or mutation executed by GetIssues.
const GetIssues_Operation = `
query GetIssues ($owner: String!, $repoName: String!) {
	repository(owner: $owner, name: $repoName) {
		issues(first: 30, states: [OPEN], orderBy: {field:CREATED_AT,direction:DESC}) {
			edges {
				node {
					title
					number
					createdAt
					author {
						__typename
						login
					}
					labels(first: 5) {
						nodes {
							name
						}
					}
				}
			}
		}
	}
}
`

func GetIssues(
	ctx context.Context,
	client graphql.Client,
	owner string,
	repoName string,
) (*GetIssuesResponse, error) {
	req := &graphql.Request{
		OpName: "GetIssues",
		Query:  GetIssues_Operation,
		Variables: &__GetIssuesInput{
			Owner:    owner,
			RepoName: repoName,
		},
	}
	var err error

	var data GetIssuesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
