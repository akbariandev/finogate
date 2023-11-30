package gateway

import (
	"encoding/json"
	"errors"
	"finogate/gateway/shepa"
)

type GateWay struct {
	sandbox     bool
	transaction any
}

type IGateWay interface {
	DoPayment(handler func(any)) error
}

func New(gateWayName string, inTransaction []byte, sandBox bool) (IGateWay, error) {
	switch gateWayName {
	case shepa.GateWayName:
		t, err := newTransaction[shepa.Transaction](inTransaction)
		if err != nil {
			return nil, err
		}
		return shepa.New(t, sandBox), nil
	default:
		return nil, errors.New("gateway not implemented")
	}
}

func newTransaction[T any](inTransaction []byte) (transaction *T, err error) {
	if err = json.Unmarshal(inTransaction, &transaction); err != nil {
		return nil, err
	}

	return transaction, err
}
