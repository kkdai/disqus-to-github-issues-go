package disqusimportorgo

import (
	"context"
	"fmt"
	"log"
	"time"

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

//CheckIfExist : Implement later.
func (b *CommentClient) CheckIfExist() bool {
	return false
}

//CreateIssue :
func (b *CommentClient) CreateIssue(i *Issue) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: b.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	commentBody := fmt.Sprintf("# %s \n \n \n [%s](%s)", i.ArticleTitle, i.ArticleLink, i.ArticleLink)

	input := &github.IssueRequest{
		Title:    String(i.ShortLink),
		Body:     String(commentBody),
		Assignee: String(""),
		Labels:   &[]string{}, //&tags,
	}

	var gIssue *github.Issue
	var err error
	gIssue, _, err = client.Issues.Create(ctx, b.User, b.Repo, input)
	if err != nil {
		fmt.Println("Issues.Create returned error: ", err, " retry after 2 seconds.")
		time.Sleep(2 * time.Second)

		//retry once
		gIssue, _, err = client.Issues.Create(ctx, b.User, b.Repo, input)
		if err != nil {
			fmt.Println("Issues.Create returned error: ", err, " retry after 2 seconds.")
			return err
		}
	}

	///Sort it before use it.
	i.SortComments()

	var id int64
	for _, c := range i.Comments {
		id = id + 1
		body := fmt.Sprintf("comment written by %s, created at %s, \n\n %s", c.Author, c.CreatedAt.Format(time.RFC822), c.Body)
		cm := &github.IssueComment{
			ID:   &(id),
			Body: &body}

		if _, res, err := client.Issues.CreateComment(ctx, b.User, b.Repo, *gIssue.Number, cm); err != nil {
			log.Println("Create comment res", res, " error code:", err)
			return err
		}

		//sleep 500 millisecond to avoid github limited.
		time.Sleep(300 * time.Millisecond)
	}
	return nil
}
