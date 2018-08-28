package storage

import (
	metrics "github.com/pepperdb/pepperdb-core/common/metrics"
)

// Metrics for storage
var (
	// rocksdb metrics
	metricsRocksdbFlushTime = metrics.NewGauge("neb.rocksdb.flushtime")
	metricsRocksdbFlushLen  = metrics.NewGauge("neb.rocksdb.flushlen")

	metricsBlocksdbCacheSize       = metrics.NewGauge("neb.rocksdb.cache.size")       //block_cache->GetUsage()
	metricsBlocksdbCachePinnedSize = metrics.NewGauge("neb.rocksdb.cachepinned.size") //block_cache->GetPinnedUsage()
	metricsBlocksdbTableReaderMem  = metrics.NewGauge("neb.rocksdb.tablereader.mem")  //estimate-table-readers-mem
	metricsBlocksdbAllMemTables    = metrics.NewGauge("neb.rocksdb.alltables.mem")    //cur-size-all-mem-tables
)
