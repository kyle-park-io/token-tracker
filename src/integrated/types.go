package integrated

type Result struct {
	Account         string            `json:"account"`
	TokenAddress    string            `json:"tokenAddress"`
	Balance         string            `json:"balance"`
	BalanceHex      string            `json:"balanceHex"`
	TransferHistory []TransferHistory `json:"transfer_history"`
}

type TransferHistory struct {
	TxHash    string `json:"txHash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	ValueHex  string `json:"valueHex"`
	Timestamp int64  `json:"timestamp"`
}
