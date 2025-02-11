package db

type Transaction struct {
	// 트랜잭션 메타데이터
	Hash             string `gorm:"primaryKey;type:varchar(66)" json:"hash"`  // 트랜잭션 해시 (기본 키)
	BlockHash        string `gorm:"type:varchar(66)" json:"blockHash"`        // 블록 해시
	BlockNumber      string `gorm:"type:varchar(32)" json:"blockNumber"`      // 블록 번호
	TransactionIndex string `gorm:"type:varchar(32)" json:"transactionIndex"` // 트랜잭션 인덱스
	ChainID          string `gorm:"type:varchar(32)" json:"chainId"`          // 체인 ID
	Type             string `gorm:"type:varchar(32)" json:"type"`             // 트랜잭션 유형

	// 트랜잭션 실행 정보
	From                 string `gorm:"type:varchar(42)" json:"from"`                 // 송신 주소
	To                   string `gorm:"type:varchar(42)" json:"to"`                   // 수신 주소
	Value                string `gorm:"type:varchar(64)" json:"value"`                // 전송된 금액 (Wei 단위)
	Gas                  string `gorm:"type:varchar(32)" json:"gas"`                  // 제공된 가스량
	GasPrice             string `gorm:"type:varchar(64)" json:"gasPrice"`             // 가스 가격
	MaxFeePerGas         string `gorm:"type:varchar(64)" json:"maxFeePerGas"`         // 최대 가스 요금
	MaxPriorityFeePerGas string `gorm:"type:varchar(64)" json:"maxPriorityFeePerGas"` // 최대 우선순위 가스 요금

	// 트랜잭션 데이터 및 서명
	Input   string `gorm:"type:text" json:"input"`          // 트랜잭션 데이터
	Nonce   string `gorm:"type:varchar(32)" json:"nonce"`   // 논스 값
	R       string `gorm:"type:varchar(66)" json:"r"`       // 서명 R 값
	S       string `gorm:"type:varchar(66)" json:"s"`       // 서명 S 값
	V       string `gorm:"type:varchar(10)" json:"v"`       // 서명 V 값
	YParity string `gorm:"type:varchar(10)" json:"yParity"` // Y 좌표 패리티 값

	// 트랜잭션 접근 정보
	AccessList string `gorm:"type:jsonb" json:"accessList"` // JSON 형식으로 저장
}

type TransactionReceipt struct {
	// 트랜잭션 메타데이터
	TransactionHash  string `gorm:"primaryKey;type:varchar(66)" json:"transactionHash"` // 트랜잭션 해시값
	TransactionIndex string `gorm:"type:varchar(32)" json:"transactionIndex"`           // 트랜잭션 인덱스
	Type             string `gorm:"type:varchar(32)" json:"type"`                       // 트랜잭션 타입
	Status           string `gorm:"type:varchar(32)" json:"status"`                     // 트랜잭션 상태

	// 블록 정보
	BlockHash   string `gorm:"type:varchar(66)" json:"blockHash"`   // 블록의 해시값
	BlockNumber string `gorm:"type:varchar(32)" json:"blockNumber"` // 블록 번호

	// 가스 정보
	CumulativeGasUsed string `gorm:"type:varchar(64)" json:"cumulativeGasUsed"` // 누적 사용된 가스
	EffectiveGasPrice string `gorm:"type:varchar(64)" json:"effectiveGasPrice"` // 실제 가스 가격
	GasUsed           string `gorm:"type:varchar(64)" json:"gasUsed"`           // 사용된 가스 양

	// 주소 정보
	From            string  `gorm:"type:varchar(42)" json:"from"`            // 송신자 주소
	To              string  `gorm:"type:varchar(42)" json:"to"`              // 수신자 주소
	ContractAddress *string `gorm:"type:varchar(42)" json:"contractAddress"` // 스마트 계약 주소 (없을 수도 있음)

	// 로그 정보
	Logs      []Log  `gorm:"foreignKey:TransactionHash;constraint:OnDelete:CASCADE;" json:"logs"` // 트랜잭션에 포함된 로그들
	LogsBloom string `gorm:"type:text" json:"logsBloom"`                                          // 로그 블룸 필터
}

type Log struct {
	ID               uint   `gorm:"primaryKey" json:"id"`                          // 고유 ID (자동 증가)
	Address          string `gorm:"type:varchar(42)" json:"address"`               // 로그의 주소
	BlockHash        string `gorm:"type:varchar(66)" json:"blockHash"`             // 블록의 해시값
	BlockNumber      string `gorm:"type:varchar(32)" json:"blockNumber"`           // 블록 번호
	Data             string `gorm:"type:text" json:"data"`                         // 로그 데이터
	Topics           string `gorm:"type:jsonb" json:"topics"`                      // 로그의 주제들
	LogIndex         string `gorm:"type:varchar(32)" json:"logIndex"`              // 로그 인덱스
	Removed          bool   `gorm:"type:boolean" json:"removed"`                   // 로그 제거 여부
	TransactionHash  string `gorm:"type:varchar(66);index" json:"transactionHash"` // 트랜잭션 해시값
	TransactionIndex string `gorm:"type:varchar(32)" json:"transactionIndex"`      // 트랜잭션 인덱스
}
