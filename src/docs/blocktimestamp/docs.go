// Package blocktimestamp Code generated by swaggo/swag. DO NOT EDIT
package blocktimestamp

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/example/helloworld": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get/getBlock": {
            "get": {
                "description": "Fetches a block from the blockchain based on its number with optional transaction details.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Retrieve a block by number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block number in decimal or hexadecimal format",
                        "name": "number",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Include transactions in the block (true or false)",
                        "name": "withTxs",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BlockWithTransactions"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get/getBlockPosition": {
            "get": {
                "description": "Fetches the block position for a specific date in the given timezone. Returns block timestamp and position details.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Retrieve block position by date",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Date in 'YYYY-MM-DD' format",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/router.ResponseBlockPosition"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get/getBlockTimestamp": {
            "get": {
                "description": "Fetches the block timestamp for a specific block number in both Unix and hexadecimal formats.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Retrieve block timestamp by block number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Block number in decimal or hexadecimal format",
                        "name": "number",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/router.ResponseBlockTimestamp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get/getERC20Balance": {
            "get": {
                "description": "Fetches the balance of a specific ERC-20 token for a given Ethereum account, returning it in both decimal and hexadecimal formats.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Retrieve the balance of an ERC-20 token for a given Ethereum account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum account address",
                        "name": "account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ERC-20 token contract address",
                        "name": "tokenAddress",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(optional) Block tag (e.g., 'latest', 'earliest', or a block number in decimal/hexadecimal)",
                        "name": "tag",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved the ERC-20 token balance",
                        "schema": {
                            "$ref": "#/definitions/router.ResponseBalance"
                        }
                    },
                    "400": {
                        "description": "Invalid input, such as incorrect Ethereum address, token address, or tag",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get/getETHBalance": {
            "get": {
                "description": "Fetches the balance of a given Ethereum account, returning it in both decimal and hexadecimal formats.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Retrieve the balance of an Ethereum account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum account address",
                        "name": "account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "(optional) Block tag (e.g., 'latest', 'earliest', or a block number in decimal/hexadecimal)",
                        "name": "tag",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved the Ethereum balance",
                        "schema": {
                            "$ref": "#/definitions/router.ResponseBalance"
                        }
                    },
                    "400": {
                        "description": "Invalid input, such as incorrect Ethereum address or tag",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get/getLatestBlockNumber": {
            "get": {
                "description": "Fetches the latest block number from the blockchain in both decimal and hexadecimal formats.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Retrieve the latest block number",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/router.ResponseBlockNumber"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/get/getRandomBlock": {
            "get": {
                "description": "Fetches a random block from the blockchain with optional transaction details.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Block"
                ],
                "summary": "Retrieve a random block",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Include transactions in the block (true or false)",
                        "name": "withTxs",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BlockWithTransactions"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/track/trackERC20": {
            "get": {
                "description": "Tracks ERC20 token transfers for a specific account and token contract within a given date and target count.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Track"
                ],
                "summary": "Track ERC20 token transfers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum account address (e.g., '0x1234567890abcdef1234567890abcdef12345678')",
                        "name": "account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ERC20 token contract address (e.g., '0xabc123')",
                        "name": "tokenAddress",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Date in 'YYYY-MM-DD' format",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of transactions to retrieve",
                        "name": "targetCount",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Time limit for processing in seconds",
                        "name": "timeLimit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/integrated.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/track/trackETH": {
            "get": {
                "description": "Tracks Ethereum account activity, including transactions and balance changes, within a given date and target count.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Track"
                ],
                "summary": "Track Ethereum account balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum account address (e.g., '0x1234567890abcdef1234567890abcdef12345678')",
                        "name": "account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Date in 'YYYY-MM-DD' format",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of transactions to retrieve",
                        "name": "targetCount",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Time limit for processing in seconds",
                        "name": "timeLimit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/integrated.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/track/trackETH2": {
            "get": {
                "description": "Tracks Ethereum account activity, including transactions and balance changes, within a given date and target count.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Track"
                ],
                "summary": "Track Ethereum account balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum account address (e.g., '0x1234567890abcdef1234567890abcdef12345678')",
                        "name": "account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Date in 'YYYY-MM-DD' format",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of transactions to retrieve",
                        "name": "targetCount",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Time limit for processing in seconds",
                        "name": "timeLimit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/integrated.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/router.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "integrated.Result": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "balance": {
                    "type": "string"
                },
                "balanceHex": {
                    "type": "string"
                },
                "tokenAddress": {
                    "type": "string"
                },
                "transfer_history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/integrated.TransferHistory"
                    }
                }
            }
        },
        "integrated.TransferHistory": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "integer"
                },
                "to": {
                    "type": "string"
                },
                "txHash": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                },
                "valueHex": {
                    "type": "string"
                }
            }
        },
        "response.BlockWithTransactions": {
            "type": "object",
            "properties": {
                "baseFeePerGas": {
                    "description": "블록의 기본 가스 요금 (EIP-1559 기준)",
                    "type": "string"
                },
                "blobGasUsed": {
                    "description": "블록에서 사용된 블롭(blob) 가스량",
                    "type": "string"
                },
                "difficulty": {
                    "description": "채굴 관련 정보",
                    "type": "string"
                },
                "excessBlobGas": {
                    "description": "블록의 초과 블롭(blob) 가스량",
                    "type": "string"
                },
                "extraData": {
                    "description": "추가 정보",
                    "type": "string"
                },
                "gasLimit": {
                    "description": "블록 가스 정보",
                    "type": "string"
                },
                "gasUsed": {
                    "description": "블록에서 실제 사용된 가스량",
                    "type": "string"
                },
                "hash": {
                    "description": "블록 해시",
                    "type": "string"
                },
                "logsBloom": {
                    "description": "로그 필터링을 위한 블룸 필터",
                    "type": "string"
                },
                "miner": {
                    "description": "블록을 생성한 채굴자의 주소",
                    "type": "string"
                },
                "mixHash": {
                    "description": "채굴의 유효성을 확인하는 해시 값",
                    "type": "string"
                },
                "nonce": {
                    "description": "블록 채굴 시 사용된 논스(nonce)",
                    "type": "string"
                },
                "number": {
                    "description": "블록 메타데이터",
                    "type": "string"
                },
                "parentBeaconBlockRoot": {
                    "description": "비콘 체인 상 부모 블록의 루트 해시",
                    "type": "string"
                },
                "parentHash": {
                    "description": "부모 블록의 해시",
                    "type": "string"
                },
                "receiptsRoot": {
                    "description": "트랜잭션 영수증 루트 해시",
                    "type": "string"
                },
                "sha3Uncles": {
                    "description": "엉클 블록의 해시 값(SHA3)",
                    "type": "string"
                },
                "size": {
                    "description": "블록의 크기 (바이트 단위)",
                    "type": "string"
                },
                "stateRoot": {
                    "description": "블록 상태 정보",
                    "type": "string"
                },
                "timestamp": {
                    "description": "블록 생성 타임스탬프 (Unix 시간)",
                    "type": "string"
                },
                "transactions": {
                    "description": "트랜잭션 목록",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Transaction"
                    }
                },
                "transactionsRoot": {
                    "description": "트랜잭션 루트 해시",
                    "type": "string"
                },
                "uncles": {
                    "description": "엉클 블록 정보",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "withdrawals": {
                    "description": "출금 정보",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Withdrawal"
                    }
                },
                "withdrawalsRoot": {
                    "description": "출금 루트 해시",
                    "type": "string"
                }
            }
        },
        "response.Transaction": {
            "type": "object",
            "properties": {
                "accessList": {
                    "description": "트랜잭션 접근 정보",
                    "type": "array",
                    "items": {}
                },
                "blockHash": {
                    "description": "트랜잭션이 포함된 블록의 해시",
                    "type": "string"
                },
                "blockNumber": {
                    "description": "트랜잭션이 포함된 블록의 번호",
                    "type": "string"
                },
                "chainId": {
                    "description": "트랜잭션이 발생한 체인의 ID (예: 이더리움 메인넷 0x1)",
                    "type": "string"
                },
                "from": {
                    "description": "트랜잭션 실행 정보",
                    "type": "string"
                },
                "gas": {
                    "description": "트랜잭션 실행을 위해 제공된 가스량",
                    "type": "string"
                },
                "gasPrice": {
                    "description": "트랜잭션에 지불한 가스 가격",
                    "type": "string"
                },
                "hash": {
                    "description": "트랜잭션 메타데이터",
                    "type": "string"
                },
                "input": {
                    "description": "트랜잭션 데이터 및 서명",
                    "type": "string"
                },
                "maxFeePerGas": {
                    "description": "트랜잭션에서 허용된 최대 가스 요금 (EIP-1559 기준)",
                    "type": "string"
                },
                "maxPriorityFeePerGas": {
                    "description": "트랜잭션에서 허용된 최대 우선순위 가스 요금 (채굴자 우선순위 보상을 위한 최대 비용)",
                    "type": "string"
                },
                "nonce": {
                    "description": "트랜잭션 발신자의 논스 값 (계정에서 보낸 트랜잭션의 순서)",
                    "type": "string"
                },
                "r": {
                    "description": "트랜잭션 서명의 R 값",
                    "type": "string"
                },
                "s": {
                    "description": "트랜잭션 서명의 S 값",
                    "type": "string"
                },
                "to": {
                    "description": "트랜잭션을 받는 주소",
                    "type": "string"
                },
                "transactionIndex": {
                    "description": "블록 내 트랜잭션의 인덱스",
                    "type": "string"
                },
                "type": {
                    "description": "트랜잭션 유형 (예: Legacy, EIP-1559 등)",
                    "type": "string"
                },
                "v": {
                    "description": "트랜잭션 서명의 V 값",
                    "type": "string"
                },
                "value": {
                    "description": "트랜잭션에서 전송된 금액 (Wei 단위)",
                    "type": "string"
                },
                "yParity": {
                    "description": "서명 검증에 사용된 Y 좌표의 패리티",
                    "type": "string"
                }
            }
        },
        "response.Withdrawal": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "출금을 받는 주소",
                    "type": "string"
                },
                "amount": {
                    "description": "출금된 금액",
                    "type": "string"
                },
                "index": {
                    "description": "출금의 인덱스",
                    "type": "string"
                },
                "validatorIndex": {
                    "description": "출금을 요청한 검증자의 인덱스",
                    "type": "string"
                }
            }
        },
        "router.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Description of the error"
                }
            }
        },
        "router.ResponseBalance": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string",
                    "example": "100000000"
                },
                "hexBalance": {
                    "type": "string",
                    "example": "0x5f5e100"
                }
            }
        },
        "router.ResponseBlockNumber": {
            "type": "object",
            "properties": {
                "blockNumber": {
                    "description": "Decimal block number",
                    "type": "integer",
                    "example": 12345678
                },
                "hexBlockNumber": {
                    "description": "Hexadecimal block number",
                    "type": "string",
                    "example": "0xabcdef"
                }
            }
        },
        "router.ResponseBlockPosition": {
            "type": "object",
            "properties": {
                "blockPosition": {
                    "$ref": "#/definitions/tracker.BlockPosition"
                },
                "blockTimestamp": {
                    "$ref": "#/definitions/router.ResponseBlockTimestamp"
                }
            }
        },
        "router.ResponseBlockTimestamp": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "Original date",
                    "type": "string",
                    "example": "2024-01-01"
                },
                "hexTimestamp": {
                    "description": "Hexadecimal timestamp",
                    "type": "string",
                    "example": "0x62d43b1f"
                },
                "timestamp": {
                    "description": "Unix timestamp",
                    "type": "integer",
                    "example": 1672531199
                }
            }
        },
        "tracker.BlockPosition": {
            "type": "object",
            "properties": {
                "high": {
                    "type": "integer"
                },
                "highHex": {
                    "type": "string"
                },
                "highTimestamp": {
                    "type": "integer"
                },
                "highTimestampHex": {
                    "type": "string"
                },
                "low": {
                    "type": "integer"
                },
                "lowHex": {
                    "type": "string"
                },
                "lowTimestamp": {
                    "type": "integer"
                },
                "lowTimestampHex": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
