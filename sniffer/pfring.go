package sniffer

import (
	"github.com/jgallartm/gopacket/pfring"
	//"golang.org/x/net/bpf"
	//"time"
)

type pfringHandle struct {
	PFRing *pfring.Ring
}
