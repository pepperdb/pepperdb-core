package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/pepperdb/pepperdb-core/core"
	"github.com/pepperdb/pepperdb-core/core/state"
	"github.com/pepperdb/pepperdb-core/nf/nvm"
	"github.com/pepperdb/pepperdb-core/storage"
)

func main() {
	data, _ := ioutil.ReadFile(os.Args[1])

	mem, _ := storage.NewMemoryStorage()
	context, _ := state.NewWorldState(nil, mem)
	contract, _ := context.CreateContractAccount([]byte("account2"), nil, nil)

	ctx, err := nvm.NewContext(core.MockBlock(nil, 1), nil, contract, context)
	if err == nil {
		engine := nvm.NewV8Engine(ctx)
		result, err := engine.RunScriptSource(string(data), 0)

		log.Fatalf("Result is %s. Err is %s", result, err)

		time.Sleep(10 * time.Second)
		engine.Dispose()
	} else {
		os.Exit(1)
	}
}
