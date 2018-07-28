package main

import "fmt"

// generated with: https://mholt.github.io/json-to-go/

type RPCErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   struct {
		Code    int           `json:"code"`
		Name    string        `json:"name"`
		What    string        `json:"what"`
		Details []interface{} `json:"details"`
	} `json:"error"`
}

func (r RPCErrorResponse) String() string {
	return fmt.Sprintf("[%d] %s: %s", r.Error.Code, r.Error.Name, r.Error.What)
}

type GetInfoResponse struct {
	ServerVersion            string `json:"server_version"`
	ChainID                  string `json:"chain_id"`
	HeadBlockNum             int    `json:"head_block_num"`
	LastIrreversibleBlockNum int    `json:"last_irreversible_block_num"`
	LastIrreversibleBlockID  string `json:"last_irreversible_block_id"`
	HeadBlockID              string `json:"head_block_id"`
	HeadBlockTime            string `json:"head_block_time"`
	HeadBlockProducer        string `json:"head_block_producer"`
	VirtualBlockCPULimit     int    `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit     int    `json:"virtual_block_net_limit"`
	BlockCPULimit            int    `json:"block_cpu_limit"`
	BlockNetLimit            int    `json:"block_net_limit"`
}

type GetBlockRequest struct {
	BlockNumOrID string `json:"block_num_or_id"`
}

type GetBlockResponse struct {
	RefBlockPrefix  int64         `json:"ref_block_prefix"`
	BlockNum        int           `json:"block_num"`
	ID              string        `json:"id"`
	BlockExtensions []interface{} `json:"block_extensions"`
	Transactions    []struct {
		Trx struct {
			Transaction struct {
				TransactionExtensions []interface{} `json:"transaction_extensions"`
				Expiration            string        `json:"expiration"`
				RefBlockNum           int           `json:"ref_block_num"`
				RefBlockPrefix        int64         `json:"ref_block_prefix"`
				MaxNetUsageWords      int           `json:"max_net_usage_words"`
				MaxCPUUsageMs         int           `json:"max_cpu_usage_ms"`
				DelaySec              int           `json:"delay_sec"`
				ContextFreeActions    []interface{} `json:"context_free_actions"`
				Actions               []struct {
					HexData       string      `json:"hex_data"`
					Data          interface{} `json:"data"`
					Authorization []struct {
						Permission string `json:"permission"`
						Actor      string `json:"actor"`
					} `json:"authorization"`
					Name    string `json:"name"`
					Account string `json:"account"`
				} `json:"actions"`
			} `json:"transaction"`
			PackedTrx             string        `json:"packed_trx"`
			ContextFreeData       []interface{} `json:"context_free_data"`
			PackedContextFreeData string        `json:"packed_context_free_data"`
			Compression           string        `json:"compression"`
			Signatures            []string      `json:"signatures"`
			ID                    string        `json:"id"`
		} `json:"trx"`
		NetUsageWords int    `json:"net_usage_words"`
		CPUUsageUs    int    `json:"cpu_usage_us"`
		Status        string `json:"status"`
	} `json:"transactions"`
	ProducerSignature string        `json:"producer_signature"`
	HeaderExtensions  []interface{} `json:"header_extensions"`
	Timestamp         string        `json:"timestamp"`
	Producer          string        `json:"producer"`
	Confirmed         int           `json:"confirmed"`
	Previous          string        `json:"previous"`
	TransactionMroot  string        `json:"transaction_mroot"`
	ActionMroot       string        `json:"action_mroot"`
	ScheduleVersion   int           `json:"schedule_version"`
	NewProducers      interface{}   `json:"new_producers"`
}
