package swos

func (c *Client) switchLoadLink() {
	var links []Link
	for i := 0; i < int(c.Config.Raw.linkb.Prt); i++ {
		var l Link
		index := uint8(i)
		l.Index = index
		l.Enabled = getBoolFromUint32(c.Config.Raw.linkb.En, index)
		l.Name = c.Config.Raw.linkb.Nm[index]
		l.AutoNegotiation = getBoolFromUint32(c.Config.Raw.linkb.An, index)
		l.FlowControlTx = getBoolFromUint32(c.Config.Raw.linkb.Fctc, index)
		l.FlowControlRx = getBoolFromUint32(c.Config.Raw.linkb.Fctr, index)

		links = append(links, l)
	}

	c.Sw.Links = links
}

func (c *Client) switchLoadLag() {
	var lags []Lag
	for i := 0; i < int(c.Config.Raw.linkb.Prt); i++ {
		var l Lag
		index := uint8(i)
		l.Index = index
		l.Mode = string(c.Config.Raw.lacpb.Mode[i])
		l.Group = c.Config.Raw.lacpb.Grp[i]
		lags = append(lags, l)
	}

	c.Sw.Lags = lags
}
