package disqusimportorgo

import (
	"os"
	"testing"
)

func TestPostGithubIssue(t *testing.T) {
	token := os.Getenv("Token")
	user := os.Getenv("User")
	repo := os.Getenv("Repo")

	if len(token) == 0 || len(user) == 0 || len(user) == 0 {
		t.Skip("Please input github env")
		return
	}

	i := NewIssue(Article{Title: "test", Link: "https://sss.ccc/path1/"})
	i.AppendComment(Comment{Author: AuthorStruct{Name: "a1"}, CreatedAt: "2020-11-15T04:34:00Z", Message: "msg 1"})
	i.AppendComment(Comment{Author: AuthorStruct{Name: "a2"}, CreatedAt: "2019-09-15T05:34:00Z", Message: "msg 2"})
	i.AppendComment(Comment{Author: AuthorStruct{Name: "a3"}, CreatedAt: "2018-11-08T04:34:00Z", Message: "msg 3"})

	i.SortComments()

	c := NewCommentClient(user, repo, token)
	if err := c.CreateIssue(i); err != nil {
		t.Error("Err on commment client: ", err)
	}
}
