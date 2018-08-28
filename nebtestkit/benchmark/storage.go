package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pepperdb/pepperdb-core/common/trie"
	"github.com/pepperdb/pepperdb-core/storage"
	"github.com/pepperdb/pepperdb-core/common/util/byteutils"
)

func main() {
	file := os.Args[1]
	cnt, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		fmt.Println("Parse Int Error ", err)
		return
	}
	rootHash := os.Args[3]
	stor, err := storage.NewDiskStorage(file)
	if err != nil {
		fmt.Println("OpenDB Error ", err)
		return
	}
	fmt.Println("cnt:", cnt, " file:", file, " root:", rootHash)

	startAt := time.Now().Unix()
	root, err := byteutils.FromHex(rootHash)
	if err != nil {
		fmt.Println("Parse Hex Error ", err)
		return
	}
	txsState, err := trie.NewTrie(root, stor, false)
	if err != nil {
		fmt.Println("NewTrie Error ", err)
		return
	}
	iter, err := txsState.Iterator(nil)
	if err != nil {
		fmt.Println("Iterator Error ", err)
		return
	}
	exist, err := iter.Next()
	if err != nil {
		fmt.Println("Next Error1 ", err)
		return
	}
	i := cnt
	for exist {
		exist, err = iter.Next()
		i--
		if i == 0 {
			fmt.Println("Read Over")
			break
		}
		if err != nil {
			fmt.Println("Next Error2 ", err)
			return
		}
	}
	endAt := time.Now().Unix()
	diff := endAt - startAt
	if diff == 0 {
		fmt.Println("Diff is zero")
		return
	}
	op := cnt - i
	fmt.Println("TPS ", op, "/", diff, "s = ", op/diff)
	fmt.Println("Done")
}
