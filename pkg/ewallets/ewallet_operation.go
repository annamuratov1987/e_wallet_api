package ewallets

import "gorm.io/gorm"

type EwalletOperationType uint8

const (
	TYPE_CREDIT EwalletOperationType = iota
	TYPE_DEBIT
)

type EwalletOperation struct {
	gorm.Model
	EwalletId int64 `gorm:"not null"`
	Type EwalletOperationType `gorm:"not null"`
	Amount int64
}
