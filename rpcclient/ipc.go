package rpcclient

import (
	"encoding/json"
	"github.com/icloudland/btcdx/btcjson"
	"github.com/icloudland/btcdx/ipcjson"
)

func (c *Client) ListTokenTransactionsCountFromAsync(token string, count, from int) FutureListTransactionsResult {
	cmd := ipcjson.NewListTokenTransactionsCmd(&token, &count, &from, nil)
	return c.sendCmd(cmd)
}

func (c *Client) ListTokenTransactionsCountFrom(account string, count, from int) ([]btcjson.ListTransactionsResult, error) {
	return c.ListTokenTransactionsCountFromAsync(account, count, from).Receive()
}

type FutureGetTokenBalanceResult chan *response

func (r FutureGetTokenBalanceResult) Receive() (float64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	var balance float64
	err = json.Unmarshal(res, &balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (c *Client) GetTokenBalanceTAsync(coinName string) FutureGetTokenBalanceResult {
	cmd := ipcjson.NewGetTokenBalanceCmd(coinName)
	return c.sendCmd(cmd)
}

func (c *Client) GetTokenBalanceT(coinName string) (float64, error) {
	return c.GetTokenBalanceTAsync(coinName).Receive()
}

type FutureIPCTokenSendToAddressResult chan *response

func (r FutureIPCTokenSendToAddressResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txHash string
	err = json.Unmarshal(res, &txHash)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (c *Client) IPCTokenSendToAddressAsync(tokenSymbol, address string, value float64) FutureIPCTokenSendToAddressResult {
	cmd := ipcjson.NewIPCTokenSendToAddressCmd(tokenSymbol, address, value)
	return c.sendCmd(cmd)
}

func (c *Client) IPCTokenSendToAddress(tokenSymbol, address string, amount float64) (string, error) {
	return c.IPCTokenSendToAddressAsync(tokenSymbol, address, amount).Receive()
}
