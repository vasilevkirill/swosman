package swos

import (
	"fmt"
	dac "github.com/xinsnake/go-http-digest-auth-client"
	"log"
)

func NewClient(config *Config) (*Client, error) {
	config.url = fmt.Sprintf("%s://%s:%d", config.HttpProtocol, config.Host, config.Port)
	client := Client{Config: *config}
	client.Config.httpTransport = dac.NewRequest(client.Config.Username, client.Config.Password, "GET", client.Config.url, "")
	if err := client.pullLinkb(); err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *Client) testConnect() error {
	_, err := getHTTP("/get_link.b", c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) pullLinkb() error {
	b, err := getHTTP("link.b", c)
	if err != nil {
		return err
	}
	//log.Printf("%s",b)
	var l linkB
	err = unmarshalBody(&b, &l)
	if err != nil {
		return err
	}
	c.Config.linkb = l
	return nil
}

func (c *Client) pushLinkb() error {
	post := linkBPost{
		En:   c.Config.linkb.En,
		Nm:   c.Config.linkb.Nm,
		An:   c.Config.linkb.An,
		Spdc: c.Config.linkb.Spdc,
		Dpxc: c.Config.linkb.Dpxc,
		Fctc: c.Config.linkb.Fctc,
		Fctr: c.Config.linkb.Fctr,
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
