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
	linkb         linkB
	httpTransport dac.DigestRequest
}

type Client struct {
	Config Config
}

type linkB struct {
	Comb uint32
	Qsfp uint32
	En   uint32
	Blkp uint32
	An   uint32
	Dpxc uint32
	Fctc uint32
	Fctr uint32
	Lnk  uint32
	Dpx  uint32
	Tfct uint32
	Rfct uint32
	Paus uint32
	Spd  []uint8
	Spdc []uint8
	Cm   []uint8
	Qtyp []uint8
	Prt  uint8
	Sfp  uint8
	Sfpo uint8
	Nm   []string
	Hop  []uint8
	Hops []uint8
	Len  []uint8
	Flt  []uint8
	Pair []uint16
}

type linkBPost struct {
	En   uint32
	Nm   []string
	An   uint32
	Spdc []uint8
	Dpxc uint32
	Fctc uint32
	Fctr uint32
}

type sfpB struct {
	Vnd []string
	Pnr []string
	Rev []string
	Ser []string
	Dat []string
	Typ []string
	Wln []uint32
	Tmp []uint32
	Vcc []uint16
	Tbs []uint16
	Tpw []uint16
	Rpw []uint16
}

type fwdB struct {
	Fp1  uint32
	Fp2  uint32
	Fp3  uint32
	Fp4  uint32
	Fp5  uint32
	Fp6  uint32
	Fp7  uint32
	Fp8  uint32
	Fp9  uint32
	Fp10 uint32
	Fp11 uint32
	Fp12 uint32
	Fp13 uint32
	Fp14 uint32
	Fp15 uint32
	Fp16 uint32
	Fp17 uint32
	Fp18 uint32
	Fp19 uint32
	Fp20 uint32
	Fp21 uint32
	Fp22 uint32
	Fp23 uint32
	Fp24 uint32
	Fp25 uint32
	Fp26 uint32
	Lck  uint32
	Lckf uint32
	Imr  uint32
	Omr  uint32
	Mrto uint32
	Vlan []uint8
	Vlni []uint8
	Dvid []uint16
	Fvid uint32
	Srt  []uint8
	Suni uint32
	Fmc  uint32
	Ir   []uint32
}

type lacpB struct {
	Mode []uint8
	Grp  []uint8
	Sgrp []uint8
	Mac  []string
}

type rstpB struct {
	Ena  uint32
	Role []uint8
	Rpc  []uint32
	Cst  []uint32
	Rstp uint32
	P2p  uint32
	Edge uint32
	Lrn  uint32
	Fwd  uint32
}

type statsB struct {
	Rb   []uint32
	Rbh  []uint32
	Tur  []uint32
	Rup  []uint32
	Tdf  []uint32
	Rbp  []uint32
	Rmp  []uint32
	P64  []uint32
	P65  []uint32
	P128 []uint32
	P256 []uint32
	P512 []uint32
	P1k  []uint32
	Tb   []uint32
	Tbh  []uint32
	Tup  []uint32
	Tec  []uint32
	Tmp  []uint32
	Tbp  []uint32
	Tmc  []uint32
	Tpp  []uint32
	Rpp  []uint32
	Rov  []uint32
	Rr   []uint32
	Fr   []uint32
	Rae  []uint32
	Rte  []uint32
	Rfcs []uint32
	Tcl  []uint32
	Tlc  []uint32
	Rtp  []uint32
	Ttp  []uint32
	Rrb  []uint32
	Trb  []uint32
	Rrp  []uint32
	Trp  []uint32
	Ruph []uint32
	Rmph []uint32
	Rbph []uint32
	Tuph []uint32
	Tmph []uint32
	Tbph []uint32
}
type vlanB struct {
	//todo надо подумать как парсить, javascriptObject
}

type snmpB struct {
	En  uint8
	Com string
	Ci  string
	Loc string
}

type sysB struct {
	Upt  uint32
	Cip  uint32
	Mac  string
	Sid  string
	Id   string
	Ver  string
	Brd  string
	Bld  uint32
	Wdt  uint8
	Dsc  uint8
	Ivl  uint8
	Alla uint32
	Allm uint8
	Allp uint32
	Avln uint16
	Temp uint32
	Btmp uint32
	Fan1 uint32
	Fan2 uint32
	Fan3 uint32
	Fan4 uint32
	P1v  uint32
	P2v  uint32
	P1c  uint32
	P2c  uint32
	P1s  uint8
	P2s  uint8
	Prio uint16
	Cost uint8
	Rpr  uint16
	Rmac string
	Igmp uint8
	Ip   uint32
	Iptp uint8
	Dtrp uint32
	Ainf uint8
	Upgr uint8
	Poe  uint8
	Poes uint8
}
