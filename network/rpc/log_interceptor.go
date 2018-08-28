package rpc

import (
	"strings"

	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func loggingStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logging.VLog().WithFields(logrus.Fields{
		"method": info.FullMethod,
	}).Info("Rpc request.")
	metricsRPCCounter.Mark(1)

	return handler(srv, ss)
}

func loggingUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	if strings.Contains(info.FullMethod, "ApiService") {
		logging.VLog().WithFields(logrus.Fields{
			"method": info.FullMethod,
			"params": req,
		}).Info("Rpc request.")
	} else {
		logging.VLog().WithFields(logrus.Fields{
			"method": info.FullMethod,
		}).Info("Rpc request.")
	}

	metricsRPCCounter.Mark(1)

	return handler(ctx, req)
}
