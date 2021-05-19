package disqusimportorgo

import (
	x "encoding/xml"
	"fmt"
)

type Disqus struct {
	data *DisqusStruct
}

func NewDisqus(xml []byte) *Disqus {
	var comments DisqusStruct
	if err := x.Unmarshal(xml, &comments); err != nil {
		fmt.Printf("err: %s \n", err)
		return nil
	}
	return &Disqus{data: &comments}
}

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

func (d *Disqus) GetArticleByComment(c *Commment) *Article {

	articles := d.GetAllArticles()
	if len(articles) == 0 || c == nil {
		return nil
	}

	for _, a := range articles {
		if a.AttrID == c.Article.ID {
			return &a
		}
	}

	return nil
}
