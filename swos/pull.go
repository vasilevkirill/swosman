package swos

func (c *Client) pullAll() error {
	if err := c.pullLinkb(); err != nil {
		return err
	}
	if err := c.pullLacpb(); err != nil {
		return err
	}
	if err := c.pullFwdb(); err != nil {
		return err
	}
	return nil
}

func (c *Client) pullLinkb() error {
	b, err := getHTTP("link.b", c)
	if err != nil {
		return err
	}
	var l linkB
	err = unmarshalBody(&b, &l)
	if err != nil {
		return err
	}
	c.Config.Raw.linkb = l
	c.switchLoadLink()
	return nil
}

func (c *Client) pullLacpb() error {
	b, err := getHTTP("lacp.b", c)
	if err != nil {
		return err
	}
	var l lacpB
	err = unmarshalBody(&b, &l)
	if err != nil {
		return err
	}
	c.Config.Raw.lacpb = l
	c.switchLoadLag()
	return nil
}
func (c *Client) pullFwdb() error {
	b, err := getHTTP("fwd.b", c)
	if err != nil {
		return err
	}
	var l fwdB
	err = unmarshalBody(&b, &l)
	if err != nil {
		return err
	}
	c.Config.Raw.fwdb = l
	return nil
}
