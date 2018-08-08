package dmjson

type DmBlockTransaction struct {
	Amount        string `json:"amount"`
	AuthoritySign string `json:"authority_sign"`
	Black         string `json:"black"`
	BlockNum      string `json:"blockNum"`
	Chaincode     string `json:"chaincode"`
	Error         string `json:"error"`
	Facc          string `json:"facc"`
	MinerNormal   string `json:"minerNormal"`
	MinerSuper    string `json:"minerSuper"`
	NMinerFee     string `json:"nMinerFee"`
	Remark        string `json:"remark"`
	SMinerFee     string `json:"sMinerFee"`
	Status        string `json:"status"`
	Tacc          string `json:"tacc"`
	Time          string `json:"time"`
	TxID          string `json:"txId"`
	Type          string `json:"type"`
}

type DmGetTransactionIdResutl struct {
	Txid  string
	Nonce string
}

type DmGetTransactionDetailResult struct {
	Amount        string `json:"amount"`
	AuthoritySign string `json:"authority_sign"`
	BlockNum      string `json:"blockNum"`
	Error         string `json:"error"`
	Fee           string `json:"fee"`
	In            string `json:"in"`
	Out           string `json:"out"`
	Remark        string `json:"remark"`
	Status        string `json:"status"`
	Sum           string `json:"sum"`
	Time          string `json:"time"`
	Token         string `json:"token"`
	TxID          string `json:"txId"`
	Type          string `json:"type"`
}

type DmGetNewAddressResult struct {
	Address string       `json:"address"`
	Crypto  DmCryptoInfo `json:"crypto"`
}

type DmCryptoInfo struct {
	Cipher     string `json:"cipher"`
	CipherText string `json:"ciphertext"`
}

type DmTokenInfo struct {
	Name        string `json:"name"`
	Cc          string `json:"cc"`
	Creator     string `json:"creator"`
	Desc        string `json:"desc"`
	Total       string `json:"total"`
	URL         string `json:"url"`
	Decimal     string `json:"decimal"`
	Logo        string `json:"logo"`
	Mineral     string `json:"mineral"`
	Award       string `json:"award"`
	Email       string `json:"email"`
	TokenCharge string `json:"tokenCharge"`
	Charge      string `json:"charge"`
	PublishTime string `json:"publishTime"`
	Deflation   string `json:"deflation"`
}
