package futures

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type stockContractServiceTestSuite struct {
	baseTestSuite
}

func TestStockContractService(t *testing.T) {
	suite.Run(t, new(stockContractServiceTestSuite))
}

func (s *stockContractServiceTestSuite) TestSignStockContract() {
	data := []byte("SUCCESS\n")
	s.mockDo(data, nil)
	defer s.assertDo()
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{}).setParams(params{
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	err := s.client.NewSignStockContractService().Do(newContext(), WithRecvWindow(recvWindow))
	s.r().NoError(err)
}

func (s *stockContractServiceTestSuite) TestSignStockContractNoRecvWindow() {
	data := []byte("SUCCESS")
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{})
		s.assertRequestEqual(e, r)
	})
	err := s.client.NewSignStockContractService().Do(newContext())
	s.r().NoError(err)
}

func (s *stockContractServiceTestSuite) TestSignStockContractNotSuccess() {
	s.mockDo([]byte("ALREADY_SIGNED"), nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		s.assertRequestEqual(newSignedRequest().setFormParams(params{}), r)
	})
	err := s.client.NewSignStockContractService().Do(newContext())
	s.r().Error(err)
	s.r().True(strings.Contains(err.Error(), "ALREADY_SIGNED"), err.Error())
}
