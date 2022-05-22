package entity

type MineRequestMessage struct {
	Bytes []byte `json:"bytes"`
	Rule  int64  `json:"rule"`
}
