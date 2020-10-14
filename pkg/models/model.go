package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

const (
	OpTypeAdd = 1
	OpTypeSub = 2
)

type CustomerOperation struct {
	CustomerId int
	Amount     int
	OpType     int
	BsType     int
	Detail     string
	Ext        string
}
