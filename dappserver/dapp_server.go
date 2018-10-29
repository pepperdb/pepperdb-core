package dappserver

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/dappserver/pb"
	"github.com/pepperdb/pepperdb-core/storage"
)

// DAppServer server of DApp file
type DAppServer struct {
	config *dappserverpb.Config

	server *http.Server
}

// NewDAppServer create the dapp server
func NewDAppServer(config *dappserverpb.Config) (*DAppServer, error) {
	/*if networkConf := n.Config().GetNetwork(); networkConf == nil {
		logging.CLog().Fatal("Failed to find dapp_server config in config file")
		return nil, errors.New("config.conf should has dapp_server")
	}
	config := &dappserverpb.DAppServerConfig{
		Host:           "127.0.0.1",
		Port:           8000,
		ReadTimeoutMs:  300,
		WriteTimeoutMs: 300,
	}*/

	mux := http.NewServeMux()
	fh, err := newFileHandler(20*1024*1024, "./dapp-db")
	if err != nil {
		logging.CLog().Error("Create file handler error")
	}
	mux.Handle("/upload", fh)

	server := &http.Server{
		Addr:         config.Dappserver.GetHost() + ":" + strconv.Itoa(int(config.Dappserver.GetPort())),
		ReadTimeout:  time.Duration(config.Dappserver.GetReadTimeoutMs()) * time.Millisecond,
		WriteTimeout: time.Duration(config.Dappserver.GetWriteTimeoutMs()) * time.Millisecond,
		Handler:      mux,
	}

	if true {
		logFile, err := os.Create("./dappserver.log")
		if err != nil {
			logging.CLog().Error("Create dapp server log file error")
			return nil, errors.New("Create dapp server log file error")
		}
		logger := log.New(logFile, "dapp_server_", log.Ldate|log.Ltime|log.Lshortfile)
		server.ErrorLog = logger
	}

	ds := &DAppServer{
		config: config,
		server: server,
	}

	return ds, nil
}

// Config returns DApp server configuration.
func (ds *DAppServer) Config() *dappserverpb.Config {
	return ds.config
}

// Start start dapp server
func (ds *DAppServer) Start() error {
	logging.CLog().Info("Starting DAppServer...")
	ds.loop()

	return nil
}

func (ds *DAppServer) loop() {
	ds.server.ListenAndServe()
	logging.CLog().Info("DAppServer started")
}

// fileHandler handler of file
type fileHandler struct {
	maxFileSize int64
	db          *storage.RocksStorage
}

// NewfileHandler create a fileHandler
func newFileHandler(maxFileSize int64, dbPath string) (*fileHandler, error) {
	if maxFileSize == 0 {
		maxFileSize = 20 * 1024 * 1024
	}
	if dbPath == "" {
		dbPath = "./dapp-db"
	}

	rocksdb, _ := storage.NewRocksStorage(dbPath)

	fh := &fileHandler{
		maxFileSize: maxFileSize,
		db:          rocksdb,
	}
	return fh, nil
}

func (fh *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, fh.maxFileSize)
	if err := r.ParseMultipartForm(fh.maxFileSize); err != nil {
		logging.CLog().Errorf("dapp_server http error: %s\n", err)
		renderError(w, "ERROR: file too big", http.StatusBadRequest)
		return
	}
	fileType := r.PostFormValue("type")
	if fileType == "" {
		renderError(w, "ERROR: need value: type", http.StatusBadRequest)
		return
	}

	// TODO other needed post value

	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		//logging.CLog().Errorf("dapp_server read file error: %s\n", err)
		renderError(w, "ERROR: invalid upload file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		renderError(w, "ERROR: read upload file fail", http.StatusBadRequest)
		return
	}
	address := "test_dapp_server_addr"
	key := address + strconv.FormatInt(time.Now().UnixNano(), 10)
	dbErr := fh.db.Put([]byte(key), fileBytes)
	if err != nil {
		logging.CLog().Errorf("dapp_server rocksdb save data error: %s\n", dbErr)
		renderError(w, "ERROR: internal error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("uplaod success"))
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
