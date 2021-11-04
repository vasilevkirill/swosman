package swos

func (c *Client) InterfaceGetAll() ([]SWInterface, error) {
	body, err := getHTTP("/link.b", c)
	if err != nil {
		return nil, err
	}
	listInterfaces, err := interfaceParse(&body)
	if err != nil {
		return nil, err
	}
	c.Interfaces = listInterfaces
	return c.Interfaces, nil
}
