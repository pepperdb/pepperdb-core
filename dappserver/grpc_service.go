package dappserver

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/dappserver/pb"
	"github.com/pepperdb/pepperdb-core/storage"
)

// GRPCService implement dapserver.GrpcServer
type GRPCService struct {
	db *storage.RocksStorage
}

// Store a dapp file to storage
func (g *GRPCService) Store(ctx context.Context, in *dappserverpb.StoreRequest) (*dappserverpb.StoreResponse, error) {
	if in.Name == "" || len(in.File) == 0 || in.Md5 == "" || in.Type == "" {
		logging.CLog().Info("Bad store request parameter")
		return nil, errors.New("Bad store request parameter")
	}
	addr := "test_dapp_server_addr"
	key := addr + strconv.FormatInt(time.Now().UnixNano(), 10)
	if err := g.db.Put([]byte(key), in.File); err != nil {
		logging.CLog().Errorf("DAppServer rocksdb save data error: %s", err)
		return nil, err
	}
	res := &dappserverpb.StoreResponse{
		Ok:      true,
		Address: key,
	}
	return res, nil
}

// NewGRPCService create a GRPCService
func NewGRPCService(dbPath string, config *dappserverpb.RocksDBConfig) (*GRPCService, error) {
	rocksdb, err := storage.NewRocksStorage(dbPath)
	if err != nil {
		return nil, err
	}

	grpc := &GRPCService{db: rocksdb}
	return grpc, nil
}
