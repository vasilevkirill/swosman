package swos

import (
	"fmt"
	"log"
)

func (c *Client) pushLinkb() error {
	post := linkBPost{
		En:   c.Config.Raw.linkb.En,
		Nm:   c.Config.Raw.linkb.Nm,
		An:   c.Config.Raw.linkb.An,
		Spdc: c.Config.Raw.linkb.Spdc,
		Dpxc: c.Config.Raw.linkb.Dpxc,
		Fctc: c.Config.Raw.linkb.Fctc,
		Fctr: c.Config.Raw.linkb.Fctr,
	}
	body, err := marshalToBody(post)
	if err != nil {
		return err
	}
	log.Printf("%s", body)
	body1, err := postHTTP("link.b", c, fmt.Sprintf("%s", body))
	if err != nil {
		return err
	}
	log.Printf("%s", body1)

	return nil
}
