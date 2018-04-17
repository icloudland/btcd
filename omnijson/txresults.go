package omnijson

type OmniGetTransactionResult struct {
	TxId             string `json:"txid"`
	SendingAddress   string `json:"sendingaddress"`
	ReferenceAddress string `json:"referenceaddress"`
	IsMine           bool   `json:"ismine"`
	Confirmations    int64  `json:"confirmations"`
	Fee              string `json:"fee"`
	BlockTime        int64  `json:"blocktime"`
	Valid            bool   `json:"valid"`
	InvalidReason    string `json:"invalidreason"`
	Version          int    `json:"version"`
	TypeInt          int    `json:"type_int"`
	Type             string `json:"type"`
	PropertyId       int    `json:"propertyid"`
	Amount           string `json:"amount"`
}
