package swos

import (
	"log"
)

type SWInterface struct {
	Index           int
	Description     string
	Enable          bool
	AutoNegotiation bool
	StatusLink      bool
	Speed           int
	FullDuplex      bool
	FlowControl     bool
}

var speeds = [...]int{0, 10, 1000, 10000}

func interfaceParse(body *[]byte) ([]SWInterface, error) {
	var SWinterfaceList []SWInterface
	linkb, err := parseLinkB(body)
	if err != nil {
		return nil, err
	}
	log.Printf("%+v", linkb)
	for i := 0; i < int(linkb.Prt); i++ {
		var SWi SWInterface
		SWi.Index = i
		SWi.Description = linkb.Nm[i]
		SWi.Enable = getBoolFromUint32(linkb.En, i)
		SWi.StatusLink = getBoolFromUint32(linkb.Lnk, i)
		SWi.AutoNegotiation = getBoolFromUint32(linkb.An, i)
		SWi.FullDuplex = getBoolFromUint32(linkb.Dpx, i)
		SWi.Speed = speeds[linkb.Spdc[i]]
		SWi.FlowControl = getBoolFromUint32(linkb.Fctc, i)
		SWinterfaceList = append(SWinterfaceList, SWi)
	}
	return SWinterfaceList, nil
}

func parseLinkB(body *[]byte) (linkB, error) {
	var linkb linkB
	err := unmarshalBody(body, &linkb)
	if err != nil {
		return linkb, err
	}

	return linkb, nil
}
