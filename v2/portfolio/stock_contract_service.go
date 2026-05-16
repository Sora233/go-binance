package portfolio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// SignStockContractService signs the USD-margined futures TradFi-Perps (stock) contract user agreement.
// Endpoint: POST /papi/v1/um/stock/contract
// See: https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/TradFi-Perps
type SignStockContractService struct {
	c *Client
}

// signStockContractSuccess is the plain text response body on success (per API docs).
const signStockContractSuccess = "success"

// Do sends the request. It returns nil only when the trimmed response body is SUCCESS; otherwise it returns an error.
func (s *SignStockContractService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/stock/contract",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	resp := new(APIResponse)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return err
	}
	if strings.EqualFold(resp.Msg, signStockContractSuccess) {
		return fmt.Errorf("stock contract: %s", resp.Msg)
	}
	return nil
}
