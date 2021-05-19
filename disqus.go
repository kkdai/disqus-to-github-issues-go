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
func (d *Disqus) GetAllComments() []Commment {
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
func (d *Disqus) GetArticleByComment(c Commment) *Article {
	articles := d.GetAllArticles()
	if len(articles) == 0 || c.CreatedAt == "" {
		return nil
	}

	for _, a := range articles {
		if a.AttrID == c.Article.ID {
			return &a
		}
	}

	return nil
}
