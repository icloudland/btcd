package rpcclient

import (
	"github.com/icloudland/btcdx/btsjson"
	"encoding/json"
)

// FutureWalletLock is a future promise to deliver the result
// of a WalletLockBtsAsync RPC invocation (or an applicable error).
type FutureWalletLock chan *response

// Receive waits for the response promised by the future and returns a hash string
func (r FutureWalletLock) Receive() (error) {
	_, err := receiveFuture(r)
	return err
}

// WalletLockBtsAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) WalletLockBtsAsync() FutureWalletLock {

	cmd := btsjson.NewWalletLockCmd()
	return c.sendCmd(cmd)
}

// WalletLockBts lock wallet
func (c *Client) WalletLockBts() (error) {
	return c.WalletLockBtsAsync().Receive()
}

// FutureWalletUnLock is a future promise to deliver the result
// of a WalletUnLockBtsAsync RPC invocation (or an applicable error).
type FutureWalletUnLock chan *response

// Receive waits for the response promised by the future and returns a hash string
func (r FutureWalletUnLock) Receive() (error) {
	_, err := receiveFuture(r)
	return err
}

// WalletUnLockBtsAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) WalletUnLockBtsAsync(pwd string) FutureWalletUnLock {

	cmd := btsjson.NewWalletUnLockCmd(pwd)
	return c.sendCmd(cmd)
}

// WalletUnLockBts lock bts wallet for a given pwd
func (c *Client) WalletUnLockBts(pwd string) (error) {
	return c.WalletUnLockBtsAsync(pwd).Receive()
}

// FuturGetRelativeAccountHistory is a future promise to deliver the result
// of a GetRelativeAccountHistoryAsync RPC invocation (or an applicable error).
type FutureGetRelativeAccountHistory chan *response

// Receive waits for the response promised by the future and return a list of operation detail
func (r FutureGetRelativeAccountHistory) Receive() ([]btsjson.OperationDetail, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as an array of getRelativeAccountHistory result objects.
	var operationDetails []btsjson.OperationDetail
	err = json.Unmarshal(res, &operationDetails)
	if err != nil {
		return nil, err
	}

	return operationDetails, nil
}

// GetRelativeAccountHistoryAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) GetRelativeAccountHistoryAsync(account string, start int, limit int, end int) FutureGetRelativeAccountHistory {

	cmd := btsjson.NewGetRelativeAccountHistoryCmd(account, start, limit, end)
	return c.sendCmd(cmd)
}

// GetRelativeAccountHistory Returns a list of operation detail
func (c *Client) GetRelativeAccountHistory(account string, start int, limit int, end int) ([]btsjson.OperationDetail, error) {
	return c.GetRelativeAccountHistoryAsync(account, start, limit, end).Receive()
}

// FutureTransfer2 is a future promise to deliver the result
// of a Transfer2Async RPC invocation (or an applicable error).
type FutureTransfer2 chan *response

func (r FutureTransfer2) Receive() (*btsjson.Transfer2Result, error) {

	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as Transfer2Result result objects.
	var transfer2Result btsjson.Transfer2Result
	err = json.Unmarshal(res, &transfer2Result)
	if err != nil {
		return nil, err
	}

	return &transfer2Result, nil
}

// Transfer2Async returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) Transfer2Async(from string, to string, amount string, symbol string, memo string) FutureTransfer2 {
	cmd := btsjson.NewTransfer2Cmd(from, to, amount, symbol, memo)
	return c.sendCmd(cmd)
}

// Transfer2 create a bts transaction
func (c *Client) Transfer2(from string, to string, amount string, symbol string, memo string) (*btsjson.Transfer2Result, error) {
	return c.Transfer2Async(from, to, amount, symbol, memo).Receive()
}

// FutureGetAccountHistoryByOperations is a future promise to deliver the result
// of a GetAccountHistoryByOperationsAsync RPC invocation (or an applicable error).
type FutureGetAccountHistoryByOperations chan *response

// Receive waits for the response promised by the future and return a list of operation detail
func (r FutureGetAccountHistoryByOperations) Receive() (*btsjson.AccountHistoryOperationDetail, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as an array of getRelativeAccountHistory result objects.
	var operationDetails btsjson.AccountHistoryOperationDetail
	err = json.Unmarshal(res, &operationDetails)
	if err != nil {
		return nil, err
	}

	return &operationDetails, nil
}

// GetAccountHistoryByOperationsAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) GetAccountHistoryByOperationsAsync(account string, operations []int, start int, limit int) FutureGetAccountHistoryByOperations {

	cmd := btsjson.NewGetAccountHistoryByOperationsCmd(account, operations, start, limit)
	return c.sendCmd(cmd)
}

// GetAccountHistoryByOperations Returns a list of operation detail
func (c *Client) GetAccountHistoryByOperations(account string, operations []int, start int, limit int) (*btsjson.AccountHistoryOperationDetail, error) {
	return c.GetAccountHistoryByOperationsAsync(account, operations, start, limit).Receive()
}

// FutureGetBtsAccountResult is a future promise to deliver the result of a
// GetBtsAccountAsync RPC invocation (or an applicable error).
type FutureGetBtsAccountResult chan *response

// Receive waits for the response promised by the future and returns the account
func (r FutureGetBtsAccountResult) Receive() (error) {
	_, err := receiveFuture(r)
	return err

}

// GetBtsAccountAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
func (c *Client) GetBtsAccountAsync(account string) FutureGetBtsAccountResult {
	cmd := btsjson.NewGetBtsAccountCmd(account)
	return c.sendCmd(cmd)
}

// GetBtsAccount
func (c *Client) GetBtsAccount(account string) (error) {
	return c.GetBtsAccountAsync(account).Receive()
}

// FutureListAccountBalances is a future promise to deliver the result of a
// ListAccountBalancesAsync RPC invocation (or an applicable error).
type FutureListAccountBalances chan *response

// Receive waits for the response promised by the future and returns the account
func (r FutureListAccountBalances) Receive() ([]btsjson.AssetAmount, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as an array of getRelativeAccountHistory result objects.
	var balances []btsjson.AssetAmount
	err = json.Unmarshal(res, &balances)
	if err != nil {
		return nil, err
	}

	return balances, nil

}

// ListAccountBalancesAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
func (c *Client) ListAccountBalancesAsync(account string) FutureListAccountBalances {
	cmd := btsjson.NewListAccountBalancesCmd(account)
	return c.sendCmd(cmd)
}

// ListAccountBalances, return a list of the given account's balances
func (c *Client) ListAccountBalances(account string) ([]btsjson.AssetAmount, error) {
	return c.ListAccountBalancesAsync(account).Receive()
}
