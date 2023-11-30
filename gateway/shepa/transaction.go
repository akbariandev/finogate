package shepa

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Transaction struct {
	Api         string `json:"api"`
	Amount      int    `json:"amount"`
	Callback    string `json:"callback"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	CardNumber  string `json:"cardnumber"`
	Description string `json:"description"`
}

type transactionResult struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

type transactionResponse struct {
	Success string            `json:"success"`
	Result  transactionResult `json:"result"`
	Errors  []string          `json:"errors"`
}

func (t *Transaction) sendTransaction() (*transactionResponse, error) {
	u := url.URL{Host: _base_api, Path: _send_transaction}
	var bodyBuff bytes.Buffer
	if err := json.NewEncoder(&bodyBuff).Encode(t); err != nil {
		return nil, err
	}

	httpClient := http.Client{}
	r, err := http.NewRequest(http.MethodPost, u.Path, &bodyBuff)
	r.Header.Add("Content-Type", "application/json")
	jsonResp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	response := new(transactionResponse)
	if err = json.NewDecoder(jsonResp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
