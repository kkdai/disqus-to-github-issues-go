package disqusimportorgo

import "testing"

func TestIssueAppendSort(t *testing.T) {
	i := NewIssue(Article{Title: "test", Link: "https://sss.ccc/path1/"})
	i.AppendComment(Comment{Author: AuthorStruct{Name: "a1"}, CreatedAt: "2020-11-15T04:34:00Z", Message: "msg 1"})
	i.AppendComment(Comment{Author: AuthorStruct{Name: "a2"}, CreatedAt: "2019-09-15T05:34:00Z", Message: "msg 2"})
	i.AppendComment(Comment{Author: AuthorStruct{Name: "a3"}, CreatedAt: "2018-11-08T04:34:00Z", Message: "msg 3"})

	i.SortComments()

	if i.Comments[0].CreatedAt.After(i.Comments[1].CreatedAt) {
		t.Errorf("Sort error on issues %v, %v", i.Comments[0], i.Comments[1])
	}
}
