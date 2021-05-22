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
	filePtr := flag.String("f", "", "xml file address (e.g. `../DISQUS_EXPORT.xml`)")
	token := flag.String("t", "", "github token")
	user := flag.String("u", "", "github user name")
	repo := flag.String("r", "", "github import repo name")

	flag.Parse()

	if *filePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	xmlFile, err := os.Open(*filePtr)
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

	if err := disqus.PrepareImportData(); err != nil {
		fmt.Println("Prepare data error:", err)
		return
	}

	if *token == "" || *repo == "" || *user == "" {
		fmt.Println("Your disqus xml has ", len(disqus.GetAllComments()), " comments in ", disqus.GetAllImportCommentArticle())
		fmt.Println("Please input github related option to import into github")
		flag.PrintDefaults()
		return
	}

	if err := disqus.PostToGithubIssue(*user, *repo, *token); err != nil {
		fmt.Println("PostToGithubIssue error:", err)
		return
	}
	fmt.Println("Import into github done! repo=", *repo)
}
