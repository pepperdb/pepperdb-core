package net

import (
	"fmt"

	metrics "github.com/pepperdb/pepperdb-core/common/metrics"
)

// Metrics map for different in/out network msg types
var (
	metricsPacketsIn = metrics.NewMeter("neb.net.packets.in")
	metricsBytesIn   = metrics.NewMeter("neb.net.bytes.in")

	metricsPacketsOut = metrics.NewMeter("neb.net.packets.out")
	metricsBytesOut   = metrics.NewMeter("neb.net.bytes.out")
)

func metricsPacketsInByMessageName(messageName string, size uint64) {
	meter := metrics.NewMeter(fmt.Sprintf("neb.net.packets.in.%s", messageName))
	meter.Mark(1)

	meter = metrics.NewMeter(fmt.Sprintf("neb.net.bytes.in.%s", messageName))
	meter.Mark(int64(size))
}

func metricsPacketsOutByMessageName(messageName string, size uint64) {
	meter := metrics.NewMeter(fmt.Sprintf("neb.net.packets.out.%s", messageName))
	meter.Mark(1)

	meter = metrics.NewMeter(fmt.Sprintf("neb.net.bytes.out.%s", messageName))
	meter.Mark(int64(size))
}
