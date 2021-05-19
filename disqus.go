package disqusimportorgo

import (
	x "encoding/xml"
	"fmt"
)

//Disqus :
type Disqus struct {
	data *DisqusStruct
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

func isCommentBelongArticle(c Comment, a Article) bool {
	return c.Article.ID == a.AttrID
}
