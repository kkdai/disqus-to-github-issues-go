package disqusimportorgo

import (
	x "encoding/xml"
	"fmt"
)

type Disqus struct {
	data DisqusFormat
}

func NewDisqus(xml []byte) *Disqus {
	var comments DisqusFormat
	if err := x.Unmarshal(xml, &comments); err != nil {
		fmt.Printf("err: %s \n", err)
		return nil
	}
	return &Disqus{data: comments}
}
