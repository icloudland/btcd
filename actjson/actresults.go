package actjson

import (
	"encoding/json"
	"github.com/juju/errors"
)

type BlockChainGetBlockResult struct {
	BlockNum           int32    `json:"block_num"`
	BlockSize          int32    `json:"block_size"`
	Timestamp          string   `json:"timestamp"`
	UserTransactionIds []string `json:"user_transaction_ids"`
}

type BlockChainGetTransactionResult struct {
	TxId   string
	Detail BlockChainGetTransactionResultDetail
}

type BlockChainGetTransactionResultDetail struct {
	Trx TrxContent `json:"trx"`
}

type TrxContent struct {
	AlpAccount string      `json:"alp_account"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Type string `json:"type"`
}

func (p *BlockChainGetTransactionResult) UnmarshalJSON(data []byte) error {

	raw := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "Unmarshal raw object")
	}

	if len(raw) != 2 {
		return errors.Errorf("Invalid operation data: %v", string(data))
	}

	if err := json.Unmarshal(raw[0], &p.TxId); err != nil {
		return errors.Annotate(err, "Unmarshal TxId")
	}

	if err := json.Unmarshal(raw[1], &p.Detail); err != nil {
		return errors.Annotate(err, "Unmarshal BlockChainGetTransactionResultDetail")
	}

	return nil
}

type BlockChainGetPrettyTransactionResult struct {
	BlockNum            int           `json:"block_num"`
	BlockPosition       int           `json:"block_position"`
	ExpirationTimestamp string        `json:"expiration_timestamp"`
	Fee                 AssetAmount   `json:"fee"`
	IsConfirmed         bool          `json:"is_confirmed"`
	IsMarket            bool          `json:"is_market"`
	IsMarketCancel      bool          `json:"is_market_cancel"`
	IsVirtual           bool          `json:"is_virtual"`
	LedgerEntries       []LedgerEntry `json:"ledger_entries"`
	Timestamp           string        `json:"timestamp"`
	TrxID               string        `json:"trx_id"`
	TrxType             int           `json:"trx_type"`
}

type LedgerEntry struct {
	FromAccount     string        `json:"from_account"`
	FromAccountName string        `json:"from_account_name"`
	Memo            string        `json:"memo"`
	RunningBalances []interface{} `json:"running_balances"`
	ToAccount       string        `json:"to_account"`
	ToAccountName   string        `json:"to_account_name"`
	Amount          AssetAmount   `json:"amount"`
}

type AssetAmount struct {
	AssetID int   `json:"asset_id"`
	Amount  int64 `json:"amount"`
}

type BlockChainGetPrettyContractTransactionResult struct {
	BlockNum                  int                   `json:"block_num"`
	BlockPosition             int                   `json:"block_position"`
	ExpirationTimestamp       string                `json:"expiration_timestamp"`
	FromContractLedgerEntries []interface{}         `json:"from_contract_ledger_entries"`
	IsCompleted               bool                  `json:"is_completed"`
	OrigTrxID                 string                `json:"orig_trx_id"`
	Reserved                  []string              `json:"reserved"`
	ResultTrxID               string                `json:"result_trx_id"`
	Timestamp                 string                `json:"timestamp"`
	LedgerEntry               ToContractLedgerEntry `json:"to_contract_ledger_entry"`
	TrxType                   int                   `json:"trx_type"`
}

type ToContractLedgerEntry struct {
	Amount          AssetAmount `json:"amount"`
	Fee             AssetAmount `json:"fee"`
	FromAccount     string      `json:"from_account"`
	FromAccountName string      `json:"from_account_name"`
	Memo            string      `json:"memo"`
	ToAccount       string      `json:"to_account"`
	ToAccountName   string      `json:"to_account_name"`
}

type BlockChainGetEvent struct {
	EventParam  string `json:"event_param"`
	EventType   string `json:"event_type"`
	ID          string `json:"id"`
	IsTruncated bool   `json:"is_truncated"`
}

type CallContractResult struct {
	BlockNum      int           `json:"block_num"`
	CreatedTime   string        `json:"created_time"`
	EntryID       string        `json:"entry_id"`
	Fee           AssetAmount   `json:"fee"`
	Index         int           `json:"index"`
	IsConfirmed   bool          `json:"is_confirmed"`
	IsMarket      bool          `json:"is_market"`
	IsVirtual     bool          `json:"is_virtual"`
	LedgerEntries []interface{} `json:"ledger_entries"`
	ReceivedTime  string        `json:"received_time"`
}

type BlockChainGetContractResultResult struct {
	BlockNum int    `json:"block_num"`
	TrxId    string `json:"trx_id"`
}

type AddressBalance struct {
	Balance   int64            `json:"balance"`
	Condition BalanceCondition `json:"condition"`
}

type BalanceCondition struct {
	AssetID int `json:"asset_id"`
}

type BlockChainListAddressBalanceResult struct {
	TxId   string
	Detail AddressBalance
}

func (p *BlockChainListAddressBalanceResult) UnmarshalJSON(data []byte) error {

	raw := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "Unmarshal raw object")
	}

	if len(raw) != 2 {
		return errors.Errorf("Invalid operation data: %v", string(data))
	}

	if err := json.Unmarshal(raw[0], &p.TxId); err != nil {
		return errors.Annotate(err, "Unmarshal TxId")
	}

	if err := json.Unmarshal(raw[1], &p.Detail); err != nil {
		return errors.Annotate(err, "Unmarshal BlockChainGetTransactionResultDetail")
	}

	return nil
}

type WalletInfo struct {
	AutomaticBackups           bool        `json:"automatic_backups"`
	DataDir                    string      `json:"data_dir"`
	LastScannedBlockNum        int         `json:"last_scanned_block_num"`
	LastScannedBlockTimestamp  string      `json:"last_scanned_block_timestamp"`
	Name                       string      `json:"name"`
	NumScanningThreads         int         `json:"num_scanning_threads"`
	Open                       bool        `json:"open"`
	ScanProgress               string      `json:"scan_progress"`
	TransactionExpirationSecs  int         `json:"transaction_expiration_secs"`
	TransactionFee             AssetAmount `json:"transaction_fee"`
	TransactionScanningEnabled bool        `json:"transaction_scanning_enabled"`
	Unlocked                   bool        `json:"unlocked"`
	UnlockedUntil              int         `json:"unlocked_until"`
	UnlockedUntilTimestamp     string      `json:"unlocked_until_timestamp"`
	Version                    int         `json:"version"`
}
