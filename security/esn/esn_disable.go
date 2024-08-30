package esn

import "github.com/free5gc/ike/message"

const string_ESN_DISABLE string = "ESN_DISABLE"

func toString_ESN_DISABLE(attrType uint16, intValue uint16, bytesValue []byte) string {
	return string_ESN_DISABLE
}

var _ ESNType = &ESN_DISABLE{}

type ESN_DISABLE struct {
	priority uint32
}

func (t *ESN_DISABLE) transformID() uint16 {
	return message.ESN_DISABLE
}

func (t *ESN_DISABLE) getAttribute() (bool, uint16, uint16, []byte) {
	return false, 0, 0, nil
}

func (t *ESN_DISABLE) Init() bool {
	return false
}
