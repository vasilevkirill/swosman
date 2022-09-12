package swos

import (
	"errors"
)

// InterfaceGetFromIndex getting interface information from index
// this uint8 starting 0
func (c *Client) InterfaceGetFromIndex(index uint8) (Interface, error) {

	var Interface Interface
	Interface.Index = index
	Interface.Enabled = getBoolFromUint32(c.Config.Raw.linkb.En, index)
	Interface.Name = c.Config.Raw.linkb.Nm[index]
	Interface.AutoNegotiation = getBoolFromUint32(c.Config.Raw.linkb.An, index)
	Interface.FlowControlTx = getBoolFromUint32(c.Config.Raw.linkb.Fctc, index)
	Interface.FlowControlRx = getBoolFromUint32(c.Config.Raw.linkb.Fctr, index)

	return Interface, nil
}

//InterfaceGetAllSlice getting all interface with slice
func (c *Client) InterfaceGetAllSlice() ([]Interface, error) {
	var x []Interface

	for i := 0; i < int(c.Config.Raw.linkb.Prt); i++ {
		swi, err := c.InterfaceGetFromIndex(uint8(i))
		if err != nil {
			return nil, err
		}
		x = append(x, swi)
	}
	return x, nil
}

func (c *Client) InterfaceSetNameFromIndex(name string, index uint8) error {
	if len(c.Config.Raw.linkb.Nm) < int(index) {
		return errors.New("index not set")
	}
	if c.Config.Raw.linkb.Nm[index] == name {
		return nil
	}
	c.Config.Raw.linkb.Nm[index] = name
	if err := c.pushLinkb(); err != nil {
		return err
	}
	return nil
}
