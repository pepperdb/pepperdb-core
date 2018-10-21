package dapp_server

import (
	"time"

	"github.com/pepperdb/pepperdb-core/neblet/pb"
)

// Config dapp_server config
// TODO: move to proto config.
type Config struct {
	Host         string
	Port         int64
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	SaveLog      bool
	LogFile      string
}

// Neblet interface
type Neblet interface {
	Config() *nebletpb.Config
}
