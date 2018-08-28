package rpc

import (
	"github.com/pepperdb/pepperdb-core/common/metrics"
)

// Metrics for rpc
var (
	metricsRPCCounter = metrics.NewMeter("neb.rpc.request")

	metricsAccountStateSuccess = metrics.NewMeter("neb.rpc.account.success")
	metricsAccountStateFailed  = metrics.NewMeter("neb.rpc.account.failed")

	metricsSendTxSuccess = metrics.NewMeter("neb.rpc.sendTx.success")
	metricsSendTxFailed  = metrics.NewMeter("neb.rpc.sendTx.failed")

	metricsSignTxSuccess = metrics.NewMeter("neb.rpc.signTx.success")
	metricsSignTxFailed  = metrics.NewMeter("neb.rpc.signTx.failed")

	metricsUnlockSuccess = metrics.NewMeter("neb.rpc.unlock.success")
	metricsUnlockFailed  = metrics.NewMeter("neb.rpc.unlock.failed")
)
