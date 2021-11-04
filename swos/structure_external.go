package swos

type Interface struct {
	Index           uint8  `yaml:"Id" json:"Id"`
	Enabled         bool   `yaml:"AdminStatus" json:"AdminStatus"`
	Name            string `yaml:"Desription" json:"Desription"`
	AutoNegotiation bool   `yaml:"AutoNegotiation" json:"AutoNegotiation"`
	FlowControlTx   bool   `yaml:"FlowControlTx" json:"FlowControlTx"`
	FlowControlRx   bool   `yaml:"FlowControlRx" json:"FlowControlRx"`
}
