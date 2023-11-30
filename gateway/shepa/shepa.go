package shepa

import (
	"errors"
	"finogate/gateway"
	"net/http"
)

const (
	GateWayName  = "shepa"
	_base_api    = "https://merchant.shepa.com"
	_sandbox_api = "https://sandbox.shepa.com"

	_send_transaction = "api/v1/token"
)

type ShepaGateWay struct {
	sandBox     bool
	transaction *Transaction
}

func New(transaction *Transaction, sandBox bool) gateway.IGateWay {
	return &ShepaGateWay{transaction: transaction, sandBox: sandBox}
}

func (s ShepaGateWay) DoPayment(transactionHandler func(transaction any)) error {
	// send transaction to SHEPA GateWay
	tr, err := s.transaction.sendTransaction()
	if err != nil {
		return err
	}

	transactionHandler(tr)
	if tr.Success == "true" {
		http.RedirectHandler(tr.Result.Url, http.StatusMovedPermanently)
		return nil
	}

	return errors.New("transaction failed")
}
