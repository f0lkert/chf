package cdrType

import "github.com/f0lkert/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type SubscriberEquipmentNumber struct { /* Set Type */
	SubscriberEquipmentNumberType SubscriberEquipmentType `ber:"tagNum:0"`
	SubscriberEquipmentNumberData asn.OctetString         `ber:"tagNum:1"`
}
