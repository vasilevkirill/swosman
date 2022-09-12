package swos

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func getHTTP(url string, client *Client) ([]byte, error) {
	client.Config.httpTransport.UpdateRequest(client.Config.Username, client.Config.Password, "GET", fmt.Sprintf("%s/%s", client.Config.url, url), "")
	response, err := client.Config.httpTransport.Execute()
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
func postHTTP(url string, client *Client, sendbody string) ([]byte, error) {
	client.Config.httpTransport.UpdateRequest(client.Config.Username, client.Config.Password, "POST", fmt.Sprintf("%s/%s", client.Config.url, url), sendbody)
	client.Config.httpTransport.Header.Set("Pragma", "no-cache")
	client.Config.httpTransport.Header.Set("Cache-Control", "no-cache")
	client.Config.httpTransport.Header.Set("User-Agent", "Vasilevkirill swosman")
	client.Config.httpTransport.Header.Set("Content-Type", "text/plain")
	client.Config.httpTransport.Header.Set("Cookie", "username=admin")
	response, err := client.Config.httpTransport.Execute()
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
