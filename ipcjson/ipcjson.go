package ipcjson

import "github.com/icloudland/btcdx/btcjson"

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("listtokentransactions", (*ListTokenTransactionsCmd)(nil), flags)
	btcjson.MustRegisterCmd("gettokenbalance", (*GetTokenBalanceCmd)(nil), flags)
	btcjson.MustRegisterCmd("IPCTokenSendToAddress", (*IPCTokenSendToAddressCmd)(nil), flags)

}

type ListTokenTransactionsCmd struct {
	Token            *string
	Count            *int  `jsonrpcdefault:"10"`
	From             *int  `jsonrpcdefault:"0"`
	IncludeWatchOnly *bool `jsonrpcdefault:"false"`
}

func NewListTokenTransactionsCmd(token *string, count, from *int, includeWatchOnly *bool) *ListTokenTransactionsCmd {
	return &ListTokenTransactionsCmd{
		Token:            token,
		Count:            count,
		From:             from,
		IncludeWatchOnly: includeWatchOnly,
	}
}

type GetTokenBalanceCmd struct {
	TokenSymbol string
}

func NewGetTokenBalanceCmd(account string) *GetTokenBalanceCmd {
	return &GetTokenBalanceCmd{
		TokenSymbol: account,
	}
}

type IPCTokenSendToAddressCmd struct {
	TokenSymbol string
	Address     string
	Amount      float64
}

func NewIPCTokenSendToAddressCmd(tokenSymbol, address string, amount float64) *IPCTokenSendToAddressCmd {
	return &IPCTokenSendToAddressCmd{
		TokenSymbol: tokenSymbol,
		Address:     address,
		Amount:      amount,
	}
}
