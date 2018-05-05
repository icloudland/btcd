package btsjson

import "github.com/icloudland/btcdx/btcjson"

// GetRelativeAccountHistoryCmd defines the get_relative_account_history JSON-RPC command.

// Returns the relative operations on the named account from start number.
type GetRelativeAccountHistoryCmd struct {
	Account string
	Start   int
	Limit   int
	End     int
}

// NewGetRelativeAccountHistoryCmd returns a new instance which can be used to issue a
// get_relative_account_history JSON-RPC command.
func NewGetRelativeAccountHistoryCmd(account string, start int, limit int, end int) *GetRelativeAccountHistoryCmd {
	return &GetRelativeAccountHistoryCmd{
		Account: account,
		Start:   start,
		Limit:   limit,
		End:     end,
	}
}

// GetAccountHistoryByOperationsCmd defines the get_account_history_by_operations JSON-RPC command.

// Returns the relative operations on the named account from start number.
type GetAccountHistoryByOperationsCmd struct {
	Account    string
	Operations []int
	Start      int
	Limit      int
}

// NewGetAccountHistoryByOperationsCmd returns a new instance which can be used to issue a
// get_account_history_by_operations JSON-RPC command.
func NewGetAccountHistoryByOperationsCmd(account string, operations []int, start int, limit int) *GetAccountHistoryByOperationsCmd {
	return &GetAccountHistoryByOperationsCmd{
		Account:    account,
		Operations: operations,
		Start:      start,
		Limit:      limit,
	}
}

// Transfer2Cmd defines the transfer2 JSON-RPC command.
type Transfer2Cmd struct {
	From   string
	To     string
	Amount string
	Symbol string
	Memo   string
}

// NewTransfer2Cmd returns a new instance which can be used to issue a
// transfer2 JSON-RPC command.
func NewTransfer2Cmd(from string, to string, amount string, symbol string, memo string) *Transfer2Cmd {
	return &Transfer2Cmd{
		From:   from,
		To:     to,
		Amount: amount,
		Symbol: symbol,
		Memo:   memo,
	}
}

// WalletLockCmd defines the lock JSON-RPC command.
type WalletLockCmd struct{}

// NewWalletLockCmd returns a new instance which can be used to issue a
// lock JSON-RPC command.
func NewWalletLockCmd() *WalletLockCmd {
	return &WalletLockCmd{}
}

// UnWalletLockCmd defines the unlock JSON-RPC command.
type UnWalletLockCmd struct {
	Password string
}

// NewWalletUnLockCmd returns a new instance which can be used to issue a
// unlock JSON-RPC command.
func NewWalletUnLockCmd(password string) *UnWalletLockCmd {
	return &UnWalletLockCmd{
		Password: password,
	}
}


// ListAccountBalancesCmd defines the unlock JSON-RPC command.
type ListAccountBalancesCmd struct {
	Account string
}

// NewListAccountBalancesCmd returns a new instance which can be used to issue a
// unlock JSON-RPC command.
func NewListAccountBalancesCmd(account string) *ListAccountBalancesCmd {
	return &ListAccountBalancesCmd{
		Account: account,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("get_relative_account_history", (*GetRelativeAccountHistoryCmd)(nil), flags)
	btcjson.MustRegisterCmd("get_account_history_by_operations", (*GetAccountHistoryByOperationsCmd)(nil), flags)
	btcjson.MustRegisterCmd("transfer2", (*Transfer2Cmd)(nil), flags)
	btcjson.MustRegisterCmd("lock", (*WalletLockCmd)(nil), flags)
	btcjson.MustRegisterCmd("unlock", (*UnWalletLockCmd)(nil), flags)

	btcjson.MustRegisterCmd("get_account", (*GetBtsAccountCmd)(nil), flags)
	btcjson.MustRegisterCmd("list_account_balances", (*ListAccountBalancesCmd)(nil), flags)

}


// GetBtcAccountCmd defines the getaccount JSON-RPC command.
type GetBtsAccountCmd struct {
	Account string
}

// NewGetAccountCmd returns a new instance which can be used to issue a
// getaccount JSON-RPC command.
func NewGetBtsAccountCmd(account string) *GetBtsAccountCmd {
	return &GetBtsAccountCmd{
		Account: account,
	}
}
