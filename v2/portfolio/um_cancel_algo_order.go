package portfolio

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type UMCancelAlgoOrderService struct {
	c            *Client
	algoId       *int64
	clientAlgoId *string
}

func (s *UMCancelAlgoOrderService) AlgoId(algoId int64) *UMCancelAlgoOrderService {
	s.algoId = &algoId
	return s
}

func (s *UMCancelAlgoOrderService) ClientAlgoId(clientAlgoId string) *UMCancelAlgoOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

func (s *UMCancelAlgoOrderService) Do(ctx context.Context) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/algo/order",
		secType:  secTypeSigned,
	}
	if s.algoId != nil {
		r.setParam("algoId", *s.algoId)
	}
	if s.clientAlgoId != nil {
		r.setParam("clientAlgoId", *s.clientAlgoId)
	}
	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	res := new(uMCancelAlgoOrderServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return err
	}
	if !res.Complete {
		return errors.New("failed to cancel algo order")
	}
	return nil
}

type uMCancelAlgoOrderServiceResponse struct {
	Complete bool `json:"complete"`
}
