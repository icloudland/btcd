package omnijson

import (
	"github.com/icloudland/btcutil"
	"github.com/icloudland/btcdx/btcjson"
)

// OmniSendCmd defines the omni_send JSON-RPC command.
type OmniSendCmd struct {
	FromAddress string
	ToAddress   string
	PropertyId  int
	Amount      btcutil.Amount
}

// NewOmniSendCmd returns a new instance which can be used to issue a
// omni_send JSON-RPC command.
func NewOmniSendCmd(fromAddress string, toAddress string, propertyId int,
	amount btcutil.Amount) *OmniSendCmd {

	return &OmniSendCmd{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		PropertyId:  propertyId,
		Amount:      amount,
	}
}

// OmniGetSeedBlocksCmd defines the omni_getseedblocks JSON-RPC command.
type OmniGetSeedBlocksCmd struct {
	StartBlock int64
	EndBlock   int64
}

// NewOmniGetSeedBlocksCmd returns a new instance which can be used to issue a
// omni_getseedblocks JSON-RPC command.
func NewOmniGetSeedBlocksCmd(startBlock int64, endBlock int64) *OmniGetSeedBlocksCmd {
	return &OmniGetSeedBlocksCmd{
		StartBlock: startBlock,
		EndBlock:   endBlock,
	}
}

// OmniListBlockTransactionsCmd defines the omni_listblocktransactions JSON-RPC command.
type OmniListBlockTransactionsCmd struct {
	Index int64
}

// NewOmniGetSeedBlocksCmd returns a new instance which can be used to issue a
// omni_listblocktransactions JSON-RPC command.
func NewOmniListBlockTransactionsCmd(index int64) *OmniListBlockTransactionsCmd {
	return &OmniListBlockTransactionsCmd{
		Index: index,
	}
}

// OmniGetTransactionsCmd defines the omni_gettransaction JSON-RPC command.

// Get detailed information about an Omni transaction
type OmniGetTransactionsCmd struct {
	TxId string
}

// NewOmniGetTransactionsCmd returns a new instance which can be used to issue a
// omni_gettransaction JSON-RPC command.
func NewOmniGetTransactionsCmd(txId string) *OmniGetTransactionsCmd {
	return &OmniGetTransactionsCmd{
		TxId: txId,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("omni_send", (*OmniSendCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_getseedblocks", (*OmniGetSeedBlocksCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_listblocktransactions", (*OmniListBlockTransactionsCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_gettransaction", (*OmniGetTransactionsCmd)(nil), flags)
}
