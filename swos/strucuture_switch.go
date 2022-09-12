package swos

type Switch struct {
	Links          []Link          `yaml:"Links" json:"Links"`
	PortIsolations []PortIsolation `yaml:"PortIsolations" json:"PortIsolations"`
	Lags           []Lag           `yaml:"Lags" json:"Lags"`
	Forwardings    []Forwarding    `yaml:"Forwardings" json:"Forwardings"`
	Rstp           Rstp            `yaml:"Rstp" json:"Rstp"`
	Vlan           []Vlan          `yaml:"Vlan" json:"Vlan"`
	Vlans          []Vlans         `yaml:"Vlans" json:"Vlans"`
	Hosts          []Hosts         `yaml:"Hosts" json:"Hosts"`
	Snmp           Snmp            `yaml:"Snmp" json:"Snmp"`
}

type Link struct {
	Index           uint8  `yaml:"Id" json:"Id"`
	Enabled         bool   `yaml:"AdminStatus" json:"AdminStatus"`
	Name            string `yaml:"Description" json:"Description"`
	AutoNegotiation bool   `yaml:"AutoNegotiation" json:"AutoNegotiation"`
	FlowControlTx   bool   `yaml:"FlowControlTx" json:"FlowControlTx"`
	FlowControlRx   bool   `yaml:"FlowControlRx" json:"FlowControlRx"`
}
type Lag struct {
	Index uint8  `yaml:"Id" json:"Id"`
	Mode  string `yaml:"Mode" json:"Mode"`
	Group uint8  `yaml:"Group" json:"Group"`
}
type PortIsolation struct {
	Index uint8   `yaml:"Id" json:"Id"`
	Ports []uint8 `yaml:"Ports" json:"Ports"`
}
type Forwarding struct {
	Index                 uint8 `yaml:"Id" json:"Id"`
	PortLock              bool  `yaml:"PortLock" json:"PortLock"`
	LockOnFirst           bool  `yaml:"LockOnFirst" json:"LockOnFirst"`
	MirrorIngress         bool  `yaml:"MirrorIngress" json:"MirrorIngress"`
	MirrorEgress          bool  `yaml:"MirrorEgress" json:"MirrorEgress"`
	MirrorTo              uint8 `yaml:"MirrorTo" json:"MirrorTo"`
	StormRatePercent      int   `yaml:"StormRatePercent" json:"StormRatePercent"`
	LimitUnknownUnicast   bool  `yaml:"LimitUnknownUnicast" json:"LimitUnknownUnicast"`
	FloodUnknownMulticast bool  `yaml:"FloodUnknownMulticast" json:"FloodUnknownMulticast"`
	IngressRate           int   `yaml:"IngressRate" json:"IngressRate"`
}
type Rstp struct {
	General RstpGeneral `yaml:"General" json:"General"`
	Ports   []RstPort   `yaml:"Ports" json:"Ports"`
}
type RstpGeneral struct {
	BridgePriority int    `yaml:"BridgePriority" json:"BridgePriority"`
	PortCostMode   string `yaml:"PortCostMode" json:"PortCostMode"`
}
type RstPort struct {
	Index  uint8 `yaml:"Id" json:"Id"`
	Enable bool  `yaml:"Enable" json:"Enable"`
}
type Vlan struct {
	Index         uint8  `yaml:"Id" json:"Id"`
	VlanMode      string `yaml:"VlanMode" json:"VlanMode"`
	VlanReceive   string `yaml:"VlanReceive" json:"VlanReceive"`
	DefaultVlanId int    `yaml:"DefaultVlanId" json:"DefaultVlanId"`
	ForceVlanId   int    `yaml:"ForceVlanId" json:"ForceVlanId"`
}

type Vlans struct {
	VlanId        int    `yaml:"VlanId" json:"VlanId"`
	Name          string `yaml:"Name" json:"Name"`
	PortIsolation bool   `yaml:"PortIsolation" json:"PortIsolation"`
	Learning      bool   `yaml:"Learning" json:"Learning"`
	Mirror        bool   `yaml:"Mirror" json:"Mirror"`
	IgmpSnooping  bool   `yaml:"IgmpSnooping" json:"IgmpSnooping"`
	Members       []int  `yaml:"Members" json:"Members"`
}

type Hosts struct {
	Port   int    `yaml:"Enable" json:"Enable"`
	Mac    string `yaml:"Mac" json:"Mac"`
	VlanId int    `yaml:"VlanId" json:"VlanId"`
	Drop   bool   `yaml:"Drop" json:"Drop"`
	Mirror bool   `yaml:"Mirror" json:"Mirror"`
}

type Snmp struct {
	Enabled     bool   `yaml:"Enable" json:"Enable"`
	Community   string `yaml:"Community" json:"Community"`
	ContactInfo string `yaml:"ContactInfo" json:"ContactInfo"`
	Location    string `yaml:"Location" json:"Location"`
}
