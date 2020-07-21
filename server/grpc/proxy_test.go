package grpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.network = network.New(s.T(), network.DefaultConfig())
	s.Require().NotNil(s.network)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestGRPC() {
	val := s.network.Validators[0]
	conn, err := grpc.Dial(
		val.AppConfig.GRPC.Address,
		grpc.WithInsecure(), // Or else we get "no transport security set"
	)
	s.Require().NoError(err)
	testClient := testdata.NewTestServiceClient(conn)
	res, err := testClient.Echo(context.Background(), &testdata.EchoRequest{Message: "hello"})
	s.Require().NoError(err)
	s.Require().Equal("hello", res.Message)
}

func TestCLIIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &IntegrationTestSuite{})
}
