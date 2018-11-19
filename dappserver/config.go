package dappserver

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"

	"github.com/gogo/protobuf/proto"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/dappserver/pb"
)

// LoadConfig loads configuration from file.
func LoadConfig(file string) *dappserverpb.Config {
	var conf string
	b, err := ioutil.ReadFile(file)
	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to read the config file: %s.", file)
	}

	conf = string(b)

	pb := new(dappserverpb.Config)
	if err := proto.UnmarshalText(conf, pb); err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to parse the config file: ", file)
	}
	return pb
}
