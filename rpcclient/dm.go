package rpcclient

import (
	"encoding/json"
	"fmt"
	"github.com/icloudland/btcdx/qstring"
	"net/http"
	"sync"
	"io/ioutil"
	"bytes"
	"github.com/icloudland/btcdx/btcjson"
	"strings"
	"github.com/icloudland/btcdx/dmjson"
	"strconv"
	"github.com/tidwall/gjson"
	"errors"
	"net/http/httputil"
	"net/url"
	"crypto/tls"
	"crypto/x509"
	"time"
	"net"
)

type DmClient struct {
	config *ConnConfig

	httpClient *http.Client

	sendPostChan chan *sendDmHttpDetails
	shutdown     chan struct{}
	wg           sync.WaitGroup

	Debug bool
}

func NewDmClient(config *ConnConfig) (*DmClient, error) {
	var httpClient *http.Client
	var start bool
	if config.HTTPPostMode {
		start = true

		var err error
		httpClient, err = newDmHTTPClient(config)
		if err != nil {
			return nil, err
		}
	}

	client := &DmClient{
		config:       config,
		httpClient:   httpClient,
		sendPostChan: make(chan *sendDmHttpDetails, sendPostBufferSize),
		shutdown:     make(chan struct{}),
	}

	if start {
		client.start()
	}

	return client, nil
}

func newDmHTTPClient(config *ConnConfig) (*http.Client, error) {
	// Set proxy function if there is a proxy configured.
	var proxyFunc func(*http.Request) (*url.URL, error)
	if config.Proxy != "" {
		proxyURL, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, err
		}
		proxyFunc = http.ProxyURL(proxyURL)
	}

	// Configure TLS if needed.
	var tlsConfig *tls.Config
	if !config.DisableTLS {
		if len(config.Certificates) > 0 {
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(config.Certificates)
			tlsConfig = &tls.Config{
				RootCAs: pool,
			}
		}

		if tlsConfig == nil {
			tlsConfig = &tls.Config{}
		}
		tlsConfig.InsecureSkipVerify = true
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: tlsConfig,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	return &client, nil
}

// start begins processing input and output messages.
func (c *DmClient) start() {
	// Start the I/O processing handlers depending on whether the client is
	// in HTTP POST mode or the default websocket mode.
	if c.config.HTTPPostMode {
		c.wg.Add(1)
		go c.sendGetHandler()
	}
}

func (c *DmClient) sendGetHandler() {
out:
	for {
		// Send any messages ready for send until the shutdown channel
		// is closed.
		select {
		case details := <-c.sendPostChan:
			c.handleSendHttpMessage(details)

		case <-c.shutdown:
			break out
		}
	}
cleanup:
	for {
		select {
		case details := <-c.sendPostChan:
			details.jsonDmRequest.responseChan <- &dmResponse{
				status: "0",
				err:    ErrClientShutdown,
			}

		default:
			break cleanup
		}
	}
	c.wg.Done()
}

func (c *DmClient) handleSendHttpMessage(details *sendDmHttpDetails) {
	jReq := details.jsonDmRequest
	httpResponse, err := c.httpClient.Do(details.httpRequest)
	if err != nil {
		jReq.responseChan <- &dmResponse{err: err}
		return
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		jReq.responseChan <- &dmResponse{err: err}
		return
	}

	// Try to unmarshal the response as a regular JSON-RPC response.
	var resp dmHttpResponse
	fmt.Println(string(respBytes[:]))
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		// When the response itself isn't a valid JSON-RPC response
		// return an error which includes the HTTP status code and raw
		// response bytes.
		err = fmt.Errorf("status code: %d, response: %q",
			httpResponse.StatusCode, string(respBytes))
		jReq.responseChan <- &dmResponse{err: err}
		return
	}

	jReq.responseChan <- &dmResponse{data: resp.Data, info: resp.Info, status: resp.Status,}
}

type dmResponse struct {
	status string
	info   string
	data   []byte
	err    error
}

type dmHttpResponse struct {
	Status string          `json:"status"`
	Info   string          `json:"info"`
	Data   json.RawMessage `json:"data"`
}

type jsonDmRequest struct {
	requestType    string
	method         string
	cmd            interface{}
	marshalledJSON []byte
	responseChan   chan *dmResponse
}

func newDmFutureError(err error) chan *dmResponse {
	responseChan := make(chan *dmResponse, 1)
	responseChan <- &dmResponse{err: err}
	return responseChan
}

func (c *DmClient) sendDmGetCmd(cmd interface{}) chan *dmResponse {
	// Get the method associated with the command.
	method, err := btcjson.CmdMethod(cmd)
	if err != nil {
		return newDmFutureError(err)
	}
	methods := strings.Split(method, ":")
	method = methods[1]

	// Marshal the command.
	queryString, err := qstring.MarshalString(cmd)
	if err != nil {
		return newDmFutureError(err)
	}
	if queryString != "" {
		method = method + "?" + queryString
	}
	// Generate the request and send it along with a channel to respond on.
	responseChan := make(chan *dmResponse, 1)
	jReq := &jsonDmRequest{
		requestType:    "GET",
		method:         method,
		cmd:            cmd,
		marshalledJSON: nil,
		responseChan:   responseChan,
	}
	c.sendDmRequest(jReq)

	return responseChan
}

func (c *DmClient) sendDmRequest(jReq *jsonDmRequest) {
	//c.sendDmGet(jReq)
	c.call(jReq)
}

func (c *DmClient) sendDmGet(jReq *jsonDmRequest) {
	//Generate a request to the configured RPC server.
	protocol := "http"
	if !c.config.DisableTLS {
		protocol = "https"
	}
	url := protocol + "://" + c.config.Host + "/" + jReq.method
	fmt.Println(url)
	bodyReader := bytes.NewReader(jReq.marshalledJSON)
	httpReq, err := http.NewRequest(jReq.requestType, url, bodyReader)
	if err != nil {
		jReq.responseChan <- &dmResponse{status: "0", err: err}
		return
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")

	// Configure basic access authorization.
	if c.config.User != "" {
		httpReq.SetBasicAuth(c.config.User, c.config.Pass)
	}

	c.sendDmHttpRequest(httpReq, jReq)
}

func (c *DmClient) call(jReq *jsonDmRequest) {

	protocol := "http"
	if !c.config.DisableTLS {
		protocol = "https"
	}
	url := protocol + "://" + c.config.Host + "/" + jReq.method

	req, err := http.NewRequest(jReq.requestType, url, nil)
	if err != nil {
		jReq.responseChan <- &dmResponse{status: "0", err: err}
		return
	}

	if c.Debug {
		// Useful when debugging API calls
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("-------------------------------")
		fmt.Println(string(requestDump))
	}

	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		jReq.responseChan <- &dmResponse{err: err}
		return
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		jReq.responseChan <- &dmResponse{err: err}
		return
	}

	if c.Debug {
		fmt.Println("RESPONSE:")
		fmt.Println(string(respBytes[:]))
	}

	// Try to unmarshal the response as a regular JSON-RPC response.
	var resp dmHttpResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		// When the response itself isn't a valid JSON-RPC response
		// return an error which includes the HTTP status code and raw
		// response bytes.
		err = fmt.Errorf("status code: %d, response: %q",
			httpResponse.StatusCode, string(respBytes))
		jReq.responseChan <- &dmResponse{err: err}
		return
	}

	jReq.responseChan <- &dmResponse{data: resp.Data, info: resp.Info, status: resp.Status,}

}

func (c *DmClient) sendDmHttpRequest(httpReq *http.Request, jReq *jsonDmRequest) {

	c.sendPostChan <- &sendDmHttpDetails{
		jsonDmRequest: jReq,
		httpRequest:   httpReq,
	}
}

type sendDmHttpDetails struct {
	httpRequest   *http.Request
	jsonDmRequest *jsonDmRequest
}

func receiveDmFuture(f chan *dmResponse) ([]byte, error) {
	// Wait for a response on the returned channel.
	r := <-f
	if r.status == "0" {
		return nil, errors.New(r.info)
	}
	return r.data, r.err
}

type FutureDmGetBlockCount chan *dmResponse

func (r FutureDmGetBlockCount) Receive() (int64, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return 0, err
	}

	var count string
	err = json.Unmarshal(res, &count)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(count, 10, 64)

}

func (c *DmClient) DmGetBlockCountAsync() FutureDmGetBlockCount {

	cmd := dmjson.NewDmGetBlockCountCmd()
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetBlockCount() (int64, error) {
	return c.DmGetBlockCountAsync().Receive()
}

type FutureDmGetBlockHeight chan *dmResponse

func (r FutureDmGetBlockHeight) Receive() (int64, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return 0, err
	}

	rawJson := string(res[:])
	height := gjson.Get(rawJson, "Height").String()

	return strconv.ParseInt(height, 10, 64)

}

func (c *DmClient) DmGetBlockHeightAsync() FutureDmGetBlockHeight {

	cmd := dmjson.NewDmGetBlockHeightCmd()
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetBlockHeight() (int64, error) {
	return c.DmGetBlockHeightAsync().Receive()
}

type FutureDmGetTransactionsByBlockId chan *dmResponse

func (r FutureDmGetTransactionsByBlockId) Receive() ([]dmjson.DmBlockTransaction, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res[:]))
	var transactions []dmjson.DmBlockTransaction
	err = json.Unmarshal(res, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil

}

func (c *DmClient) DmGetTransactionsByBlockIdAsync(
	blockHeight int64) FutureDmGetTransactionsByBlockId {

	cmd := dmjson.NewDmGetTransactionsByBlockIdCmd(blockHeight)
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetTransactionsByBlockId(
	blockHeight int64) ([]dmjson.DmBlockTransaction, error) {
	return c.DmGetTransactionsByBlockIdAsync(blockHeight).Receive()
}

type FutureDmCreateTransaction chan *dmResponse

func (r FutureDmCreateTransaction) Receive() (error) {
	_, err := receiveDmFuture(r)
	if err != nil {
		return err
	}

	return nil
}

func (c *DmClient) DmCreateTransactionAsync(cc string, from string, to string,
	amo string, remark string, txid string, nonce string, sign string, fee string) FutureDmCreateTransaction {

	cmd := dmjson.NewDmCreateTransactionCmd(cc, from, to, amo, remark, txid, nonce, sign, fee)
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmCreateTransaction(cc string, from string, to string,
	amo string, remark string, txid string, nonce string, sign string, fee string) (error) {
	return c.DmCreateTransactionAsync(cc, from, to, amo, remark, txid, nonce, sign, fee).Receive()
}

type FutureDmGetTransactionId chan *dmResponse

func (r FutureDmGetTransactionId) Receive() (*dmjson.DmGetTransactionIdResutl, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return nil, err
	}

	var tRes dmjson.DmGetTransactionIdResutl
	err = json.Unmarshal(res, &tRes)
	if err != nil {
		return nil, err
	}

	return &tRes, nil
}

func (c *DmClient) DmGetTransactionIdAsync(cnt int) FutureDmGetTransactionId {

	cmd := dmjson.NewDmGetTransactionIdCmd(cnt)
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetTransactionId(cnt int) (*dmjson.DmGetTransactionIdResutl, error) {
	return c.DmGetTransactionIdAsync(cnt).Receive()
}

type FutureGetTransactionDetail chan *dmResponse

func (r FutureGetTransactionDetail) Receive() (*dmjson.DmGetTransactionDetailResult, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return nil, err
	}

	var tRes dmjson.DmGetTransactionDetailResult
	err = json.Unmarshal(res, &tRes)
	if err != nil {
		return nil, err
	}

	return &tRes, nil
}

func (r FutureGetTransactionDetail) ReceiveNew() ([]*dmjson.DmGetTransactionDetailResult, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return nil, err
	}

	var tRes []*dmjson.DmGetTransactionDetailResult
	err = json.Unmarshal(res, &tRes)
	if err != nil {
		return nil, err
	}

	return tRes, nil
}

func (c *DmClient) DmGetTransactionDetailAsync(tx string) FutureGetTransactionDetail {

	cmd := dmjson.NewDmGetTransactionDetailCmd(tx)
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetTransactionDetail(tx string) (*dmjson.DmGetTransactionDetailResult, error) {
	return c.DmGetTransactionDetailAsync(tx).Receive()
}

func (c *DmClient) DmGetTransactionDetailNew(tx string) ([]*dmjson.DmGetTransactionDetailResult, error) {
	return c.DmGetTransactionDetailAsync(tx).ReceiveNew()
}

type FutureDmGetNewAddress chan *dmResponse

func (r FutureDmGetNewAddress) Receive() (*dmjson.DmGetNewAddressResult, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return nil, err
	}

	var tRes dmjson.DmGetNewAddressResult
	err = json.Unmarshal(res, &tRes)
	if err != nil {
		return nil, err
	}

	return &tRes, nil
}

func (c *DmClient) DmGetNewAddressAsync(psw string) FutureDmGetNewAddress {

	cmd := dmjson.NewDmGetNewAddressCmd(psw)
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetNewAddress(psw string) (*dmjson.DmGetNewAddressResult, error) {
	return c.DmGetNewAddressAsync(psw).Receive()
}

type FutureDmGetBalance chan *dmResponse

func (r FutureDmGetBalance) Receive(coinName string) (string, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return "", err
	}
	rawJson := string(res[:])
	balStr := gjson.Get(rawJson, coinName).String()
	return balStr, nil
}

func (c *DmClient) DmGetBalanceAsync(address string,
	coinName string) FutureDmGetBalance {

	cmd := dmjson.NewDmGetBalanceCmd(address, coinName)
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetBalance(address string, coinName string) (
	string, error) {
	return c.DmGetBalanceAsync(address, coinName).Receive(coinName)
}

type FutureDmGetTokenInfo chan *dmResponse

func (r FutureDmGetTokenInfo) Receive() (*dmjson.DmTokenInfo, error) {
	res, err := receiveDmFuture(r)
	if err != nil {
		return nil, err
	}

	var tokenInfo dmjson.DmTokenInfo
	err = json.Unmarshal(res, &tokenInfo)
	if err != nil {
		return nil, err
	}

	return &tokenInfo, nil
}

func (c *DmClient) DmGetTokenInfoAsync(cc string) FutureDmGetTokenInfo {

	cmd := dmjson.NewDmGetTokenInfoCmd(cc)
	return c.sendDmGetCmd(cmd)
}

func (c *DmClient) DmGetTokenInfo(cc string) (*dmjson.DmTokenInfo, error) {
	return c.DmGetTokenInfoAsync(cc).Receive()
}
