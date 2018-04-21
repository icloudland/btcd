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

// OmniCreatePayloadSimpleSendCmd defines the omni_createpayload_simplesend JSON-RPC command.

// Create the payload for a simple send transaction.
type OmniCreatePayloadSimpleSendCmd struct {
	PropertyId int
	Amount     string
}

// NewOmniCreatePayloadSimpleSendCmd returns a new instance which can be used to issue a
// omni_createpayload_simplesend JSON-RPC command.
func NewOmniCreatePayloadSimpleSendCmd(propertyId int, amount string) *OmniCreatePayloadSimpleSendCmd {
	return &OmniCreatePayloadSimpleSendCmd{
		PropertyId: propertyId,
		Amount:     amount,
	}
}

// OmniCreateRawTxOpReturnCmd defines the omni_createrawtx_opreturn JSON-RPC command.

// Adds a payload with class C (op-return) encoding to the transaction.
type OmniCreateRawTxOpReturnCmd struct {
	RawTx   string
	Payload string
}

// NewOmniCreateRawTxOpReturn returns a new instance which can be used to issue a
// omni_createrawtx_opreturn JSON-RPC command.
func NewOmniCreateRawTxOpReturn(rawTx string, payload string) *OmniCreateRawTxOpReturnCmd {
	return &OmniCreateRawTxOpReturnCmd{
		RawTx:   rawTx,
		Payload: payload,
	}
}

// OmniCreateRawTxReferenceCmd defines the omni_createrawtx_reference JSON-RPC command.

// Adds a reference output to the transaction
type OmniCreateRawTxReferenceCmd struct {
	RawTx       string
	Destination string
	Amount      *btcutil.Amount
}

// NewOmniCreateRawTxReferenceCmd returns a new instance which can be used to issue a
// omni_createrawtx_reference JSON-RPC command.
func NewOmniCreateRawTxReferenceCmd(rawTx string, destination string, amount *btcutil.Amount) *OmniCreateRawTxReferenceCmd {
	return &OmniCreateRawTxReferenceCmd{
		RawTx:       rawTx,
		Destination: destination,
		Amount:      amount,
	}
}

type PreTx struct {
	Txid         string  `json:"txid"`
	Vout         uint32  `json:"vout"`
	ScriptPubKey string  `json:"scriptPubKey"`
	Value        float64 `json:"value"`
}

// OmniCreateRawTxChangeCmd defines the omni_createrawtx_change JSON-RPC command.

// Adds a reference output to the transaction
type OmniCreateRawTxChangeCmd struct {
	RawTx       string
	PreTxs      []*PreTx
	Destination string
	Fee         float64
}

// NewOmniCreateRawTxChangeCmd returns a new instance which can be used to issue a
// omni_createrawtx_reference JSON-RPC command.
func NewOmniCreateRawTxChangeCmd(rawTx string, preTxs []*PreTx, destination string, fee float64) *OmniCreateRawTxChangeCmd {
	return &OmniCreateRawTxChangeCmd{
		RawTx:       rawTx,
		PreTxs:      preTxs,
		Destination: destination,
		Fee:         fee,
	}
}

// OmniGetBalanceCmd defines the omni_getbalance JSON-RPC command.

// Adds a reference output to the transaction
type OmniGetBalanceCmd struct {
	Address    string
	PropertyId int
}

// NewOmniGetBalanceCmd returns a new instance which can be used to issue a
// omni_getbalance JSON-RPC command.
func NewOmniGetBalanceCmd(address string, propertyId int) *OmniGetBalanceCmd {
	return &OmniGetBalanceCmd{
		Address:    address,
		PropertyId: propertyId,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("omni_send", (*OmniSendCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_getseedblocks", (*OmniGetSeedBlocksCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_listblocktransactions", (*OmniListBlockTransactionsCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_gettransaction", (*OmniGetTransactionsCmd)(nil), flags)

	btcjson.MustRegisterCmd("omni_createpayload_simplesend", (*OmniCreatePayloadSimpleSendCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_createrawtx_opreturn", (*OmniCreateRawTxOpReturnCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_createrawtx_reference", (*OmniCreateRawTxReferenceCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_createrawtx_change", (*OmniCreateRawTxChangeCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_getbalance", (*OmniGetBalanceCmd)(nil), flags)

}
