package dmjson

import "github.com/icloudland/btcdx/btcjson"

type DmGetNewAddressCmd struct {
	Cmd string
	Psw string
}

func NewDmGetNewAddressCmd(psw string) *DmGetNewAddressCmd {
	return &DmGetNewAddressCmd{
		Cmd: "newac",
		Psw: psw,
	}
}

type DmGetBlockCountCmd struct {
	Method string
}

func NewDmGetBlockCountCmd() *DmGetBlockCountCmd {
	return &DmGetBlockCountCmd{
		Method: "max",
	}
}

type DmGetTransactionsByBlockIdCmd struct {
	Method string
	Blk    int64
}

func NewDmGetTransactionsByBlockIdCmd(blockHeight int64) *DmGetTransactionsByBlockIdCmd {
	return &DmGetTransactionsByBlockIdCmd{
		Method: "qryTxByBlk",
		Blk:    blockHeight,
	}
}

type DmGetTransactionIdCmd struct {
	Cnt int
}

func NewDmGetTransactionIdCmd(cnt int) *DmGetTransactionIdCmd {
	return &DmGetTransactionIdCmd{
		Cnt: cnt,
	}
}

type DmGetTransactionDetailCmd struct {
	Method string
	Tx     string
}

func NewDmGetTransactionDetailCmd(tx string) *DmGetTransactionDetailCmd {
	return &DmGetTransactionDetailCmd{
		Method: "txrst",
		Tx:     tx,
	}
}

type DmGetBalanceCmd struct {
	Cmd  string
	Args string
	Cc   string
}

func NewDmGetBalanceCmd(address string, token string) *DmGetBalanceCmd {
	return &DmGetBalanceCmd{
		Cmd:  "qryBal",
		Args: address,
		Cc:   token,
	}
}

type DmCreateTransactionCmd struct {
	Cc     string
	From   string
	To     string
	Amo    string
	Remark string
	Txid   string
	Nonce  string
	Sign   string
}

func NewDmCreateTransactionCmd(cc string, from string, to string,
	amo string, remark string, txid string, nonce string, sign string) *DmCreateTransactionCmd {
	return &DmCreateTransactionCmd{
		Cc:     cc,
		From:   from,
		To:     to,
		Amo:    amo,
		Remark: remark,
		Txid:   txid,
		Nonce:  nonce,
		Sign:   sign,
	}
}

type DmGetTokenInfoCmd struct {
	Cc string
}

func NewDmGetTokenInfoCmd(cc string) *DmGetTokenInfoCmd {
	return &DmGetTokenInfoCmd{
		Cc: cc,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("dm:acc:1", (*DmGetNewAddressCmd)(nil), flags)
	btcjson.MustRegisterCmd("dm:block:1", (*DmGetBlockCountCmd)(nil), flags)
	btcjson.MustRegisterCmd("dm:trade:1", (*DmGetTransactionsByBlockIdCmd)(nil), flags)
	btcjson.MustRegisterCmd("dm:trade:2", (*DmGetTransactionDetailCmd)(nil), flags)
	btcjson.MustRegisterCmd("dm:crtTx:1", (*DmGetTransactionIdCmd)(nil), flags)
	btcjson.MustRegisterCmd("dm:qry:1", (*DmGetBalanceCmd)(nil), flags)
	btcjson.MustRegisterCmd("dm:ivk:1", (*DmCreateTransactionCmd)(nil), flags)
	btcjson.MustRegisterCmd("dm:tkinfo:1", (*DmGetTokenInfoCmd)(nil), flags)


}
