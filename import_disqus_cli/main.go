// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	worker "github.com/kkdai/disqus-importor-go"
)

func main() {
	textPtr := flag.String("f", "../example/evanlin_20210517.xml", "xml file address (e.g. `../example/evanlin_20210517.xml`)")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	xmlFile, err := os.Open(*textPtr)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	disqus := worker.NewDisqus(byteValue)
	if disqus == nil {
		fmt.Println("XML parsing failed.")
		return
	}

	comments := disqus.GetAllComments()
	for i, c := range comments {
		article := disqus.GetArticleByComment(c)
		fmt.Printf("Post: aticle ID:%s authur:%s Msg:%s  title:%s \n", c.ID, c.GetAuthorName(), c.Message, article.Title)
		if i > 5 {
			break
		}
	}
}
