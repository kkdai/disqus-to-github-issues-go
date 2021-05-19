package disqusimportorgo

import "encoding/xml"

// Disqus: Disqus comment export structure in go
type Disqus struct {
	XMLName        xml.Name `xml:"disqus"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Dsq            string   `xml:"dsq,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Category       struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		Forum     string `xml:"forum"`
		Title     string `xml:"title"`
		IsDefault string `xml:"isDefault"`
	} `xml:"category"`
	Thread []struct {
		Text     string `xml:",chardata"`
		AttrID   string `xml:"id,attr"`
		ID       string `xml:"id"`
		Forum    string `xml:"forum"`
		Category struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"category"`
		Link      string `xml:"link"`
		Title     string `xml:"title"`
		Message   string `xml:"message"`
		CreatedAt string `xml:"createdAt"`
		Author    struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			IsAnonymous string `xml:"isAnonymous"`
			Username    string `xml:"username"`
		} `xml:"author"`
		IsClosed  string `xml:"isClosed"`
		IsDeleted string `xml:"isDeleted"`
	} `xml:"thread"`
	Post []struct {
		Text      string `xml:",chardata"`
		AttrID    string `xml:"id,attr"`
		ID        string `xml:"id"`
		Message   string `xml:"message"`
		CreatedAt string `xml:"createdAt"`
		IsDeleted string `xml:"isDeleted"`
		IsSpam    string `xml:"isSpam"`
		Author    struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			IsAnonymous string `xml:"isAnonymous"`
			Username    string `xml:"username"`
		} `xml:"author"`
		Thread struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"thread"`
		Parent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"parent"`
	} `xml:"post"`
}
