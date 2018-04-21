package rpcclient

import (
	"github.com/icloudland/btcutil"
	"github.com/icloudland/btcdx/omnijson"
	"encoding/json"
)

// FutureOmniSendResult is a future promise to deliver the result
// of a OmniSendAsync RPC invocation (or an applicable error).
type FutureOmniSendResult chan *response

func (c *Client) OmniSendAsync(fromAddress string, toAddress string,
	propertyId int, amount btcutil.Amount) FutureOmniSendResult {

	cmd := omnijson.NewOmniSendCmd(fromAddress, toAddress, propertyId, amount)
	return c.sendCmd(cmd)
}

func (c *Client) OmniSend(fromAddress string, toAddress string,
	propertyId int, amount btcutil.Amount) (string, error) {

	return c.OmniSendAsync(fromAddress, toAddress, propertyId, amount).Receive()
}

// Receive waits for the response promised by the future and returns a new
// transaction hash
func (r FutureOmniSendResult) Receive() (string, error) {
	return handleStringResult(r)

}

// FutureOmniListBlockTransactionsResult is a future promise to deliver the result
// of a OmniListBlockTransactionsAsync RPC invocation (or an applicable error).
type FutureOmniListBlockTransactionsResult chan *response

// Receive waits for the response promised by the future and returns a list of transaction hash
func (r FutureOmniListBlockTransactionsResult) Receive() ([]string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var txs []string
	err = json.Unmarshal(res, &txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}

// OmniListBlockTransactionsAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See OmniListBlockTransactions for the blocking version and more details.
func (c *Client) OmniListBlockTransactionsAsync(index int64) FutureOmniListBlockTransactionsResult {

	cmd := omnijson.NewOmniListBlockTransactionsCmd(index)
	return c.sendCmd(cmd)
}

// OmniListBlockTransactions Lists all Omni transactions in a block
func (c *Client) OmniListBlockTransactions(index int64) ([]string, error) {

	return c.OmniListBlockTransactionsAsync(index).Receive()
}

// OmniGetTransactionResult is a future promise to deliver the result
// of a OmniGetTransactionAsync RPC invocation (or an applicable error).
type FutureOmniGetTransactionResult chan *response

// Receive waits for the response promised by the future and returns a new
// transaction spending the provided inputs and sending to the provided
// addresses.
func (r FutureOmniGetTransactionResult) Receive() (*omnijson.OmniGetTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result omnijson.OmniGetTransactionResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}

// OmniGetTransactionAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See OmniGetTransaction for the blocking version and more details.
func (c *Client) OmniGetTransactionAsync(txId string) FutureOmniGetTransactionResult {

	cmd := omnijson.NewOmniGetTransactionsCmd(txId)
	return c.sendCmd(cmd)
}

// OmniGetTransaction Get detailed information about an Omni transaction
func (c *Client) OmniGetTransaction(txId string) (*omnijson.OmniGetTransactionResult, error) {

	return c.OmniGetTransactionAsync(txId).Receive()
}

// FutureOmniCreatePayloadSimpleSend is a future promise to deliver the result
// of a OmniCreatePayloadSimpleSendAsync RPC invocation (or an applicable error).
type FutureOmniCreatePayloadSimpleSend chan *response

// Receive waits for the response promised by the future and returns a payload string
func (r FutureOmniCreatePayloadSimpleSend) Receive() (string, error) {
	return handleStringResult(r)

}

// OmniCreatePayloadSimpleSendAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See OmniGetTransaction for the blocking version and more details.
func (c *Client) OmniCreatePayloadSimpleSendAsync(propertyId int, amount string) FutureOmniCreatePayloadSimpleSend {

	cmd := omnijson.NewOmniCreatePayloadSimpleSendCmd(propertyId, amount)
	return c.sendCmd(cmd)
}

// OmniCreatePayloadSimpleSend Create the payload for a simple send transaction
func (c *Client) OmniCreatePayloadSimpleSend(propertyId int, amount string) (string, error) {

	return c.OmniCreatePayloadSimpleSendAsync(propertyId, amount).Receive()
}

// FutureOmniCreatePayloadSimpleSend is a future promise to deliver the result
// of a OmniCreatePayloadSimpleSendAsync RPC invocation (or an applicable error).
type FutureOmniCreateRawTxOpReturn chan *response

// Receive waits for the response promised by the future and returns a payload string
func (r FutureOmniCreateRawTxOpReturn) Receive() (string, error) {
	return handleStringResult(r)

}

// OmniCreatePayloadSimpleSendAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See OmniGetTransaction for the blocking version and more details.
func (c *Client) OmniCreateRawTxOpReturnAsync(rawTx string, payload string) FutureOmniCreateRawTxOpReturn {

	cmd := omnijson.NewOmniCreateRawTxOpReturn(rawTx, payload)
	return c.sendCmd(cmd)
}

// OmniCreatePayloadSimpleSend Create the payload for a simple send transaction
func (c *Client) OmniCreateRawTxOpReturn(rawTx string, payload string) (string, error) {

	return c.OmniCreateRawTxOpReturnAsync(rawTx, payload).Receive()
}

// FuturOmniCreateRawTxReference is a future promise to deliver the result
// of a OmniCreateRawTxReferenceAsync RPC invocation (or an applicable error).
type FuturOmniCreateRawTxReference chan *response

// Receive waits for the response promised by the future and returns a hash string
func (r FuturOmniCreateRawTxReference) Receive() (string, error) {
	return handleStringResult(r)

}

// OmniCreateRawTxReferenceAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// Adds a reference output to the transaction.
func (c *Client) OmniCreateRawTxReferenceAsync(rawTx string, destination string, amount *btcutil.Amount) FuturOmniCreateRawTxReference {

	cmd := omnijson.NewOmniCreateRawTxReferenceCmd(rawTx, destination, amount)
	return c.sendCmd(cmd)
}

// OmniCreateRawTxReference Adds a reference output to the transaction.
func (c *Client) OmniCreateRawTxReference(rawTx string, destination string, amount *btcutil.Amount) (string, error) {
	return c.OmniCreateRawTxReferenceAsync(rawTx, destination, amount).Receive()
}

// FuturOmniCreateRawTxReference is a future promise to deliver the result
// of a OmniCreateRawTxReferenceAsync RPC invocation (or an applicable error).
type FuturOmniCreateRawTxChange chan *response

// Receive waits for the response promised by the future and returns a hash string
func (r FuturOmniCreateRawTxChange) Receive() (string, error) {
	return handleStringResult(r)
}

// OmniCreateRawTxReferenceAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// Adds a reference output to the transaction.
func (c *Client) OmniCreateRawTxChangeAsync(rawTx string, prevTxs []*omnijson.PreTx, destination string, fee float64) FuturOmniCreateRawTxChange {

	cmd := omnijson.NewOmniCreateRawTxChangeCmd(rawTx, prevTxs, destination, fee)
	return c.sendCmd(cmd)
}

// OmniCreateRawTxChange Adds a change output to the transaction
func (c *Client) OmniCreateRawTxChange(rawTx string, prevTxs []*omnijson.PreTx, destination string, fee float64) (string, error) {
	return c.OmniCreateRawTxChangeAsync(rawTx, prevTxs, destination, fee).Receive()
}

// FuturOmniGetBalance is a future promise to deliver the result
// of a OmniGetBalanceAsync RPC invocation (or an applicable error).
type FuturOmniGetBalance chan *response

// Receive waits for the response promised by the future and returns a hash string
func (r FuturOmniGetBalance) Receive() (string, error) {

	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	type getBalanceResult struct {
		Balance  string
		Reserved string
	}

	var result getBalanceResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result.Balance, nil
}

// OmniGetBalanceAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// Returns the token balance for a given address and property
func (c *Client) OmniGetBalanceAsync(address string, propertyId int) FuturOmniGetBalance {

	cmd := omnijson.NewOmniGetBalanceCmd(address, propertyId)
	return c.sendCmd(cmd)
}

// OmniGetBalance Returns the token balance for a given address and property
func (c *Client) OmniGetBalance(address string, propertyId int) (string, error) {
	return c.OmniGetBalanceAsync(address, propertyId).Receive()
}

func handleStringResult(r chan *response) (string, error) {

	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}
