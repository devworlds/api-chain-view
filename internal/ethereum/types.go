package ethereum

type Wallet struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

type Transaction struct {
	From     string
	To       string
	Value    string
	Gas      string
	GasPrice string
	Nonce    string
	Data     string
}

type Block struct {
	Number       string
	Hash         string
	ParentHash   string
	Nonce        string
	Difficulty   string
	ExtraData    string
	GasLimit     string
	GasUsed      string
	Timestamp    string
	Transactions []Transaction
}
