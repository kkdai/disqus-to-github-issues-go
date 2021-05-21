package disqusimportorgo

import "time"

type ByCreateAt []IssueComment

func (a ByCreateAt) Len() int           { return len(a) }
func (a ByCreateAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreateAt) Less(i, j int) bool { return a[i].CreatedAt.Before(a[j].CreatedAt) }

type IssueComment struct {
	Author    string
	CreatedAt time.Time
	Body      string
}
