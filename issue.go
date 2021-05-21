package disqusimportorgo

import (
	"sort"
	"time"
)

type Issue struct {
	ArticleTitle string
	ArticleLink  string
	ShortLink    string
	Comments     []IssueComment
}

func NewIssue(article Article) *Issue {
	return &Issue{
		ArticleTitle: article.Title,
		ArticleLink:  article.Link,
		ShortLink:    getShortPath(article.Link),
	}
}

func (issue *Issue) AppendComment(c Comment) {
	t, _ := time.Parse(time.RFC3339, c.CreatedAt)

	issue.Comments = append(issue.Comments, IssueComment{
		Author:    c.GetAuthorName(),
		CreatedAt: t,
		Body:      c.Message,
	})
}

func (issue *Issue) SortComments() {
	if issue != nil && len(issue.Comments) > 2 {
		sort.Sort(ByCreateAt(issue.Comments))
	}
}
