package dappserver

import (
	"errors"
	"net"

	"golang.org/x/net/netutil"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/dappserver/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Errors
var (
	ErrEmptyRPCListenList = errors.New("empty rpc listen list")
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
	rpc := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(loggingStream)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(loggingUnary)),
		grpc.MaxRecvMsgSize(int(config.Dappserver.MaxUploadSize)))
	server := &DAppServer{rpcServer: rpc, config: config}
	grpcService, err := NewGRPCService(config.Dappserver.DbPath, config.Dappserver.Rocksdb)

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

	if len(ds.config.Dappserver.RpcListen) == 0 {
		return ErrEmptyRPCListenList
	}

	for _, v := range ds.config.Dappserver.RpcListen {
		if err := ds.start(v); err != nil {
			return err
		}
	}

	return nil
}

func (ds *DAppServer) start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to listen to DAppServer GRPCServer")
		return err
	}

	logging.CLog().WithFields(logrus.Fields{
		"address": addr,
	}).Info("Stared DAppServer.")

	connectionLimits := ds.config.Dappserver.ConnectionLimits
	if connectionLimits == 0 {
		connectionLimits = 128
	}

	listener = netutil.LimitListener(listener, int(connectionLimits))

	go func() {
		if err := ds.rpcServer.Serve(listener); err != nil {
			logging.CLog().WithFields(logrus.Fields{
				"err": err,
			}).Info("DAppServer exited.")
		}
	}()

	return nil
}

// Stop stops the DAppServer and closes listener.
func (ds *DAppServer) Stop() {
	logging.CLog().WithFields(logrus.Fields{
		"listen": ds.config.Dappserver.RpcListen,
	}).Info("Stopping DAppServer...")

	ds.rpcServer.Stop()

	logging.CLog().Info("Stopped DAppServer.")
}
