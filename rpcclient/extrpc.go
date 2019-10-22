package rpcclient

import (
	"encoding/json"
	"github.com/icloudland/btcdx/btcjson"
	"github.com/icloudland/btcdx/chaincfg/chainhash"
	"github.com/icloudland/btcdx/tcoinjson"
	"github.com/icloudland/btcutil"
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

// ValidateAddressAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See ValidateAddress for the blocking version and more details.
func (c *Client) ValidateAddressTAsync(address string) FutureValidateAddressResult {
	cmd := btcjson.NewValidateAddressCmd(address)
	return c.sendCmd(cmd)
}

// ValidateAddress returns information about the given bitcoin address.
func (c *Client) ValidateAddressT(address string) (*btcjson.ValidateAddressWalletResult, error) {
	return c.ValidateAddressTAsync(address).Receive()
}

func (r FutureGetNewAddressResult) ReceiveT() (string, error) {
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

	return addr, nil
}

// GetNewAddress returns a new address.
func (c *Client) GetNewAddressT(account string) (string, error) {
	return c.GetNewAddressAsync(account).ReceiveT()
}

func (r FutureDumpPrivKeyResult) ReceiveT() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var privKeyWIF string
	err = json.Unmarshal(res, &privKeyWIF)
	if err != nil {
		return "", err
	}

	return privKeyWIF, nil
}

// DumpPrivKeyAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See DumpPrivKey for the blocking version and more details.
func (c *Client) DumpPrivKeyTAsync(address string) FutureDumpPrivKeyResult {
	cmd := btcjson.NewDumpPrivKeyCmd(address)
	return c.sendCmd(cmd)
}

// DumpPrivKey gets the private key corresponding to the passed address encoded
// in the wallet import format (WIF).
//
// NOTE: This function requires to the wallet to be unlocked.  See the
// WalletPassphrase function for more details.
func (c *Client) DumpPrivKeyT(address string) (string, error) {
	return c.DumpPrivKeyTAsync(address).ReceiveT()
}

// Receive waits for the response promised by the future and returns the list of
// addresses associated with the passed account.
func (r FutureGetAddressesByAccountResult) ReceiveT() ([]string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmashal result as an array of string.
	var addrStrings []string
	err = json.Unmarshal(res, &addrStrings)
	if err != nil {
		return nil, err
	}

	return addrStrings, nil
}

// GetAddressesByAccountAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetAddressesByAccount for the blocking version and more details.
func (c *Client) GetAddressesByAccountTAsync(account string) FutureGetAddressesByAccountResult {
	cmd := btcjson.NewGetAddressesByAccountCmd(account)
	return c.sendCmd(cmd)
}

// GetAddressesByAccount returns the list of addresses associated with the
// passed account.
func (c *Client) GetAddressesByAccountT(account string) ([]string, error) {
	return c.GetAddressesByAccountTAsync(account).ReceiveT()
}

func (c *Client) ListUnspentMinMaxAddressesT(minConf, maxConf int, addrs []string) ([]btcjson.ListUnspentResult, error) {
	return c.ListUnspentMinMaxAddressesTAsync(minConf, maxConf, addrs).Receive()
}

func (c *Client) ListUnspentMinMaxAddressesTAsync(minConf, maxConf int, addrs []string) FutureListUnspentResult {
	cmd := btcjson.NewListUnspentCmd(&minConf, &maxConf, &addrs)
	return c.sendCmd(cmd)
}

// RawTxInput models the data needed for raw transaction input that is used in
// the SignRawTransactionCmd struct.
type RawTxInput struct {
	Txid         string  `json:"txid"`
	Vout         uint32  `json:"vout"`
	ScriptPubKey string  `json:"scriptPubKey"`
	RedeemScript string  `json:"redeemScript"`
	Amount       float64 `json:"amount"`
}

// SignRawTransactionCmd defines the signrawtransaction JSON-RPC command.
type SignRawTransactionCmd struct {
	RawTx    string
	Inputs   *[]RawTxInput
	PrivKeys *[]string
	Flags    *string `jsonrpcdefault:"\"ALL\""`
}

func NewSignRawTransactionCmd(hexEncodedTx string, inputs *[]RawTxInput, privKeys *[]string, flags *string) *SignRawTransactionCmd {
	return &SignRawTransactionCmd{
		RawTx:    hexEncodedTx,
		Inputs:   inputs,
		PrivKeys: privKeys,
		Flags:    flags,
	}
}

func (c *Client) SignRawTransaction4TWithAmount(tx string,
	inputs []RawTxInput, privKeysWIF []string,
	hashType SigHashType) (string, bool, error) {

	return c.SignRawTransaction4AsyncTWithAmount(tx, inputs, privKeysWIF,
		hashType).ReceiveT()
}

func (c *Client) SignRawTransaction4AsyncTWithAmount(tx string,
	inputs []RawTxInput, privKeysWIF []string,
	hashType SigHashType) FutureSignRawTransactionResult {

	txHex := tx

	cmd := NewSignRawTransactionCmd(txHex, &inputs, &privKeysWIF,
		btcjson.String(string(hashType)))
	return c.sendCmd(cmd)
}

func (r FutureGetRawTransactionResult) ReceiveT() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var txHex string
	err = json.Unmarshal(res, &txHex)
	if err != nil {
		return "", err
	}

	return txHex, nil
}

func (c *Client) GetRawTransactionTAsync(hash string) FutureGetRawTransactionResult {

	cmd := btcjson.NewGetRawTransactionCmd(hash, btcjson.Int(0))
	return c.sendCmd(cmd)
}

func (c *Client) GetRawTransactionT(hash string) (string, error) {

	return c.GetRawTransactionTAsync(hash).ReceiveT()
}

func (r FutureSendToAddressResult) ReceiveT() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var txHash string
	err = json.Unmarshal(res, &txHash)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (c *Client) SendToAddressTAsync(address string, amount btcutil.Amount) FutureSendToAddressResult {
	cmd := btcjson.NewSendToAddressCmd(address, amount.ToBTC(), nil, nil)
	return c.sendCmd(cmd)
}

func (c *Client) SendToAddressT(address string, amount btcutil.Amount) (string, error) {
	return c.SendToAddressTAsync(address, amount).ReceiveT()
}

func (c *Client) GetBalanceTAsync() FutureGetBalanceResult {
	cmd := btcjson.NewGetBalanceCmd(nil, nil)
	return c.sendCmd(cmd)
}

func (c *Client) GetBalanceT() (btcutil.Amount, error) {
	return c.GetBalanceTAsync().Receive()
}


func (c *Client) SendFromTAsync(fromAccount string, toAddress string, amount btcutil.Amount) FutureSendFromResult {
	
	cmd := btcjson.NewSendFromCmd(fromAccount, toAddress, amount.ToBTC(), nil,
		nil, nil)
	return c.sendCmd(cmd)
}

func (c *Client) SendFromT(fromAccount string, toAddress string, amount btcutil.Amount) (string, error) {
	return c.SendFromTAsync(fromAccount, toAddress, amount).ReceiveT()
}

func (r FutureSendFromResult) ReceiveT() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var txHash string
	err = json.Unmarshal(res, &txHash)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func init() {
	flags := btcjson.UFWalletOnly
	btcjson.MustRegisterCmd("signrawtransaction:a", (*SignRawTransactionCmd)(nil), flags)
}
