package dappserver

import (
	"context"

	"github.com/pepperdb/pepperdb-core/dappserver/pb"
	"github.com/pepperdb/pepperdb-core/storage"
)

// Server implement dapserver.GrpcServer
type GRPCService struct {
	db *storage.RocksStorage
}

// Store a dapp file to storage
func (g *GRPCService) Store(ctx context.Context, in *dappserverpb.StoreRequest) (*dappserverpb.StoreResponse, error) {
	return nil, nil
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
