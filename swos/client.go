package swos

import (
	"fmt"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

func NewClient(config *Config) (*Client, error) {
	config.url = fmt.Sprintf("%s://%s:%d", config.HttpProtocol, config.Host, config.Port)
	client := Client{Config: *config}
	client.Config.httpTransport = dac.NewRequest(client.Config.Username, client.Config.Password, "GET", client.Config.url, "")
	clientInit()
	if err := client.pullAll(); err != nil {
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

func clientInit() {
	lacpMap[0] = "passive"
	lacpMap[1] = "active"
	lacpMap[2] = "static"
}
