package tcoinjson

import "github.com/icloudland/btcdx/btcjson"

type GetNewAddressAndKeyCmd struct{}

func NewGetNewAddressAndKeyCmd() *GetNewAddressAndKeyCmd {
	return &GetNewAddressAndKeyCmd{}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("getaddressandkey", (*GetNewAddressAndKeyCmd)(nil), flags)

}