package rpcclient

import (
	"encoding/json"
	"github.com/icloudland/btcdx/actjson"
	"fmt"
)

// FutureBlockChainGetBlockCount is a future promise to deliver the result
// of a BlockChainGetBlockCountAsync RPC invocation (or an applicable error).
type FutureBlockChainGetBlockCount chan *response

func (r FutureBlockChainGetBlockCount) Receive() (int32, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal the result as an int64.
	var count int32
	err = json.Unmarshal(res, &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// BlockChainGetBlockCountAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainGetBlockCountAsync() FutureBlockChainGetBlockCount {

	cmd := actjson.NewBlockChainGetBlockCountCmd()
	return c.sendCmd(cmd)
}

// BlockChainGetBlockCount The current block number of blocks
func (c *Client) BlockChainGetBlockCount() (int32, error) {
	return c.BlockChainGetBlockCountAsync().Receive()
}

// FutureBlockchainGetBlockt is a future promise to deliver the result
// of a BlockchainGetBlock Async RPC invocation (or an applicable error).
type FutureBlockChainGetBlock chan *response

func (r FutureBlockChainGetBlock) Receive() (*actjson.BlockChainGetBlockResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an BlockChainGetBlockResult.
	var block actjson.BlockChainGetBlockResult
	err = json.Unmarshal(res, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// BlockChainGetBlockAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainGetBlockAsync(blockNum string) FutureBlockChainGetBlock {

	cmd := actjson.NewBlockChainGetBlockCmd(blockNum)
	return c.sendCmd(cmd)
}

// BlockChainGetBlock To specify the block by block number or ID, and obtain the header information of its block
func (c *Client) BlockChainGetBlock(blockNum string) (*actjson.BlockChainGetBlockResult, error) {
	return c.BlockChainGetBlockAsync(blockNum).Receive()
}

// FutureBlockChainGetTransaction is a future promise to deliver the result
// of a BlockChainGetTransactionAsync RPC invocation (or an applicable error).
type FutureBlockChainGetTransaction chan *response

func (r FutureBlockChainGetTransaction) Receive() (*actjson.BlockChainGetTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an BlockChainGetTransactionResult.
	var transaction actjson.BlockChainGetTransactionResult
	err = json.Unmarshal(res, &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// BlockChainGetBlockAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainGetTransactionAsync(txId string) FutureBlockChainGetTransaction {

	cmd := actjson.NewBlockChainGetTransactionCmd(txId)
	return c.sendCmd(cmd)
}

// BlockChainGetBlock To specify the block by block number or ID, and obtain the header information of its block
func (c *Client) BlockChainGetTransaction(txId string) (*actjson.BlockChainGetTransactionResult, error) {
	return c.BlockChainGetTransactionAsync(txId).Receive()
}

// FutureBlockChainGetPrettyTransaction is a future promise to deliver the result
// of a BlockChainGetPrettyTransactionAsync RPC invocation (or an applicable error).
type FutureBlockChainGetPrettyTransaction chan *response

func (r FutureBlockChainGetPrettyTransaction) Receive() (*actjson.BlockChainGetPrettyTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an BlockChainGetPrettyTransactionResult.
	var transaction actjson.BlockChainGetPrettyTransactionResult
	err = json.Unmarshal(res, &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// BlockChainGetPrettyTransactionAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainGetPrettyTransactionAsync(txId string) FutureBlockChainGetPrettyTransaction {

	cmd := actjson.NewBlockChainGetPrettyTransactionCmd(txId)
	return c.sendCmd(cmd)
}

// BlockChainGetPrettyTransaction
func (c *Client) BlockChainGetPrettyTransaction(txId string) (*actjson.BlockChainGetPrettyTransactionResult, error) {
	return c.BlockChainGetPrettyTransactionAsync(txId).Receive()
}

// FutureBlockChainGetPrettyContractTransaction is a future promise to deliver the result
// of a BlockChainGetPrettyContractTransactionAsync RPC invocation (or an applicable error).
type FutureBlockChainGetPrettyContractTransaction chan *response

func (r FutureBlockChainGetPrettyContractTransaction) Receive() (*actjson.BlockChainGetPrettyContractTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an BlockChainGetPrettyContractTransactionResult.
	var transaction actjson.BlockChainGetPrettyContractTransactionResult
	err = json.Unmarshal(res, &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// BlockChainGetPrettyTransactionAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainGetPrettyContractTransactionAsync(txId string) FutureBlockChainGetPrettyContractTransaction {

	cmd := actjson.NewBlockChainGetPrettyContractTransactionCmd(txId)
	return c.sendCmd(cmd)
}

// BlockChainGetPrettyTransaction
func (c *Client) BlockChainGetPrettyContractTransaction(txId string) (*actjson.BlockChainGetPrettyContractTransactionResult, error) {
	return c.BlockChainGetPrettyContractTransactionAsync(txId).Receive()
}

// FutureBlockChainGetEvents is a future promise to deliver the result
// of a BlockChainGetEventsAsync RPC invocation (or an applicable error).
type FutureBlockChainGetEvents chan *response

func (r FutureBlockChainGetEvents) Receive() ([]actjson.BlockChainGetEvent, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an BlockChainGetTransactionResult.
	var getEvents []actjson.BlockChainGetEvent
	err = json.Unmarshal(res, &getEvents)
	if err != nil {
		return nil, err
	}
	return getEvents, nil
}

// BlockChainGetBlockAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainGetEventsAsync(blockNum int32, txId string) FutureBlockChainGetEvents {

	cmd := actjson.NewBlockChainGetEventsCmd(blockNum, txId)
	return c.sendCmd(cmd)
}

// BlockChainGetBlock To specify the block by block number or ID, and obtain the header information of its block
func (c *Client) BlockChainGetEvents(blockNum int32, txId string) ([]actjson.BlockChainGetEvent, error) {
	return c.BlockChainGetEventsAsync(blockNum, txId).Receive()
}
// FutureBlockChainGetEvents is a future promise to deliver the result
// of a BlockChainGetEventsAsync RPC invocation (or an applicable error).
type FutureBlockChainListAddressBalances chan *response

func (r FutureBlockChainListAddressBalances) Receive() ([]actjson.BlockChainListAddressBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an BlockChainGetTransactionResult.
	var results []actjson.BlockChainListAddressBalanceResult
	err = json.Unmarshal(res, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// BlockChainGetBlockAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainListAddressBalancesAsync(addr string) FutureBlockChainListAddressBalances {

	cmd := actjson.NewBlockChainListAddressBalancesCmd(addr)
	return c.sendCmd(cmd)
}

// BlockChainGetBlock To specify the block by block number or ID, and obtain the header information of its block
func (c *Client) BlockChainListAddressBalances(addr string) ([]actjson.BlockChainListAddressBalanceResult, error) {
	return c.BlockChainListAddressBalancesAsync(addr).Receive()
}

// FutureBlockChainGetEvents is a future promise to deliver the result
// of a BlockChainGetEventsAsync RPC invocation (or an applicable error).
type FutureCallContract chan *response

func (r FutureCallContract) Receive() (*actjson.CallContractResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an CallContractResult.
	var callResult actjson.CallContractResult
	err = json.Unmarshal(res, &callResult)
	if err != nil {
		return nil, err
	}
	return &callResult, nil
}

// BlockChainGetBlockAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) CallContractAsync(contract string, callerName string, functionName string,
	params string, assetSymbol string, callLimit float64) FutureCallContract {
	cmd := actjson.NewCallContractCmd(contract, callerName, functionName, params, assetSymbol, callLimit)
	return c.sendCmd(cmd)
}

// BlockChainGetBlock To specify the block by block number or ID, and obtain the header information of its block
func (c *Client) CallContract(contract string, callerName string, functionName string,
	params string, assetSymbol string, callLimit float64) (*actjson.CallContractResult, error) {
	return c.CallContractAsync(contract, callerName, functionName, params, assetSymbol, callLimit).Receive()
}
// FutureWalletTransferToAddress is a future promise to deliver the result
// of a WalletTransferToAddressAsync RPC invocation (or an applicable error).
type FutureWalletTransferToAddress chan *response

func (r FutureWalletTransferToAddress) Receive() (error) {
	res, err := receiveFuture(r)
	if err != nil {
		return err
	}

	fmt.Println(string(res))
	// Unmarshal the result as an CallContractResult.
	//var callResult actjson.CallContractResult
	//err = json.Unmarshal(res, &callResult)
	//if err != nil {
	//	return nil, err
	//}
	//return &callResult, nil

	return nil
}

// BlockChainGetBlockAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) WalletTransferToAddressAsync(amount string, assetSymbol string, fromAccountName string,
	toAddress string) FutureWalletTransferToAddress {
	cmd := actjson.NewWalletTransferToAddressCmd(amount, assetSymbol, fromAccountName, toAddress)
	return c.sendCmd(cmd)
}

// BlockChainGetBlock To specify the block by block number or ID, and obtain the header information of its block
func (c *Client) WalletTransferToAddress(amount string, assetSymbol string, fromAccountName string,
	toAddress string) ( error) {
	return c.WalletTransferToAddressAsync(amount, assetSymbol, fromAccountName, toAddress).Receive()
}

// FutureBlockChainGetEvents is a future promise to deliver the result
// of a BlockChainGetEventsAsync RPC invocation (or an applicable error).
type FutureBlockChainGetContractResult chan *response

func (r FutureBlockChainGetContractResult) Receive() (*actjson.BlockChainGetContractResultResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an CallContractResult.
	var callResult actjson.BlockChainGetContractResultResult
	err = json.Unmarshal(res, &callResult)
	if err != nil {
		return nil, err
	}
	return &callResult, nil
}

// BlockChainGetBlockAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) BlockChainGetContractResultAsync(resultId string) FutureBlockChainGetContractResult {
	cmd := actjson.NewBlockChainGetContractResultCmd(resultId)
	return c.sendCmd(cmd)
}

// BlockChainGetContractResult
func (c *Client) BlockChainGetContractResult(resultId string) (*actjson.BlockChainGetContractResultResult, error) {
	return c.BlockChainGetContractResultAsync(resultId).Receive()
}


// FutureWalletLock is a future promise to deliver the result
// of a WalletLockBtsAsync RPC invocation (or an applicable error).
type FutureWalletLockAct chan *response

// Receive waits for the response promised by the future and returns a hash string
func (r FutureWalletLockAct) Receive() (error) {
	_, err := receiveFuture(r)
	return err
}

// WalletLockActAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) WalletLockActAsync() FutureWalletLockAct {

	cmd := actjson.NewWalletLockCmd()
	return c.sendCmd(cmd)
}

// WalletLockAct lock wallet
func (c *Client) WalletLockAct() (error) {
	return c.WalletLockActAsync().Receive()
}

// FutureWalletUnLockAct is a future promise to deliver the result
// of a WalletUnLockActAsync RPC invocation (or an applicable error).
type FutureWalletUnLockAct chan *response

// Receive waits for the response promised by the future and returns a hash string
func (r FutureWalletUnLockAct) Receive() (error) {
	_, err := receiveFuture(r)
	return err
}

// WalletUnLockActAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func (c *Client) WalletUnLockActAsync(timeout string, password string) FutureWalletUnLockAct {

	cmd := actjson.NewWalletUnLockCmd(timeout, password)
	return c.sendCmd(cmd)
}

// WalletUnLockAct lock act wallet for timeout, password
func (c *Client) WalletUnLockAct(timeout string, password string) (error) {
	return c.WalletUnLockActAsync(timeout, password).Receive()
}
