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
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Open our xmlFile
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

	// we initialize our Users array
	var comments Disqus
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	if err := xml.Unmarshal(byteValue, &comments); err != nil {
		fmt.Printf("err: %s \n", err)
		os.Exit(13)
	}

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i, c := range comments.Post {
		fmt.Printf("Post %n : authur:%s Msg:%s \n", i, c.Author.Name, c.Message)
	}

}
