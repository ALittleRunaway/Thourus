package entity

type ValidateRequestMessage struct {
	Bytes []byte `json:"bytes"`
	Hash  string `json:"hash"`
	PoW   int64  `json:"pow"`
}
