package disqusimportorgo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var globalData []byte

func init() {
	globalData = getTestXML()
}
func getTestXML() []byte {
	xmlFile, err := os.Open("./example/evanlin_20210517.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	if err != nil {
		panic("fail on files")
	}
	return byteValue
}

func TestGetAllComments(t *testing.T) {
	disqus := NewDisqus(globalData)

	if disqus == nil {
		t.Fatal("Initilized error.")
	}

	comments := disqus.GetAllComments()
	if len(comments) == 0 {
		t.Fatal("Get comments error!")
	}

	//Test 1st comment's article
	article := disqus.GetArticleByComment(comments[0])
	if article == nil {
		t.Fatal("Get article error!")
	}

	if aName := article.GetAuthorName(); len(aName) == 0 {
		t.Fatal("Get article's Author error!", article)
	}

	aComments := disqus.GetAllCommentsByArticle(*article)
	if len(aComments) < 1 {
		t.Fatal("Get comment's by article error!", aComments, article)
	}

	if name := comments[0].GetAuthorName(); len(name) == 0 {
		t.Fatal("Get comment Author error!", comments[0])
	}
}

func TestGetPath(t *testing.T) {
	source := []string{"https://www.evanlin.com/reading-twitter/",
		"http://www.evanlin.com/e5-b0-87mt-e5-88-86-e9-a1-9e-e5-bd-99-e6-95-b4-e5-8a-a0-e4-bb-a5-e4-bf-ae-e6-94-b9-e4-ba-86/?__mode=view&amp;entry_id="}

	want := []string{"reading-twitter/",
		"e5-b0-87mt-e5-88-86-e9-a1-9e-e5-bd-99-e6-95-b4-e5-8a-a0-e4-bb-a5-e4-bf-ae-e6-94-b9-e4-ba-86/"}

	for i, str := range source {
		if out := getShortPath(str); out != want[i] {
			t.Fatal("Get path fail: str:", str, " ret:", out)
		}
	}
}

func TestPrepareData(t *testing.T) {
	var articles []Article
	articles = append(articles, Article{Title: "topic1", AttrID: "001", Link: "https:/aaa.ccc/001/"})
	articles = append(articles, Article{Title: "topic2", AttrID: "002", Link: "https:/aaa.ccc/002/"})
	articles = append(articles, Article{Title: "topic3", AttrID: "003", Link: "https:/aaa.ccc/003/"})

	var comments []Comment
	comments = append(comments, Comment{
		ArticleLink: ArticleLinkStruct{ID: "002"},
		Author:      AuthorStruct{Name: "John"},
		CreatedAt:   "2020-05-15T04:34:00Z",
		Message:     "msg1"})

	comments = append(comments, Comment{
		ArticleLink: ArticleLinkStruct{ID: "003"},
		Author:      AuthorStruct{Name: "Tom"},
		CreatedAt:   "2019-11-15T04:34:00Z",
		Message:     "msg2"})

	comments = append(comments, Comment{
		ArticleLink: ArticleLinkStruct{ID: "002"},
		Author:      AuthorStruct{Name: "May"},
		CreatedAt:   "2018-12-15T04:34:00Z",
		Message:     "msg3"})

	disqusObj := DisqusStruct{
		Articles:  articles,
		Commments: comments}

	disqusXML, _ := xml.Marshal(disqusObj)

	disqus := NewDisqus(disqusXML)
	if err := disqus.PrepareImportData(); err != nil {
		t.Error("err:", err)
	}

	if len(disqus.impData) != 2 {
		t.Error("PrepareImportData article count error!")
	}

	if len(disqus.impData["002/"].Comments) != 2 {
		t.Error("PrepareImportData comment count error!")
	}
}
