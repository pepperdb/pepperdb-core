package sync

import (
	"errors"

	"github.com/pepperdb/pepperdb-core/common/metrics"
)

// Error Types
var (
	ErrTooSmallGapToSync        = errors.New("the gap between syncpoint and current tail is smaller than a dynasty interval, ignore the sync task")
	ErrCannotFindBlockByHeight  = errors.New("cannot find the block at given height")
	ErrCannotFindBlockByHash    = errors.New("cannot find the block with the given hash")
	ErrWrongChunkHeaderRootHash = errors.New("wrong chunk header root hash")
	ErrWrongChunkDataRootHash   = errors.New("wrong chunk data root hash")
	ErrWrongChunkDataSize       = errors.New("wrong chunk data size")
	ErrInvalidBlockHashInChunk  = errors.New("invalid block hash in chunk data")
	ErrWrongBlockHashInChunk    = errors.New("wrong block hash in chunk data compared with chunk header")
)

// Contants
const (
	MaxChunkPerSyncRequest       = 10
	ConcurrentSyncChunkDataCount = 10
	GetChunkDataTimeout          = 10 // 10s.
)

// Metrics
var (
	metricsCachedSync = metrics.NewGauge("neb.sync.cached")
)
