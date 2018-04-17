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
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return "", err
	}

	return txHashStr, nil

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
