package models

import (
	"errors"
	"fmt"
	"strconv"
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

var (
	ErrAccountNotFound = errors.New("资金账户不存在")
	ErrAmountNotEnough = errors.New("可用余额不足")
)

type CustomerOperation struct {
	CustomerId int
	Amount     int64
	OpType     int
	BsType     int
	Detail     string
	Ext        string
}

func (e *CustomerOperation) Key() string {
	return fmt.Sprintf("%v", strconv.Itoa(e.CustomerId))
}
