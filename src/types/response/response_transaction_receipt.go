package response

type TransactionReceipt struct {
	// 트랜잭션 메타데이터
	TransactionHash  string `json:"transactionHash"`  // 트랜잭션 해시값
	TransactionIndex string `json:"transactionIndex"` // 트랜잭션 인덱스
	Type             string `json:"type"`             // 트랜잭션 타입
	Status           string `json:"status"`           // 트랜잭션 상태

	// 블록 정보
	BlockHash   string `json:"blockHash"`   // 블록의 해시값
	BlockNumber string `json:"blockNumber"` // 블록 번호

	// 가스 정보
	CumulativeGasUsed string `json:"cumulativeGasUsed"` // 누적 사용된 가스
	EffectiveGasPrice string `json:"effectiveGasPrice"` // 실제 가스 가격
	GasUsed           string `json:"gasUsed"`           // 사용된 가스 양

	// 주소 정보
	From            string  `json:"from"`            // 송신자 주소
	To              string  `json:"to"`              // 수신자 주소
	ContractAddress *string `json:"contractAddress"` // 스마트 계약 주소 (없을 수도 있음)

	// 로그 정보
	Logs      []Log  `json:"logs"`      // 트랜잭션에 포함된 로그들
	LogsBloom string `json:"logsBloom"` // 로그 블룸 필터
}

type Log struct {
	Address          string   `json:"address"`          // 로그의 주소
	BlockHash        string   `json:"blockHash"`        // 블록의 해시값
	BlockNumber      string   `json:"blockNumber"`      // 블록 번호
	Data             string   `json:"data"`             // 로그 데이터
	Topics           []string `json:"topics"`           // 로그의 주제들
	LogIndex         string   `json:"logIndex"`         // 로그 인덱스
	Removed          bool     `json:"removed"`          // 로그 제거 여부
	TransactionHash  string   `json:"transactionHash"`  // 트랜잭션 해시값
	TransactionIndex string   `json:"transactionIndex"` // 트랜잭션 인덱스
}
