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
	"encoding/xml"
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

	// Open our xmlFile
	xmlFile, err := os.Open(*textPtr)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var comments worker.Disqus
	if err := xml.Unmarshal(byteValue, &comments); err != nil {
		fmt.Printf("err: %s \n", err)
		os.Exit(13)
	}

	for i, c := range comments.Post {
		fmt.Printf("Post: aticle ID:%s authur:%s Msg:%s \n", c.ID, c.Author.Name, c.Message)
		if i > 5 {
			break
		}
	}

}
