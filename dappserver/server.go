package dappserver

import (
	"errors"

	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/dappserver/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// DAppServer server of dapp files
type DAppServer struct {
	rpcServer *grpc.Server

	config *dappserverpb.Config
}

// NewServer create the dapp server
func NewServer(config *dappserverpb.Config) (*DAppServer, error) {
	if config == nil {
		logging.CLog().Fatal("Failed to find dapp server config in config file.")
		return nil, errors.New("Create dapp server error: failed to find dapp server config")
	}
	/*
		rpc := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(loggingStream)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(loggingUnary)),
			grpc.MaxRecvMsgSize(config.Dappserver.MaxUploadSize))
	*/
	rpc := grpc.NewServer()

	server := &DAppServer{rpcServer: rpc, config: config}
	grpcService, err := NewGRPCService(config)

	if err != nil {
		return nil, err
	}

	dappserverpb.RegisterDAppServerServer(rpc, grpcService)
	reflection.Register(rpc)

	return server, nil
}

// Start start dapp server
func (ds *DAppServer) Start() error {
	logging.CLog().Info("Starting DAppServer...")
	ds.loop()

	return nil
}

func (ds *DAppServer) loop() {
	logging.CLog().Infof("Start DAppServer at: %s", ds.rpcServer.GetServiceInfo())
	logging.CLog().Info("DAppServer started")
}
