package dappserver

import (
	"time"

	"github.com/pepperdb/pepperdb-core/neblet/pb"
)

// Config DAppServer config
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
