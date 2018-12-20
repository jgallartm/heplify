package sniffer

import (
	"github.com/jgallartm/pfring"
	"golang.org/x/net/bpf"
	"time"
)

type pfringHandle struct {
	PFRing *pfring.Ring
}
