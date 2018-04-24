package rpcclient

import (
	"encoding/json"
	"github.com/icloudland/btcdx/chaincfg/chainhash"
	"github.com/icloudland/btcdx/btcjson"
	"github.com/icloudland/btcutil"
	"github.com/icloudland/btcdx/tcoinjson"
)

func (c *Client) GetBlockVerboseAsyncT(blockHash *chainhash.Hash) FutureGetBlockVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := btcjson.NewGetBlockCmd(hash, btcjson.Bool(true), nil)
	return c.sendCmd(cmd)
}

func (c *Client) GetBlockVerboseT(blockHash *chainhash.Hash) (*GetBlockVerboseResult, error) {
	return c.GetBlockVerboseAsyncT(blockHash).ReceiveT()
}

func (r FutureGetBlockVerboseResult) ReceiveT() (*GetBlockVerboseResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a BlockResult.
	var blockResult GetBlockVerboseResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}
	return &blockResult, nil
}

type GetBlockVerboseResult struct {
	Tx []TxRawResult `json:"tx,omitempty"`
}

type TxRawResult struct {
	Vout []Vout `json:"vout"`
}

type Vout struct {
	Value        float64            `json:"value"`
	N            uint32             `json:"n"`
	ScriptPubKey ScriptPubKeyResult `json:"scriptPubKey"`
}

type ScriptPubKeyResult struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex,omitempty"`
	ReqSigs   int32    `json:"reqSigs,omitempty"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses,omitempty"`
}

func (c *Client) GetBlockT(blockHash *chainhash.Hash) (*GetBlockeResult, error) {
	return c.GetBlockAsyncT(blockHash).ReceiveT()
}

func (c *Client) GetBlockAsyncT(blockHash *chainhash.Hash) FutureGetBlockResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	cmd := btcjson.NewGetBlockCmd(hash, btcjson.Bool(false), nil)
	return c.sendCmd(cmd)
}

func (r FutureGetBlockResult) ReceiveT() (*GetBlockeResult, error) {

	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var blockResult GetBlockeResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}
	return &blockResult, nil

}

type GetBlockeResult struct {
	Hash          string   `json:"hash"`
	Confirmations uint64   `json:"confirmations"`
	StrippedSize  int32    `json:"strippedsize"`
	Size          int32    `json:"size"`
	Weight        int32    `json:"weight"`
	Height        int64    `json:"height"`
	Version       int32    `json:"version"`
	VersionHex    string   `json:"versionHex"`
	MerkleRoot    string   `json:"merkleroot"`
	Tx            []string `json:"tx,omitempty"`
	Time          int64    `json:"time"`
	Nonce         uint32   `json:"nonce"`
	Bits          string   `json:"bits"`
	Difficulty    float64  `json:"difficulty"`
	PreviousHash  string   `json:"previousblockhash"`
	NextHash      string   `json:"nextblockhash,omitempty"`
}

//func NewCreateRawTransactionCmd(inputs []btcjson.TransactionInput, amounts map[string]float64) *CreateRawTransactionCmd {
//
//	return &CreateRawTransactionCmd{
//		Inputs:  inputs,
//		Amounts: amounts,
//	}
//}
//
//type CreateRawTransactionCmd struct {
//	Inputs  []btcjson.TransactionInput
//	Amounts map[string]float64 `jsonrpcusage:"{\"address\":amount,...}"` // In BTC
//}

func (c *Client) CreateRawTransactionAsyncT(inputs []btcjson.TransactionInput,
	amounts map[string]btcutil.Amount) FutureCreateRawTransactionResult {

	convertedAmts := make(map[string]float64, len(amounts))
	for addr, amount := range amounts {
		convertedAmts[addr] = amount.ToBTC()
	}
	cmd := btcjson.NewCreateRawTransactionCmd(inputs, convertedAmts, nil)
	return c.sendCmd(cmd)
}

// CreateRawTransaction returns a new transaction spending the provided inputs
// and sending to the provided addresses.
func (c *Client) CreateRawTransactionT(inputs []btcjson.TransactionInput,
	amounts map[string]btcutil.Amount, lockTime *int64) (string, error) {

	return c.CreateRawTransactionAsyncT(inputs, amounts).ReceiveT()
}

func (r FutureCreateRawTransactionResult) ReceiveT() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}
	var txHex string
	err = json.Unmarshal(res, &txHex)
	if err != nil {
		return "", err
	}

	return txHex, nil
}

//func init() {
//	btcjson.UnRegisterCmd("createrawtransaction")
//	btcjson.MustRegisterCmd("createrawtransaction", (*CreateRawTransactionCmd)(nil), btcjson.UsageFlag(0))
//
//}

func (c *Client) SignRawTransaction4T(tx string,
	inputs []btcjson.RawTxInput, privKeysWIF []string,
	hashType SigHashType) (string, bool, error) {

	return c.SignRawTransaction4AsyncT(tx, inputs, privKeysWIF,
		hashType).ReceiveT()
}

func (c *Client) SignRawTransaction4AsyncT(tx string,
	inputs []btcjson.RawTxInput, privKeysWIF []string,
	hashType SigHashType) FutureSignRawTransactionResult {

	txHex := tx

	cmd := btcjson.NewSignRawTransactionCmd(txHex, &inputs, &privKeysWIF,
		btcjson.String(string(hashType)))
	return c.sendCmd(cmd)
}

func (r FutureSignRawTransactionResult) ReceiveT() (string, bool, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", false, err
	}

	// Unmarshal as a signrawtransaction result.
	var signRawTxResult btcjson.SignRawTransactionResult
	err = json.Unmarshal(res, &signRawTxResult)
	if err != nil {
		return "", false, err
	}

	return signRawTxResult.Hex, signRawTxResult.Complete, nil
}

func (c *Client) SendRawTransactionAsyncT(txHex string, allowHighFees bool) FutureSendRawTransactionResult {

	cmd := btcjson.NewSendRawTransactionCmd(txHex, nil)
	return c.sendCmd(cmd)
}

// SendRawTransaction submits the encoded transaction to the server which will
// then relay it to the network.
func (c *Client) SendRawTransactionT(txHex string, allowHighFees bool) (string, error) {
	return c.SendRawTransactionAsyncT(txHex, allowHighFees).ReceiveT()
}

func (r FutureSendRawTransactionResult) ReceiveT() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return "", err
	}

	return txHashStr, nil
}

func (r FutureGetNewAddressResult) ReceiveD() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var addr string
	err = json.Unmarshal(res, &addr)
	if err != nil {
		return "", err
	}

	return addr, err
}

// FutureGetNewAddressResult is a future promise to deliver the result of a
// GetNewAddressAsync RPC invocation (or an applicable error).
type FutureGetNewAddressAndKeyResult chan *response

// Receive waits for the response promised by the future and returns a new
// address.
func (r FutureGetNewAddressAndKeyResult) Receive() (string, string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", "", err
	}

	type AddrKey struct {
		Address string
		Secret  string
	}
	// Unmarshal result as a AddrKey.
	var addrKey AddrKey
	err = json.Unmarshal(res, &addrKey)
	if err != nil {
		return "", "", err
	}

	return addrKey.Address, addrKey.Secret, nil
}

// GetNewAddressAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetNewAddress for the blocking version and more details.
func (c *Client) GetNewAddressAndKeyAsync() FutureGetNewAddressAndKeyResult {
	cmd := tcoinjson.NewGetNewAddressAndKeyCmd()
	return c.sendCmd(cmd)
}

// GetNewAddress returns a new address.
func (c *Client) GetNewAddressAndKey() (string, string, error) {
	return c.GetNewAddressAndKeyAsync().Receive()
}
