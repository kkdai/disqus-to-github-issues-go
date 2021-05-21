package disqusimportorgo

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func String(v string) *string { return &v }

//CommentClient :
type CommentClient struct {
	Token string
	User  string
	Repo  string
}

//NewCommentClient :
func NewCommentClient(user, repo, token string) *CommentClient {
	new := new(CommentClient)
	new.User = user
	new.Repo = repo
	new.Token = token
	return new
}

//CheckIfExist :
func (b *CommentClient) CheckIfExist() bool {
	return false
}

//CreateIssue :
func (b *CommentClient) CreateIssue(shortLink, fullTitle, link string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: b.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	commentBody := fmt.Sprintf("# %s \n \n \n [%s](%s)", fullTitle, link, link)

	input := &github.IssueRequest{
		Title:    String(shortLink),
		Body:     String(commentBody),
		Assignee: String(""),
		Labels:   &[]string{}, //&tags,
	}

	_, _, err := client.Issues.Create(ctx, b.User, b.Repo, input)
	if err != nil {
		fmt.Printf("Issues.Create returned error: %v", err)
		return err
	}

	return nil
}
