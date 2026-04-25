package futures

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// SignStockContractService signs the USD-margined futures TradFi-Perps (stock) contract user agreement.
// Endpoint: POST /fapi/v1/stock/contract
// See: https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/TradFi-Perps
type SignStockContractService struct {
	c *Client
}

// signStockContractSuccess is the plain text response body on success (per API docs).
const signStockContractSuccess = "SUCCESS"

// Do sends the request. It returns nil only when the trimmed response body is SUCCESS; otherwise it returns an error.
func (s *SignStockContractService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/stock/contract",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	if got := strings.TrimSpace(string(data)); got != signStockContractSuccess {
		return fmt.Errorf("stock contract: %s", got)
	}
	return nil
}
