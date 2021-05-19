package disqusimportorgo

import (
	x "encoding/xml"
	"fmt"
)

type Disqus struct {
	data DisqusStruct
}

func NewDisqus(xml []byte) *Disqus {
	var comments DisqusStruct
	if err := x.Unmarshal(xml, &comments); err != nil {
		fmt.Printf("err: %s \n", err)
		return nil
	}
	return &Disqus{data: comments}
}

func (d *Disqus) GetAllComments() []Commment {
	return d.data.Commments
}
