package crosschain

// TxInput is input data to a tx. Depending on the blockchain it can include nonce, recent block hash, account id, ...
type TxInput interface {
}

// TxInfo is a unified view of common tx info across multiple blockchains. Use it as an example to build your own.
type TxInfo struct {
	TxID            string
	From            Address
	To              Address
	ToAlt           Address
	ContractAddress ContractAddress
	Amount          AmountBlockchain
	Fee             AmountBlockchain
	BlockIndex      int64
	BlockTime       int64
	Confirmations   int64
}

// TxHash is a tx hash or id
type TxHash string

// TxDataToSign is the payload that Signer needs to sign, when "signing a tx". It's sometimes called a sighash.
type TxDataToSign []byte

// TxSignature is a tx signature
type TxSignature []byte

// Tx is a transaction
type Tx interface {
	Hash() TxHash
	Sighash() (TxDataToSign, error)
	AddSignature(TxSignature) error
}
