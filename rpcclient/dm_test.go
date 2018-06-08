package rpcclient

import (
	"testing"
	"fmt"
	"encoding/json"
	"github.com/icloudland/btcdx/dmjson"
)

type WalletTransferToAddressCmd struct {
	Amount          string
	AssetSymbol     string
	FromAccountName string
	ToAddress       string
}
type DmGetBlockCountCmd struct {
	Method string
}

func TestSendCmd(t *testing.T) {
	connCfg := &ConnConfig{
		Host:         "117.82.138.202:8081",
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	c, err := NewDmClient(connCfg)
	if err != nil {
		fmt.Println(err)
	}

	//cmd := &DmGetBlockCountCmd{
	//	Method:"max",
	//}

	aa := sendaaabb(c, dmjson.NewDmGetBlockCountCmd())
	aa.Receive()

	close(c.shutdown)
	select {

	}

}


type FutureWalletCheckAddressCmd1 chan *dmResponse

// Receive waits for the response promised by the future and returns a hash string
func (r FutureWalletCheckAddressCmd1) Receive() (int64, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal the result as an CallContractResult.
	var b string
	fmt.Println(string(res[:]))
	err = json.Unmarshal(res, &b)
	if err != nil {
		return 0, err
	}
	fmt.Println(b)
	return 0, nil
}

// WalletCheckAddressAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
func sendaaabb(c *DmClient, cmd interface{}) FutureWalletCheckAddressCmd1 {

	return c.sendDmGetCmd(cmd)
}

