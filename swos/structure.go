package swos

import (
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

type Config struct {
	Host          string
	Port          uint16
	HttpProtocol  string
	Username      string
	Password      string
	url           string
	Raw           Raw
	httpTransport dac.DigestRequest
}

type Raw struct {
	linkb linkB
	lacpb lacpB
	fwdb  fwdB
	rstpb rstpB
	vlanb vlanB
	snmpb snmpB
}

type Client struct {
	Config Config
	Sw     Switch
}
