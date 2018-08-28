package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/pepperdb/pepperdb-core/neblet/pb"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

// InitCrashReporter init crash reporter
func InitCrashReporter(conf *nebletpb.AppConfig) {
	os.Setenv("GOBACKTRACE", "crash")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to init crash reporter.")
	}
	fp := fmt.Sprintf("%vcrash_%v.log", os.TempDir(), os.Getpid())

	port := rand.Intn(0xFFFF-1024) + 1024
	s, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	for i := 0; i < 0xff; i++ {
		if err != nil {
			port = rand.Intn(0xFFFF-1024) + 1024
			s, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
		} else {
			break
		}
	}

	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to init crash reporter.")
	}
	defer s.Close()

	code := rand.Intn(0xFFFF)
	cmd := exec.Command(fmt.Sprintf("%v/pepp-crashreporter", dir),
		"-logfile",
		fp,
		"-port",
		strconv.Itoa(port),
		"-code",
		strconv.Itoa(code),
		"-pid",
		strconv.Itoa(os.Getpid()),
		"-url",
		conf.CrashReportUrl)

	err = cmd.Start()
	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to start crash reporter daemon.")
	}

	conn, err := s.Accept()
	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to init crash reporter.")
	}
	var buf = make([]byte, 10)
	n, berror := conn.Read(buf)
	if berror != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": berror,
		}).Fatal("Failed to read from conn")
	}
	rs := string(buf[:n])

	if rs == strconv.Itoa(code) {
		if crashFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664); err == nil {
			os.Stderr = crashFile
			syscall.Dup2(int(crashFile.Fd()), 2)
		}
	} else {
		logging.CLog().WithFields(logrus.Fields{
			"rs":   rs,
			"code": code,
			"err":  "code not match",
		}).Fatal("Failed to init crash reporter.")
	}

	logging.CLog().Info("Started crash reporter.")
}
