package disqusimportorgo

import "encoding/xml"

// Disqus: Disqus comment export structure in go
type DisqusStruct struct {
	XMLName        xml.Name       `xml:"disqus"`
	Text           string         `xml:",chardata"`
	Xmlns          string         `xml:"xmlns,attr"`
	Dsq            string         `xml:"dsq,attr"`
	Xsi            string         `xml:"xsi,attr"`
	SchemaLocation string         `xml:"schemaLocation,attr"`
	Category       CategoryStruct `xml:"category"`
	Articles       []Article      `xml:"thread"`
	Commments      []Comment      `xml:"post"`
}

type CategoryStruct struct {
	Text      string `xml:",chardata"`
	ID        string `xml:"id,attr"`
	Forum     string `xml:"forum"`
	Title     string `xml:"title"`
	IsDefault string `xml:"isDefault"`
}

type Comment struct {
	Text        string            `xml:",chardata"`
	AttrID      string            `xml:"id,attr"`
	ID          string            `xml:"id"`
	Message     string            `xml:"message"`
	CreatedAt   string            `xml:"createdAt"`
	IsDeleted   string            `xml:"isDeleted"`
	IsSpam      string            `xml:"isSpam"`
	Author      AuthorStruct      `xml:"author"`
	ArticleLink ArticleLinkStruct `xml:"thread"`
	Parent      struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
	} `xml:"parent"`
}

type ArticleLinkStruct struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
}
type Article struct {
	Text     string `xml:",chardata"`
	AttrID   string `xml:"id,attr"`
	ID       string `xml:"id"`
	Forum    string `xml:"forum"`
	Category struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
	} `xml:"category"`
	Link      string       `xml:"link"`
	Title     string       `xml:"title"`
	Message   string       `xml:"message"`
	CreatedAt string       `xml:"createdAt"`
	Author    AuthorStruct `xml:"author"`
	IsClosed  string       `xml:"isClosed"`
	IsDeleted string       `xml:"isDeleted"`
}

type AuthorStruct struct {
	Text        string `xml:",chardata"`
	Name        string `xml:"name"`
	IsAnonymous string `xml:"isAnonymous"`
	Username    string `xml:"username"`
}

func (a Article) GetAuthorName() string {
	return a.Author.Name
}

func (a Article) GetArticleLink() string {
	return a.Link
}

func (c Comment) GetAuthorName() string {
	return c.Author.Name
}
