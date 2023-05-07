package entity

import (
	"time"
)

type Bank struct {
	Oid                 string    `json:"oid"`
	BankCode            *string   `json:"bankCode"`
	BankName            *string   `json:"bankName"`
	Description         *string   `json:"description"`
	OptimisticLockField int       `json:"optimisticLockField"`
	GCRecord            int       `json:"gCRecord"`
	Deleted             bool      `json:"deleted"`
	UserInserted        *string   `json:"userInserted"`
	InsertedDate        time.Time `json:"insertedDate"`
	LastUserId          *string   `json:"lastUserId"`
	LastUpdate          time.Time `json:"lastUpdate"`
}
