package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

type UMOpenAlgoOrdersService struct {
	c      *Client
	symbol string
	algoId *int64
}

func (s *UMOpenAlgoOrdersService) Symbol(symbol string) *UMOpenAlgoOrdersService {
	s.symbol = symbol
	return s
}

func (s *UMOpenAlgoOrdersService) AlgoId(algoId int64) *UMOpenAlgoOrdersService {
	s.algoId = &algoId
	return s
}

func (s *UMOpenAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*UMOpenAlgoOrdersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/algo/openAlgoOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.algoId != nil {
		r.setParam("algoId", *s.algoId)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*UMOpenAlgoOrdersResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UMOpenAlgoOrdersResponse struct {
	AlgoId        int64  `json:"algoId"`
	ClientAlgoId  string `json:"clientAlgoId"`
	AlgoType      string `json:"algoType"`
	OrderType     string `json:"orderType"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	TimeInForce   string `json:"timeInForce"`
	Quantity      string `json:"quantity"`
	AlgoStatus    string `json:"algoStatus"`
	ActualOrderId string `json:"actualOrderId"`
	ActualPrice   string `json:"actualPrice"`
	TriggerPrice  string `json:"triggerPrice"`
	Price         string `json:"price"`
	// IcebergQuantity         any `json:"icebergQuantity"`
	TPTriggerPrice          string `json:"tpTriggerPrice"`
	TPPrice                 string `json:"tpPrice"`
	SLTriggerPrice          string `json:"slTriggerPrice"`
	SLPrice                 string `json:"slPrice"`
	TPOrderType             string `json:"tpOrderType"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	WorkingType             string `json:"workingType"`
	PriceMatch              string `json:"priceMatch"`
	ClosePosition           bool   `json:"closePosition"`
	PriceProtect            bool   `json:"priceProtect"`
	ReduceOnly              bool   `json:"reduceOnly"`
	CreateTime              int64  `json:"createTime"`
	UpdateTime              int64  `json:"updateTime"`
	TriggerTime             int64  `json:"triggerTime"`
	GoodTillDate            int64  `json:"goodTillDate"`
}
