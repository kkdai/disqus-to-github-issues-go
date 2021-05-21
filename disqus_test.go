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
