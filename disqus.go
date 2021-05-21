package disqusimportorgo

import (
	x "encoding/xml"
	"fmt"
	"strings"
)

//Disqus :
type Disqus struct {
	data    *DisqusStruct
	impData map[string]Issue
}

// NewDisqus: create disqus object.
func NewDisqus(xml []byte) *Disqus {
	var comments DisqusStruct
	if err := x.Unmarshal(xml, &comments); err != nil {
		fmt.Printf("err: %s \n", err)
		return nil
	}
	return &Disqus{data: &comments}
}

// GetAllComments: Get all comments.
func (d *Disqus) GetAllComments() []Comment {
	if d.data == nil {
		return nil
	}

	return d.data.Commments
}

func (d *Disqus) GetAllArticles() []Article {
	if d.data == nil {
		return nil
	}

	return d.data.Articles
}

// GetArticleByComment: Get original article by comment.
func (d *Disqus) GetArticleByComment(c Comment) *Article {
	articles := d.GetAllArticles()
	if len(articles) == 0 || c.CreatedAt == "" {
		return nil
	}

	for _, a := range articles {
		if isCommentBelongArticle(c, a) {
			return &a
		}
	}

	return nil
}

// GetAllCommentsByArticle: Get all comments by specific article
func (d *Disqus) GetAllCommentsByArticle(a Article) []Comment {
	allComments := d.GetAllComments()
	if len(allComments) == 0 {
		return nil
	}

	var retComments []Comment
	for _, c := range allComments {
		if isCommentBelongArticle(c, a) {
			retComments = append(retComments, c)
		}
	}

	return retComments
}

// PrepareImportData:
func (d *Disqus) PrepareImportData() error {
	if d.data == nil {
		return fmt.Errorf("No source data")
	}

	if d.impData != nil {
		//Data exist.
		return nil
	}

	// Make a article map for quick search cache.
	articleMap := make(map[string]Article)
	for _, a := range d.GetAllArticles() {
		articleMap[a.AttrID] = a
	}

	d.impData = make(map[string]Issue)
	for _, c := range d.GetAllComments() {
		a := articleMap[c.Article.ID]
		shortLink := getShortPath(a.GetArticleLink())
		if issue, exist := d.impData[shortLink]; !exist {
			//not exist, insert new issue.
			ii := Issue{
				ArticleTitle: a.Title,
				ArticleLink:  a.Link,
				ShortLink:    shortLink,
			}
			ii.AppendComment(c)
			d.impData[shortLink] = ii
		} else {
			//Exist, append new comment and update issue.
			issue.AppendComment(c)
			d.impData[shortLink] = issue
		}
	}
	return nil
}

func isCommentBelongArticle(c Comment, a Article) bool {
	return c.Article.ID == a.AttrID
}

//getShortPath: To extra path from disqus to github issue.
// https://www.evanlin.com/reading-twitter/ 		--> reading-twitter/
// https://www.evanlin.com/reading-twitter/sss?111 	--> reading-twitter/
func getShortPath(s1 string) string {
	end := strings.LastIndex(s1, "/")
	start := strings.LastIndex(s1[:end], "/")

	return s1[start+1 : end+1]
}
