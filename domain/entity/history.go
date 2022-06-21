package entity

import "time"

type DocumentHistory struct {
	Id                int       `json:"id"`
	Uid               string    `json:"uid"`
	Document          Document  `json:"document"`
	Hash              string    `json:"hash"`
	PoW               int       `json:"pow"`
	Initiator         User      `json:"initiator"`
	DateChanged       time.Time `json:"date_changed"`
	DateChangedString string    `json:"date_changed_string"`
}
