package actjson

import "github.com/icloudland/btcdx/btcjson"

// BlockChainGetBlockCountCmd defines the blockchain_get_block_count JSON-RPC command.

type BlockChainGetBlockCountCmd struct {
}

func NewBlockChainGetBlockCountCmd() *BlockChainGetBlockCountCmd {
	return &BlockChainGetBlockCountCmd{}
}

// BlockChainGetBlockCmd defines the blockchain_get_block JSON-RPC command.

type BlockChainGetBlockCmd struct {
	BlockNum string
}

func NewBlockChainGetBlockCmd(blockNum string) *BlockChainGetBlockCmd {
	return &BlockChainGetBlockCmd{
		BlockNum: blockNum,
	}
}

// BlockChainGetTransactionCmd defines the blockchain_get_transaction JSON-RPC command.
type BlockChainGetTransactionCmd struct {
	TxId string
}

func NewBlockChainGetTransactionCmd(txId string) *BlockChainGetTransactionCmd {
	return &BlockChainGetTransactionCmd{
		TxId: txId,
	}
}
// BlockChainListAddressBalancesCmd defines the blockchain_list_address_balances JSON-RPC command.
type BlockChainListAddressBalancesCmd struct {
	Addr string
}

func NewBlockChainListAddressBalancesCmd(addr string) *BlockChainListAddressBalancesCmd {
	return &BlockChainListAddressBalancesCmd{
		Addr: addr,
	}
}

// BlockChainGetPrettyTransactionCmd defines the blockchain_get_pretty_transaction JSON-RPC command.
type BlockChainGetPrettyTransactionCmd struct {
	TxId string
}

func NewBlockChainGetPrettyTransactionCmd(txId string) *BlockChainGetPrettyTransactionCmd {
	return &BlockChainGetPrettyTransactionCmd{
		TxId: txId,
	}
}

// BlockChainGetPrettyContractTransactionCmd defines the blockchain_get_pretty_contract_transaction JSON-RPC command.
type BlockChainGetPrettyContractTransactionCmd struct {
	TxId string
}

func NewBlockChainGetPrettyContractTransactionCmd(txId string) *BlockChainGetPrettyContractTransactionCmd {
	return &BlockChainGetPrettyContractTransactionCmd{
		TxId: txId,
	}
}

// BlockChainGetEventsCmd defines the blockchain_get_events JSON-RPC command.
type BlockChainGetEventsCmd struct {
	BlockNum int32
	TxId     string
}

func NewBlockChainGetEventsCmd(blockNum int32, txId string) *BlockChainGetEventsCmd {
	return &BlockChainGetEventsCmd{
		BlockNum: blockNum,
		TxId:     txId,
	}
}

// CallContractCmd defines the call_contract JSON-RPC command.
type CallContractCmd struct {
	Contract     string
	CallerName   string
	FunctionName string
	Params       string
	AssetSymbol  string
	CallLimit    float64
}

func NewCallContractCmd(contract string, callerName string, functionName string,
	params string, assetSymbol string, callLimit float64) *CallContractCmd {
	return &CallContractCmd{
		Contract:     contract,
		CallerName:   callerName,
		FunctionName: functionName,
		Params:       params,
		AssetSymbol:  assetSymbol,
		CallLimit:    callLimit,
	}
}

// BlockChainGetContractResultCmd defines the blockchain_get_contract_result JSON-RPC command.
type BlockChainGetContractResultCmd struct {
	ResultId string
}

func NewBlockChainGetContractResultCmd(resultId string) *BlockChainGetContractResultCmd {
	return &BlockChainGetContractResultCmd{
		ResultId: resultId,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("blockchain_get_block_count", (*BlockChainGetBlockCountCmd)(nil), flags)
	btcjson.MustRegisterCmd("blockchain_get_block", (*BlockChainGetBlockCmd)(nil), flags)
	btcjson.MustRegisterCmd("blockchain_get_transaction", (*BlockChainGetTransactionCmd)(nil), flags)
	btcjson.MustRegisterCmd("blockchain_get_pretty_transaction", (*BlockChainGetPrettyTransactionCmd)(nil), flags)
	btcjson.MustRegisterCmd("blockchain_get_pretty_contract_transaction", (*BlockChainGetPrettyContractTransactionCmd)(nil), flags)
	btcjson.MustRegisterCmd("blockchain_list_address_balances", (*BlockChainListAddressBalancesCmd)(nil), flags)

	btcjson.MustRegisterCmd("blockchain_get_events", (*BlockChainGetEventsCmd)(nil), flags)
	btcjson.MustRegisterCmd("call_contract", (*CallContractCmd)(nil), flags)
	btcjson.MustRegisterCmd("blockchain_get_contract_result", (*BlockChainGetContractResultCmd)(nil), flags)

}
