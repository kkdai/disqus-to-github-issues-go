package disqusimportorgo

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func getTestXML() []byte {
	xmlFile, err := os.Open("./example/evanlin_20210517.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	if err != nil {
		panic("fail on files")
	}
	return byteValue
}

func TestGetAllComments(t *testing.T) {
	data := getTestXML()
	disqus := NewDisqus(data)

	if disqus == nil {
		t.Fatal("Initilized error.")
	}

	comments := disqus.GetAllComments()
	if len(comments) == 0 {
		t.Fatal("Get comments error!")
	}
}
