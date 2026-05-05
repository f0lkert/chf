package cdrType

import "github.com/f0lkert/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type QoSCharacteristics struct {
	Value asn.OctetString
}
