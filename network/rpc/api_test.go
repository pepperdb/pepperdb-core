package rpc

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pepperdb/pepperdb-core/network/rpc/mock_pb"
	"github.com/pepperdb/pepperdb-core/network/rpc/pb"
	"github.com/pepperdb/pepperdb-core/common/util"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestAPIService_GetNebState(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_pb.NewMockAPIServiceClient(ctrl)

	{
		req := &rpcpb.NonParamsRequest{}
		expected := &rpcpb.GetNebStateResponse{Tail: "hac"}
		client.EXPECT().GetNebState(gomock.Any(), gomock.Any()).Return(expected, nil)
		resp, _ := client.GetNebState(context.Background(), req)
		assert.Equal(t, expected, resp)
	}

	// TODO: test with mock neblet.
}

func TestGetAccountState(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_pb.NewMockAPIServiceClient(ctrl)

	{
		req := &rpcpb.GetAccountStateRequest{Address: "0xf"}
		tmpNumber, _ := util.NewUint128FromInt(31415926)
		bal := tmpNumber.String()
		expected := &rpcpb.GetAccountStateResponse{Balance: bal, Nonce: 1}
		client.EXPECT().GetAccountState(gomock.Any(), gomock.Any()).Return(expected, nil)
		resp, _ := client.GetAccountState(context.Background(), req)
		assert.Equal(t, expected, resp)
	}

	// TODO: test with mock neblet.
}

func TestSendTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_pb.NewMockAPIServiceClient(ctrl)

	{
		req := &rpcpb.TransactionRequest{From: "0xf"}
		expected := &rpcpb.SendTransactionResponse{Txhash: "0x2"}
		client.EXPECT().SendTransaction(gomock.Any(), gomock.Any()).Return(expected, nil)
		resp, _ := client.SendTransaction(context.Background(), req)
		assert.Equal(t, expected, resp)
	}

	// TODO: test with mock neblet.
}
