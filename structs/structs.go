package structs

type RPCRequest struct {
	Jsonrpc string                 `json:"jsonrpc"`
	Id      int                    `json:"id"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	Id int `json:"id"`
}
