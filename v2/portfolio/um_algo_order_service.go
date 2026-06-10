package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMAlgoOrderService service to place UM algo orders
type UMAlgoOrderService struct {
	c                       *Client
	clientAlgoId            *string
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               string
	timeInForce             *TimeInForceType
	quantity                *string
	reduceOnly              *bool
	triggerPrice            *string
	price                   *string
	workingType             *string
	priceProtect            *bool
	activatePrice           *string
	callbackRate            *string
	priceMatch              *PriceMatchType
	selfTradePreventionMode *SelfTradePreventionMode
	goodTillDate            *int64
}

// 用户自定义的条件订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
func (s *UMAlgoOrderService) ClientAlgoId(clientAlgoId string) *UMAlgoOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

// Symbol set symbol
func (s *UMAlgoOrderService) Symbol(symbol string) *UMAlgoOrderService {
	s.symbol = symbol
	return s
}

// Side 买卖方向 SELL, BUY
func (s *UMAlgoOrderService) Side(side SideType) *UMAlgoOrderService {
	s.side = side
	return s
}

// PositionSide 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
func (s *UMAlgoOrderService) PositionSide(positionSide PositionSideType) *UMAlgoOrderService {
	s.positionSide = &positionSide
	return s
}

// OrderType 条件订单类型 STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
func (s *UMAlgoOrderService) OrderType(orderType string) *UMAlgoOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set time in force
func (s *UMAlgoOrderService) TimeInForce(timeInForce TimeInForceType) *UMAlgoOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *UMAlgoOrderService) Quantity(quantity string) *UMAlgoOrderService {
	s.quantity = &quantity
	return s
}

// TriggerPrice 触发价
func (s *UMAlgoOrderService) TriggerPrice(triggerPrice string) *UMAlgoOrderService {
	s.triggerPrice = &triggerPrice
	return s
}

// Price 委托价格
func (s *UMAlgoOrderService) Price(price string) *UMAlgoOrderService {
	s.price = &price
	return s
}

// ActivationPrice set activation price
func (s *UMAlgoOrderService) ActivatePrice(price string) *UMAlgoOrderService {
	s.activatePrice = &price
	return s
}

// CallbackRate 追踪止损回调比例，可取值范围[0.1, 10],其中 1代表1% ,仅TRAILING_STOP_MARKET 需要此参数
func (s *UMAlgoOrderService) CallbackRate(rate string) *UMAlgoOrderService {
	s.callbackRate = &rate
	return s
}

// WorkingType 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
func (s *UMAlgoOrderService) WorkingType(workingType string) *UMAlgoOrderService {
	s.workingType = &workingType
	return s
}

// PriceProtect 条件单触发保护："true","false", 默认"false".
func (s *UMAlgoOrderService) PriceProtect(protect bool) *UMAlgoOrderService {
	s.priceProtect = &protect
	return s
}

// ReduceOnly 只减仓 true, false; 非双开模式下默认false；双开模式下不接受此参数
func (s *UMAlgoOrderService) ReduceOnly(reduceOnly bool) *UMAlgoOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// GoodTillDate TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
func (s *UMAlgoOrderService) GoodTillDate(goodTillDate int64) *UMAlgoOrderService {
	s.goodTillDate = &goodTillDate
	return s
}

// PriceMatch OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
func (s *UMAlgoOrderService) PriceMatch(priceMatch PriceMatchType) *UMAlgoOrderService {
	s.priceMatch = &priceMatch
	return s
}

// Do send request
func (s *UMAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *UMAlgoOrder, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/algo/order",
		secType:  secTypeSigned,
	}
	r.setParam("algoType", "CONDITIONAL")
	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("type", s.orderType)

	if s.positionSide != nil {
		r.setParam("positionSide", *s.positionSide)
	}
	if s.timeInForce != nil {
		r.setParam("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		r.setParam("quantity", *s.quantity)
	}
	if s.reduceOnly != nil {
		r.setParam("reduceOnly", *s.reduceOnly)
	}
	if s.triggerPrice != nil {
		r.setParam("triggerPrice", *s.triggerPrice)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.activatePrice != nil {
		r.setParam("activatePrice", *s.activatePrice)
	}
	if s.callbackRate != nil {
		r.setParam("callbackRate", *s.callbackRate)
	}
	if s.workingType != nil {
		r.setParam("workingType", *s.workingType)
	}
	if s.priceProtect != nil {
		r.setParam("priceProtect", *s.priceProtect)
	}
	if s.clientAlgoId != nil {
		r.setParam("clientAlgoId", *s.clientAlgoId)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMAlgoOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMAlgoOrder define algo order info
type UMAlgoOrder struct {
	AlgoId                  int64                   `json:"algoId"`
	ClientAlgoId            string                  `json:"clientAlgoId"`
	AlgoType                string                  `json:"algoType"`  // ENUM: CONDITIONAL
	OrderType               string                  `json:"orderType"` // ENUM:STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	Symbol                  string                  `json:"symbol"`
	Side                    SideType                `json:"side"`
	PositionSide            PositionSideType        `json:"positionSide"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	Quantity                string                  `json:"quantity"`
	AlgoStatus              OrderStatusType         `json:"algoStatus"`
	TriggerPrice            string                  `json:"triggerPrice"`
	Price                   string                  `json:"price"`
	IcebergQuantity         *int64                  `json:"icebergQuantity"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	WorkingType             string                  `json:"workingType"`
	PriceMatch              string                  `json:"priceMatch"`
	PriceProtect            bool                    `json:"priceProtect"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	ActivatePrice           string                  `json:"activatePrice"`
	CallbackRate            string                  `json:"callbackRate"`
	CreateTime              int64                   `json:"createTime"`
	UpdateTime              int64                   `json:"updateTime"`
	TriggerTime             int64                   `json:"triggerTime"`
	GoodTillDate            int64                   `json:"goodTillDate"`
}
