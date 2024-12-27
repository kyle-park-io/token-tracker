package response

type BlockNumber string

type Balance string

type BlockWithoutTransactions struct {
	Block

	// 트랜잭션 해시 목록
	Transactions []string `json:"transactions"` // 블록에 포함된 트랜잭션 해시 목록
}

type BlockWithTransactions struct {
	Block

	// 트랜잭션 목록
	Transactions []Transaction `json:"transactions"` // 블록에 포함된 트랜잭션 목록
}

// Block represents the structure of the block returned by eth_getBlockByNumber
type Block struct {
	// 블록 메타데이터
	Number                string `json:"number"`                // 블록 번호
	Hash                  string `json:"hash"`                  // 블록 해시
	ParentHash            string `json:"parentHash"`            // 부모 블록의 해시
	ParentBeaconBlockRoot string `json:"parentBeaconBlockRoot"` // 비콘 체인 상 부모 블록의 루트 해시
	Timestamp             string `json:"timestamp"`             // 블록 생성 타임스탬프 (Unix 시간)

	// 채굴 관련 정보
	Difficulty string `json:"difficulty"` // 블록의 채굴 난이도
	Miner      string `json:"miner"`      // 블록을 생성한 채굴자의 주소
	Nonce      string `json:"nonce"`      // 블록 채굴 시 사용된 논스(nonce)
	MixHash    string `json:"mixHash"`    // 채굴의 유효성을 확인하는 해시 값

	// 블록 상태 정보
	StateRoot        string `json:"stateRoot"`        // 블록 상태의 루트 해시
	ReceiptsRoot     string `json:"receiptsRoot"`     // 트랜잭션 영수증 루트 해시
	TransactionsRoot string `json:"transactionsRoot"` // 트랜잭션 루트 해시
	Sha3Uncles       string `json:"sha3Uncles"`       // 엉클 블록의 해시 값(SHA3)
	LogsBloom        string `json:"logsBloom"`        // 로그 필터링을 위한 블룸 필터

	// 블록 가스 정보
	GasLimit      string `json:"gasLimit"`      // 블록의 가스 사용 한도
	GasUsed       string `json:"gasUsed"`       // 블록에서 실제 사용된 가스량
	BaseFeePerGas string `json:"baseFeePerGas"` // 블록의 기본 가스 요금 (EIP-1559 기준)
	BlobGasUsed   string `json:"blobGasUsed"`   // 블록에서 사용된 블롭(blob) 가스량
	ExcessBlobGas string `json:"excessBlobGas"` // 블록의 초과 블롭(blob) 가스량

	// 추가 정보
	ExtraData string `json:"extraData"` // 블록에 포함된 추가 데이터
	Size      string `json:"size"`      // 블록의 크기 (바이트 단위)

	// 엉클 블록 정보
	Uncles []string `json:"uncles"` // 엉클 블록 목록

	// 출금 정보
	Withdrawals     []Withdrawal `json:"withdrawals"`     // 블록에 포함된 출금 정보
	WithdrawalsRoot string       `json:"withdrawalsRoot"` // 출금 루트 해시
}

// Withdrawal 구조체 정의
type Withdrawal struct {
	Address        string `json:"address"`        // 출금을 받는 주소
	Amount         string `json:"amount"`         // 출금된 금액
	Index          string `json:"index"`          // 출금의 인덱스
	ValidatorIndex string `json:"validatorIndex"` // 출금을 요청한 검증자의 인덱스
}

type Transaction struct {
	// 트랜잭션 메타데이터
	Hash             string `json:"hash"`             // 트랜잭션 해시
	BlockHash        string `json:"blockHash"`        // 트랜잭션이 포함된 블록의 해시
	BlockNumber      string `json:"blockNumber"`      // 트랜잭션이 포함된 블록의 번호
	TransactionIndex string `json:"transactionIndex"` // 블록 내 트랜잭션의 인덱스
	ChainID          string `json:"chainId"`          // 트랜잭션이 발생한 체인의 ID (예: 이더리움 메인넷 0x1)
	Type             string `json:"type"`             // 트랜잭션 유형 (예: Legacy, EIP-1559 등)

	// 트랜잭션 실행 정보
	From                 string `json:"from"`                 // 트랜잭션을 보낸 주소
	To                   string `json:"to"`                   // 트랜잭션을 받는 주소
	Value                string `json:"value"`                // 트랜잭션에서 전송된 금액 (Wei 단위)
	Gas                  string `json:"gas"`                  // 트랜잭션 실행을 위해 제공된 가스량
	GasPrice             string `json:"gasPrice"`             // 트랜잭션에 지불한 가스 가격
	MaxFeePerGas         string `json:"maxFeePerGas"`         // 트랜잭션에서 허용된 최대 가스 요금 (EIP-1559 기준)
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"` // 트랜잭션에서 허용된 최대 우선순위 가스 요금 (채굴자 우선순위 보상을 위한 최대 비용)

	// 트랜잭션 데이터 및 서명
	Input   string `json:"input"`   // 트랜잭션 데이터 (바이트 코드 등)
	Nonce   string `json:"nonce"`   // 트랜잭션 발신자의 논스 값 (계정에서 보낸 트랜잭션의 순서)
	R       string `json:"r"`       // 트랜잭션 서명의 R 값
	S       string `json:"s"`       // 트랜잭션 서명의 S 값
	V       string `json:"v"`       // 트랜잭션 서명의 V 값
	YParity string `json:"yParity"` // 서명 검증에 사용된 Y 좌표의 패리티

	// 트랜잭션 접근 정보
	AccessList []interface{} `json:"accessList"` // 트랜잭션에서 접근 가능한 주소와 키 목록(EIP-2930에 정의된 접근 가능 목록)
}
