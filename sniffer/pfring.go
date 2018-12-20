package sniffer

import (
	"github.com/google/jgallartm/pfring"
	"golang.org/x/net/bpf"
	"time"
)

type pfringHandle struct {
	PFRing *pfring.Ring
}
