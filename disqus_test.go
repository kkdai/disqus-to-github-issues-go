package disqusimportorgo

import (
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

	aComment := disqus.GetAllCommentsByArticle(*article)
	if len(aComment) < 1 {
		t.Fatal("Get comment by article error!", aComment, comments[0])
	}
}
