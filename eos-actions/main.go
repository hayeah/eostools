package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

func apiRequest(method string, req interface{}, res interface{}) (err error) {
	url := fmt.Sprintf("http://127.0.0.1:8888/v1/%s", method)

	var postBody io.Reader

	if req != nil {
		var buf bytes.Buffer

		bodyEncoder := json.NewEncoder(&buf)
		err = bodyEncoder.Encode(req)
		if err != nil {
			return
		}

		postBody = &buf
		// wtf... why does io.TeeReader cause postBody to be empty?
		// postBody = io.TeeReader(&buf, os.Stdout)
	}

	r, err := http.Post(url, "application/json", postBody)
	if err != nil {
		return
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var rpcErr RPCErrorResponse
		dec := json.NewDecoder(r.Body)
		dec.Decode(&rpcErr)

		return errors.Errorf("POST %s %s", method, rpcErr)
	}

	body := r.Body
	// body := io.TeeReader(r.Body, os.Stdout)

	dec := json.NewDecoder(body)
	err = dec.Decode(&res)
	if err != nil {
		return
	}

	return nil
}

var lastBlockSeen int

type actionFilter struct {
	name    string
	account string
}

func showBlock(blocknum int) (err error) {
	var block GetBlockResponse
	err = apiRequest("chain/get_block", &GetBlockRequest{
		// BlockNumOrID: info.LastIrreversibleBlockID,
		BlockNumOrID: fmt.Sprintf("%d", blocknum),
		// BlockNumOrID: "0000a41517787e2294b271a55d9413504266b5c04ac03b731a8455764c1faacb",

		// BlockNumOrID: "0000a41517787e2294b271a55d9413504266b5c04ac03b731a8455764c1faace",
	}, &block)

	if err != nil {
		return
	}

	if len(block.Transactions) == 0 {
		// log.Printf("#%d is empty", block.BlockNum)
		return
	}

	// spew.Dump(block)

	for _, trx := range block.Transactions {
		for _, action := range trx.Trx.Transaction.Actions {
			if *filterAccountNameExact != "" && action.Account != *filterAccountNameExact {
				continue
			}

			if *filterActionNameExact != "" && action.Name != *filterActionNameExact {
				continue
			}

			enc := json.NewEncoder(os.Stdout)
			enc.Encode(action)

			// data, _ := json.Marshal(action.Data)
			// fmt.Println(data)
			// fmt.Printf("%s:%s\t%s\n", action.Account, action.Name, string(data))
			// action.Data
		}
	}

	return
}

func pollTx() (err error) {
	var info GetInfoResponse
	err = apiRequest("chain/get_info", nil, &info)
	if err != nil {
		return
	}

	// pretty.Println(info)
	// pretty.Println("last block id:", info.LastIrreversibleBlockID)
	for {
		// TODO may need to do parallel block lookup
		if lastBlockSeen >= info.LastIrreversibleBlockNum {
			break
		}

		if lastBlockSeen == 0 {
			lastBlockSeen = info.LastIrreversibleBlockNum
		} else {
			lastBlockSeen++
		}

		// fmt.Println("args", os.Args)
		// filter := os.Args[1]

		err = showBlock(lastBlockSeen)
		if err != nil {
			log.Println("show block err:", err)
		}
	}

	return
}

func main() {
	cli.Parse(os.Args[1:])

	tick := time.Tick(500 * time.Millisecond)
	for {
		<-tick
		err := pollTx()
		if err != nil {
			log.Println("err", err)
		}
	}
}
