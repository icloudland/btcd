package tcoinjson

import "github.com/icloudland/btcdx/btcjson"

type GetNewAddressAndKeyCmd struct{}

func NewGetNewAddressAndKeyCmd() *GetNewAddressAndKeyCmd {
	return &GetNewAddressAndKeyCmd{}
}

type ValidateAddressTCmd struct {
	Address string
}

func NewValidateAddressTCmd(address string) *ValidateAddressTCmd {
	return &ValidateAddressTCmd{
		Address: address,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("getaddressandkey", (*GetNewAddressAndKeyCmd)(nil), flags)


}
