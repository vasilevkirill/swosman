package swos

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bobziuchkovski/digest"
)

type Config struct {
	Url      string
	UserName string
	Password string
	Timeout  time.Duration
}
type Client struct {
	Config     Config
	Interfaces []SWInterface
}

func Dial(config *Config) (*Client, error) {
	client := NewClient(config)
	if err := client.TestConnect(); err != nil {
		return nil, err
	}
	return client, nil
}

func NewClient(config *Config) *Client {
	client := Client{Config: *config}

	return &client
}

func (c *Client) TestConnect() error {
	_, err := getHTTP("/link.b", c)
	if err != nil {
		return err
	}
	return nil
}
func getHTTP(url string, c *Client) ([]byte, error) {
	return sendHTTP(url, "GET", c)
}
func postHTTP(url string, c *Client) ([]byte, error) {
	return sendHTTP(url, "POST", c)
}
func sendHTTP(url string, method string, client *Client) ([]byte, error) {
	digestNewTransport := digest.NewTransport(client.Config.UserName, client.Config.Password)
	httpClient := &http.Client{
		Transport:     digestNewTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       client.Config.Timeout,
	}
	if method == "" {
		method = "GET"
	}
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", client.Config.Url, url), nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("Bad Status:" + response.Status)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
